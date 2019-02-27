package yacht

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	"github.com/olivere/elastic"
)

const defaultIndex = "yachts"
const defaultType = "_doc"

const (
	BuilderNameSuggestField = "builder_name_suggest"
	ModelNameSuggestField   = "model_name_suggest"
)

type StorageESRow struct {
	Model

	BuilderNameSuggest string `json:"builder_name_suggest"`
	ModelNameSuggest   string `json:"model_name_suggest"`
}

type StorageES struct {
	indexES string
	typeES  string
	client  *elastic.Client
}

func (t *StorageES) Add(yacht *Model) error {
	fmt.Printf("%+v\n", yacht)

	row := StorageESRow{
		Model:              *yacht,
		BuilderNameSuggest: yacht.BuilderName,
		ModelNameSuggest:   yacht.ModelName,
	}

	// @TODO: bulk upload
	_, err := t.client.Index().
		Index(t.indexES).
		Type(t.typeES).
		Id(strconv.FormatInt(int64(row.Id), 10)).
		BodyJson(row).
		Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (t *StorageES) CreateIndexIfNotExists() error {
	exists, err := t.client.IndexExists(t.Index()).Do(context.Background())
	if err != nil {
		return err
	}

	if !exists {
		createIndex, err := t.client.CreateIndex(t.Index()).
			BodyString(t.GetMapping()).
			Do(context.Background())

		if err != nil {
			return err
		}
		if !createIndex.Acknowledged {
			return errors.New("index is not acknowledged")
		}
	}

	return nil
}

func (t *StorageES) DeleteIndex() error {
	_, err := t.client.DeleteIndex(t.indexES).Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (t *StorageES) Index() string {
	return defaultIndex
}

func (t *StorageES) Type() string {
	return defaultType
}

func (t *StorageES) GetMapping() string {
	return `{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"_doc":{
			"properties":{
				"id":{
					"type":"long"
				},
				"builder_name":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"model_name":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"owner_name":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"reservation_from":{
					"type":"date"
				},
				"reservation_to":{
					"type":"date"
				},
				"builder_name_suggest":{
					"type":"completion"
				},
				"model_name_suggest":{
					"type":"completion"
				}
			}
		}
	}
}`
}

func NewStorageES(client *elastic.Client) *StorageES {
	return &StorageES{
		indexES: defaultIndex,
		typeES:  defaultType,
		client:  client,
	}
}
