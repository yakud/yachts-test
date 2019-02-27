package yacht

import (
	"testing"

	"github.com/yakud/yachts-test/gds"
)

func _client() *gds.Client {
	return gds.NewClient(&gds.ClientConfig{
		Entrypoint: "http://ws.nausys.com/",
		Login:      "rest83@TTTTT",
		Password:   "Rest59Tb",
	})
}

func _gdsLoader() *GDSLoader {
	return NewGDSLoader(_client())
}

func TestGDSLoader_LoadTo(t *testing.T) {
	storage := _storage()
	defer storage.DeleteIndex()

	if err := storage.CreateIndexIfNotExists(); err != nil {
		t.Error(err)
	}

	loader := _gdsLoader()
	if err := loader.LoadTo(storage); err != nil {
		t.Error(err)
	}

	_, total, err := _search(storage).Search(0, 100, nil)
	if err != nil {
		t.Error(err)
	}

	if total <= 0 {
		t.Errorf("no one rows loaded from gds")
	}
}
