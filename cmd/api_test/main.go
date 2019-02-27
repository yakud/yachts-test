package main

import (
	"log"

	"github.com/olivere/elastic"

	"github.com/yakud/yachts-test/yacht"

	"github.com/yakud/yachts-test/gds"
)

func main() {
	apiConfig := &gds.ApiConfig{
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

	api := gds.NewApi(apiConfig)

	storageEs := yacht.NewStorageES(esClient)
	if err := storageEs.DeleteIndex(); err != nil {
		log.Fatal(err)
	}
	if err := storageEs.CreateIndexIfNotExists(); err != nil {
		log.Fatal(err)
	}

	loader := yacht.NewGDSLoader(api)

	if err := loader.LoadTo(storageEs); err != nil {
		log.Fatal(err)
	}

	////////////////////////////////////////////////////////////////////////////////////
	//fmt.Println("Yacht free:", time.Now().Format("02.01.2006"))
	//free, err := api.FreeYachts(&gds.RestFreeYachtsRequest{
	//	PeriodFrom: time.Now().Format("02.01.2006"),
	//	PeriodTo:   time.Now().Add(time.Hour * 24 * 7).Format("02.01.2006"),
	//	Yachts:     allYachts,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, yacht := range free.FreeYachts {
	//	fmt.Printf("%+v\n", yacht)
	//}
}
