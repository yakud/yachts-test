package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/montanaflynn/stats"

	"github.com/olivere/elastic"
	"github.com/yakud/yachts-test/yacht"
)

func main() {
	wg := &sync.WaitGroup{}

	var count uint64
	var durations []time.Duration
	durationMutex := &sync.Mutex{}
	ctx := context.Background()

	workers := 4

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			esClient, err := elastic.NewClient(
				elastic.SetURL("http://0.0.0.0:9200"),
				//elastic.SetURL("http://es-yachts:9200"),
				elastic.SetSniff(false),
			)
			if err != nil {
				log.Fatal(err)
			}

			storageEs := yacht.NewStorageES(esClient)
			search := yacht.NewSearch(esClient, storageEs)
			//suggester := yacht.NewCompletionSuggester(esClient, storageEs)

			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return

				default:
					start := time.Now()
					if _, _, err := search.Search(0, 10, nil); err != nil {
						log.Fatal(err)
					}
					//if _, err := suggester.Suggest(yacht.SuggestFieldBuilderName, "b"); err != nil {
					//	log.Fatal(err)
					//}
					atomic.AddUint64(&count, 1)

					durationMutex.Lock()
					durations = append(durations, time.Now().Sub(start))
					durationMutex.Unlock()
				}
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return

			case <-time.After(time.Second):
				meanDurations := make([]float64, 0)
				for _, d := range durations {
					meanDurations = append(meanDurations, float64(d))
				}

				mean, _ := stats.Mean(meanDurations)

				fmt.Print(workers, "\t", atomic.LoadUint64(&count), "\t", time.Duration(mean), "\n")
				atomic.StoreUint64(&count, 0)
			}
		}
	}()

	wg.Wait()
}
