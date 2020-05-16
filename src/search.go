package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func search(ctx context.Context, q string) []wikiResult {
	base, err := url.Parse("https://en.wikipedia.org/w/api.php")
	if err != nil {
		log.Fatal(err)
	}

	params := url.Values{"action": {"query"}, "format": {"json"}, "list": {"search"}, "srsearch": {q}}
	base.RawQuery = params.Encode()

	log.Println("Searching for", q)

	//TODO set a context to HTTP request

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
