package main

import (
	"fmt"
	"log"
	"time"

	"github.com/olivere/elastic"

	"github.com/yakud/yachts-test/yacht"
)

func main() {
	esClient, err := elastic.NewClient(
		elastic.SetURL("http://0.0.0.0:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatal(err)
	}

	storageEs := yacht.NewStorageES(esClient)
	search := yacht.NewSearch(esClient, storageEs)

	start := time.Now()
	for i := 0; i < 1; i++ {
		_, _, err := search.Search(
			0,
			5,
			[]yacht.SearchFilter{
				yacht.Filter("builder_name", "Bava"),
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("1 search for: ", time.Now().Sub(start))
}
