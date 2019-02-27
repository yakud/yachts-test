package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/yakud/yachts-test/gds"

	"github.com/olivere/elastic"
	"github.com/yakud/yachts-test/yacht"

	"github.com/yakud/yachts-test/api"

	"github.com/gramework/gramework"
)

func main() {
	clientConfig := &gds.ClientConfig{
		Entrypoint: "http://ws.nausys.com/",
		Login:      "rest83@TTTTT",
		Password:   "Rest59Tb",
	}
	esClient, err := elastic.NewClient(
		//elastic.SetURL("http://0.0.0.0:9200"),
		elastic.SetURL("http://es-yachts:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatal(err)
	}

	gdsClient := gds.NewClient(clientConfig)
	loader := yacht.NewGDSLoader(gdsClient)

	storageEs := yacht.NewStorageES(esClient)
	suggester := yacht.NewCompletionSuggester(esClient, storageEs)
	search := yacht.NewSearch(esClient, storageEs)

	app := gramework.New()

	indexTemplate, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Fatal(err)
	}

	app.GET("/static/style/*static", app.ServeDirNoCache("./"))
	app.GET("/static/js/*static", app.ServeDirNoCache("./"))

	app.GET("/", api.NewIndex(indexTemplate))
	app.GET("/search", api.NewSearch(search))
	app.GET("/suggest/builder_name", api.NewSuggest(yacht.SuggestFieldBuilderName, suggester))
	app.GET("/suggest/model_name", api.NewSuggest(yacht.SuggestFieldModelName, suggester))

	// graceful stop
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v", sig)
		app.Shutdown()
	}()

	go func() {
		storageEs.DeleteIndex()
		if err := storageEs.CreateIndexIfNotExists(); err != nil {
			log.Fatal(err)
		}

		if err := loader.LoadTo(storageEs); err != nil {
			log.Fatal(err)
		}
	}()

	if err := app.ListenAndServe("0.0.0.0:80"); err != nil {
		log.Fatal(err)
	}
}
