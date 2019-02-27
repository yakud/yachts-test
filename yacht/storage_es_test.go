package yacht

import (
	"log"
	"testing"

	"github.com/olivere/elastic"
)

func _storage() *StorageES {
	esClient, err := elastic.NewClient(
		elastic.SetURL("http://0.0.0.0:9200"),
		elastic.SetSniff(false),
		//elastic.SetTraceLog(log.New(os.Stderr, "[[TRACE LOG]] ", 0)),
	)
	if err != nil {
		log.Fatal(err)
	}

	storage := NewStorageES(esClient)
	storage.indexES = "test_index"
	return storage
}

func TestStorageES_CreateIndexIfNotExists(t *testing.T) {
	storage := _storage()
	defer storage.DeleteIndex()

	if err := storage.CreateIndexIfNotExists(); err != nil {
		t.Error(err)
	}
}

func TestStorageES_DeleteIndex(t *testing.T) {
	storage := _storage()
	defer storage.DeleteIndex()

	if err := storage.CreateIndexIfNotExists(); err != nil {
		t.Error(err)
	}

	if err := storage.DeleteIndex(); err != nil {
		t.Error(err)
	}
}

func TestStorageES_Add(t *testing.T) {
	storage := _storage()
	defer storage.DeleteIndex()

	if err := storage.CreateIndexIfNotExists(); err != nil {
		t.Error(err)
	}

	err := storage.Add(&Model{
		Id:          1,
		BuilderName: "builder_1",
		ModelName:   "model_1",
		OwnerName:   "owner_1",
	})
	if err != nil {
		t.Error(err)
	}
}
