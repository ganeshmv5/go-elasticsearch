// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/sort.asciidoc#L392>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "sort" : [
//         {
//             "_geo_distance" : {
//                 "pin.location" : [-70, 40],
//                 "order" : "asc",
//                 "unit" : "km",
//                 "mode" : "min",
//                 "distance_type" : "arc",
//                 "ignore_unmapped": true
//             }
//         }
//     ],
//     "query" : {
//         "term" : { "user" : "kimchy" }
//     }
// }
// --------------------------------------------------------------------------------

func Test_search_request_sort_d17269bb80fb63ec0bf37d219e003dcb(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:d17269bb80fb63ec0bf37d219e003dcb[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "sort": [
		    {
		      "_geo_distance": {
		        "pin.location": [
		          -70,
		          40
		        ],
		        "order": "asc",
		        "unit": "km",
		        "mode": "min",
		        "distance_type": "arc",
		        "ignore_unmapped": true
		      }
		    }
		  ],
		  "query": {
		    "term": {
		      "user": "kimchy"
		    }
		  }
		}`)),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:d17269bb80fb63ec0bf37d219e003dcb[]
}
