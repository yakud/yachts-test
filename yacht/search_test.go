package yacht

import (
	"testing"
)

func _search(storageES *StorageES) *Search {
	return NewSearch(storageES.client, storageES)
}

func _loadStorage(storage *StorageES) error {
	model1 := &Model{Id: 1, BuilderName: "builder_1", ModelName: "model_1", OwnerName: "owner_1"}
	if err := storage.Add(model1); err != nil {
		return err
	}

	model2 := &Model{Id: 2, BuilderName: "builder_1", ModelName: "model_2", OwnerName: "owner_1"}
	if err := storage.Add(model2); err != nil {
		return err
	}

	model3 := &Model{Id: 3, BuilderName: "builder_2", ModelName: "model_3", OwnerName: "owner_1"}
	if err := storage.Add(model3); err != nil {
		return err
	}

	model4 := &Model{Id: 4, BuilderName: "super_builder_3", ModelName: "model_4", OwnerName: "owner_1"}
	if err := storage.Add(model4); err != nil {
		return err
	}

	storage.client.Refresh(storage.Index())

	return nil
}

func TestSearch_Search(t *testing.T) {
	storage := _storage()
	defer storage.DeleteIndex()

	storage.CreateIndexIfNotExists()
	_loadStorage(storage)

	search := _search(storage)
	res, total, err := search.Search(0, 10, []SearchFilter{}) // {"builder_name", "b"}
	if err != nil {
		t.Error(err)
	}

	expectedTotal := 4
	if total != expectedTotal || len(res) != expectedTotal {
		t.Errorf("should be %d elements. got: %d", expectedTotal, total)
	}
}

func TestSearch_SearchFilter(t *testing.T) {
	storage := _storage()
	defer storage.DeleteIndex()

	storage.CreateIndexIfNotExists()
	_loadStorage(storage)

	search := _search(storage)

	// builder_name: b
	res, total, err := search.Search(0, 10, []SearchFilter{{"builder_name", "b"}})
	if err != nil {
		t.Error(err)
	}

	expectedTotal := 3
	if total != expectedTotal || len(res) != expectedTotal {
		t.Errorf("should be %d elements. got: %d", expectedTotal, total)
	}

	// builder_name: b
	res, total, err = search.Search(0, 10, []SearchFilter{Filter("builder_name", "builder_1")})
	if err != nil {
		t.Error(err)
	}

	expectedTotal = 2
	if total != expectedTotal || len(res) != expectedTotal {
		t.Errorf("should be %d elements. got: %d", expectedTotal, total)
	}

	// builder_name: builder AND model_name: model
	res, total, err = search.Search(0, 10, []SearchFilter{{"builder_name", "builder"}, {"model_name", "model"}})
	if err != nil {
		t.Error(err)
	}

	expectedTotal = 3
	if total != expectedTotal || len(res) != expectedTotal {
		t.Errorf("should be %d elements. got: %d", expectedTotal, total)
	}

	// model_name: model_2
	res, total, err = search.Search(0, 10, []SearchFilter{{"model_name", "model_2"}})
	if err != nil {
		t.Error(err)
	}

	expectedTotal = 1
	if total != expectedTotal || len(res) != expectedTotal {
		t.Errorf("should be %d elements. got: %d", expectedTotal, total)
	}

	// builder_name: su
	res, total, err = search.Search(0, 10, []SearchFilter{{"builder_name", "su"}})
	if err != nil {
		t.Error(err)
	}

	expectedTotal = 1
	if total != expectedTotal || len(res) != expectedTotal {
		t.Errorf("should be %d elements. got: %d", expectedTotal, total)
	}
}
