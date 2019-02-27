package main

import (
	"log"

	"github.com/olivere/elastic"
	"github.com/yakud/yachts-test/yacht"

	"github.com/yakud/yachts-test/api"

	"github.com/gramework/gramework"
)

func main() {
	esClient, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatal(err)
	}

	storageEs := yacht.NewStorageES(esClient)
	suggester := yacht.NewCompletionSuggester(esClient, storageEs)
	search := yacht.NewSearch(esClient, storageEs)

	app := gramework.New()

	app.GET("/search", api.NewSearch(search))
	app.GET("/suggest/builder_name", api.NewSuggest(yacht.BuilderNameSuggestField, suggester))
	app.GET("/suggest/model_name", api.NewSuggest(yacht.ModelNameSuggestField, suggester))

	if err := app.ListenAndServe("127.0.0.1:8087"); err != nil {
		log.Fatal(err)
	}
}
