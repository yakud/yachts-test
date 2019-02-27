package yacht

import "testing"

func _suggester(storageES *StorageES) *CompletionSuggester {
	return NewCompletionSuggester(storageES.client, storageES)
}

func TestCompletionSuggester_Suggest(t *testing.T) {
	storage := _storage()
	defer storage.DeleteIndex()

	storage.CreateIndexIfNotExists()
	_loadStorage(storage)

	suggester := _suggester(storage)

	// builder_name: bu
	res, err := suggester.Suggest(SuggestFieldBuilderName, "bu")
	if err != nil {
		t.Error()
	}

	expectedTotal := 2
	if len(res) != expectedTotal {
		t.Errorf("should be %d elements. got: %d", expectedTotal, len(res))
	}

	// builder_name: su
	res, err = suggester.Suggest(SuggestFieldBuilderName, "su")
	if err != nil {
		t.Error()
	}

	expectedTotal = 1
	if len(res) != expectedTotal {
		t.Errorf("should be %d elements. got: %d", expectedTotal, len(res))
	}

	// fuzz example:
	// builder_name: builll
	res, err = suggester.Suggest(SuggestFieldBuilderName, "builll")
	if err != nil {
		t.Error()
	}

	expectedTotal = 2
	if len(res) != expectedTotal {
		t.Errorf("should be %d elements. got: %d", expectedTotal, len(res))
	}

	// model_name: mod
	res, err = suggester.Suggest(SuggestFieldModelName, "mod")
	if err != nil {
		t.Error()
	}

	expectedTotal = 4
	if len(res) != expectedTotal {
		t.Errorf("should be %d elements. got: %d", expectedTotal, len(res))
	}
}
