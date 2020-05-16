package main

import (
	"testing"
)

func TestCleanWikiResult(t *testing.T) {
	wr := wikiResult{Snippet: "The <span class=\"searchmatch\">Perito</span> <span class=\"searchmatch\">Moreno</span> Glacier (Spanish: Glaciar <span class=\"searchmatch\">Perito</span> <span class=\"searchmatch\">Moreno</span>) is a glacier located in the Los Glaciares National Park in southwest Santa Cruz Province, Argentina"}
	wr.Clean()
	cleaned := "The Perito Moreno Glacier (Spanish: Glaciar Perito Moreno) is a glacier located in the Los Glaciares National Park in southwest Santa Cruz Province, Argentina"
	if wr.Snippet != cleaned {
		t.Errorf("Could not clean snippet, Expected %s, Got %s", cleaned, wr.Snippet)
	}
}

// func TestSearch(t *testing.T) {
// 	results := search("Rome")
// 	if len(results) != 10 {
// 		t.Errorf("Expected 10 results, got %d", len(results))
// 	}
// }

func TestGetResults(t *testing.T) {
	apiResponse := `{"batchcomplete":"","continue":{"sroffset":10,"continue":"-||"},"query":{"searchinfo":{"totalhits":253},"search":[{"ns":0,"title":"Perito Moreno Glacier","pageid":1007154,"size":16827,"wordcount":2185,"snippet":"The <span class=\"searchmatch\">Perito</span> <span class=\"searchmatch\">Moreno</span> Glacier (Spanish: Glaciar <span class=\"searchmatch\">Perito</span> <span class=\"searchmatch\">Moreno</span>) is a glacier located in the Los Glaciares National Park in southwest Santa Cruz Province, Argentina","timestamp":"2020-03-24T03:50:37Z"},{"ns":0,"title":"Perito Moreno","pageid":1727540,"size":373,"wordcount":53,"snippet":"<span class=\"searchmatch\">Perito</span> <span class=\"searchmatch\">Moreno</span> may refer to: Francisco <span class=\"searchmatch\">Moreno</span>, Argentine explorer and scientist <span class=\"searchmatch\">Perito</span> <span class=\"searchmatch\">Moreno</span> Glacier, at the Los Glaciares National Park in Santa Cruz","timestamp":"2013-11-07T14:18:22Z"},{"ns":0,"title":"Perito Moreno, Santa Cruz","pageid":24088973,"size":8945,"wordcount":418,"snippet":"<span class=\"searchmatch\">Perito</span> <span class=\"searchmatch\">Moreno</span> is a town in the northwest of Santa Cruz Province, Argentina, 25\u00a0km east of Lake Buenos Aires. It should not be confused with the Perito","timestamp":"2020-05-04T02:02:46Z"},{"ns":0,"title":"Perito Moreno Airport","pageid":35278572,"size":3352,"wordcount":155,"snippet":"<span class=\"searchmatch\">Perito</span> <span class=\"searchmatch\">Moreno</span> Airport (IATA: PMQ, ICAO: SAWP) is an airport serving <span class=\"searchmatch\">Perito</span> <span class=\"searchmatch\">Moreno</span>, a town in the Santa Cruz Province of Argentina. The airport is 7 kilometres","timestamp":"2020-03-09T02:19:52Z"},{"ns":0,"title":"Francisco Moreno","pageid":2040206,"size":10465,"wordcount":1119,"snippet":"Pascasio <span class=\"searchmatch\">Moreno</span> (May 31, 1852 \u2013 November 22, 1919) was a prominent explorer and academic in Argentina, where he is usually referred to as <span class=\"searchmatch\">Perito</span> <span class=\"searchmatch\">Moreno</span> (perito","timestamp":"2020-03-19T12:11:08Z"},{"ns":0,"title":"Perito Moreno National Park","pageid":22272600,"size":4936,"wordcount":509,"snippet":"<span class=\"searchmatch\">Perito</span> <span class=\"searchmatch\">Moreno</span> National Park (Spanish: Parque Nacional <span class=\"searchmatch\">Perito</span> <span class=\"searchmatch\">Moreno</span>) is a national park in Argentina. It is located in the western region of Santa Cruz","timestamp":"2020-04-26T02:48:38Z"},{"ns":0,"title":"Cueva de las Manos","pageid":2420596,"size":6351,"wordcount":591,"snippet":"province of Santa Cruz, Argentina, 163\u00a0km (101\u00a0mi) south of the town of <span class=\"searchmatch\">Perito</span> <span class=\"searchmatch\">Moreno</span>. It is famous for (and gets its name from) the paintings of hands. The","timestamp":"2020-02-20T22:54:11Z"},{"ns":0,"title":"Traveler (South Korean TV series)","pageid":60057900,"size":28923,"wordcount":476,"snippet":"than their changing attire. They will go to places like Iguazu Falls, <span class=\"searchmatch\">Perito</span> <span class=\"searchmatch\">Moreno</span> Glacier and fighting against the wind on the road across Patagonia.","timestamp":"2020-05-09T06:32:48Z"},{"ns":0,"title":"Santa Cruz Province, Argentina","pageid":520495,"size":37770,"wordcount":2778,"snippet":"of <span class=\"searchmatch\">Perito</span> <span class=\"searchmatch\">Moreno</span> allows the few tourists who venture to this point to see the prehistoric wall paintings in the caves near the Pinturas River. <span class=\"searchmatch\">Perito</span> Moreno","timestamp":"2020-03-25T04:24:38Z"},{"ns":0,"title":"El Calafate","pageid":4371221,"size":13822,"wordcount":782,"snippet":"visit different parts of the Los Glaciares National Park, including the <span class=\"searchmatch\">Perito</span> <span class=\"searchmatch\">Moreno</span> Glacier and the Cerro Chalt\u00e9n and Cerro Torre. The history of El Calafate","timestamp":"2020-04-20T21:08:12Z"}]}}`
	results := getResults(apiResponse)
	if len(results) != 10 {
		t.Errorf("Expected 10 results, got %d", len(results))
	}
}
