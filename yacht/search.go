package yacht

import (
	"context"
	"reflect"

	"github.com/olivere/elastic"
)

type Search struct {
	client    *elastic.Client
	storageES *StorageES
}

func (t *Search) Search(from, size int, filters []SearchFilter) ([]StorageESRow, int, error) {
	// Search with a term query

	queries := make([]elastic.Query, 0, len(filters))
	for _, filter := range filters {
		queries = append(queries, elastic.NewMatchPhrasePrefixQuery(filter.Field, filter.Value))
	}

	query := elastic.NewBoolQuery()
	if len(queries) > 0 {
		query.Must(queries...)
	}

	res, err := t.client.Search().
		Index(t.storageES.Index()).
		Type(t.storageES.Type()).
		Query(query).
		Size(size).
		Pretty(true).
		Sort("model_name", true).
		Do(context.Background())
	if err != nil {
		return nil, 0, err
	}

	var row StorageESRow
	var rows []StorageESRow
	for _, item := range res.Each(reflect.TypeOf(row)) {
		if row, ok := item.(StorageESRow); ok {
			rows = append(rows, row)
		}
	}

	return rows, int(res.Hits.TotalHits), nil
}

func NewSearch(client *elastic.Client, storageES *StorageES) *Search {
	return &Search{
		client:    client,
		storageES: storageES,
	}
}