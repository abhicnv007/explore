package main

import (
	"context"
	"encoding/json"
	"explore/cache"
	"log"
	"net/http"
	"regexp"

	"github.com/valyala/fastjson"
)

type wikiResult struct {
	Title     string `json:"title"`
	PageID    int    `json:"pageid"`
	Size      int    `json:"size"`
	WordCount int    `json:"wordcount"`
	Snippet   string `json:"snippet"`
}

type searchQuery struct {
	Query   string       `json:"query"`
	Results []wikiResult `json:"results"`
}

func (w *wikiResult) Clean() {
	// remove all the tags
	exp := regexp.MustCompile(`<[^>]*>`)
	w.Snippet = exp.ReplaceAllString(w.Snippet, "")
}

func getResults(s string) []wikiResult {
	v, err := fastjson.Parse(s)
	if err != nil {
		log.Fatal(err)
	}

	arr := v.GetArray("query", "search")

	var wr wikiResult
	var results []wikiResult

	for _, val := range arr {
		err = json.Unmarshal([]byte(val.String()), &wr)
		if err != nil {
			log.Fatal(err)
		}
		wr.Clean()
		results = append(results, wr)
	}
	return results
}

func jsonify(v interface{}) []byte {
	j, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	return j
}

// SearchHandler handler for search
func SearchHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	queries, ok := r.URL.Query()["q"]
	if !ok || len(queries[0]) < 1 {
		log.Println("Url Param 'q' is missing")
		w.Write([]byte("[]"))
		return
	}
	q := queries[0]

	//check the cache first
	ctx := r.Context()
	present, items := cache.Check(ctx, q)

	var results []byte
	if !present {
		res := search(ctx, string(q))
		sq := searchQuery{Query: q, Results: res}
		cache.Add(ctx, q, sq)
		results = jsonify(sq)
	} else {
		results = jsonify(items)
	}

	w.Write(results)
}

func main() {
	cache.Init(context.Background())
	http.HandleFunc("/search", SearchHandler)
	http.ListenAndServe(":8000", nil)
}
