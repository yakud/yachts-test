package main

import (
	"log"

	"github.com/olivere/elastic"

	"github.com/yakud/yachts-test/yacht"

	"github.com/yakud/yachts-test/gds"
)

func main() {
	clientConfig := &gds.ClientConfig{
		Entrypoint: "http://ws.nausys.com/",
		Login:      "rest83@TTTTT",
		Password:   "Rest59Tb",
	}
	esClient, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := gds.NewClient(clientConfig)

	storageEs := yacht.NewStorageES(esClient)
	if err := storageEs.DeleteIndex(); err != nil {
		log.Fatal(err)
	}
	if err := storageEs.CreateIndexIfNotExists(); err != nil {
		log.Fatal(err)
	}

	loader := yacht.NewGDSLoader(client)

	if err := loader.LoadTo(storageEs); err != nil {
		log.Fatal(err)
	}
}
