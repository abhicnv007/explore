package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func search(q string) []wikiResult {
	base, err := url.Parse("https://en.wikipedia.org/w/api.php")
	if err != nil {
		log.Fatal(err)
	}

	params := url.Values{"action": {"query"}, "format": {"json"}, "list": {"search"}, "srsearch": {q}}
	base.RawQuery = params.Encode()

	log.Println("Searching for", q)

	resp, err := http.Get(base.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return getResults(string(body))
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
	queries, ok := r.URL.Query()["q"]
	if !ok || len(queries[0]) < 1 {
		log.Println("Url Param 'q' is missing")
		return
	}
	q := queries[0]
	results := search(string(q))
	w.Write(jsonify(results))
}

func main() {
	http.HandleFunc("/search", SearchHandler)
	http.ListenAndServe(":8000", nil)
}
