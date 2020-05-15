package main

import (
	"encoding/json"
	"fmt"
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
	exp := regexp.MustCompile(`<.*>`)
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

func main() {

	base, err := url.Parse("https://en.wikipedia.org/w/api.php")
	if err != nil {
		return
	}

	q := "Perito Moreno"

	// Query params
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

	results := getResults(string(body))

	j, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))

}
