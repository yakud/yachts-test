package yacht

import (
	"context"

	"github.com/olivere/elastic"
)

type CompletionSuggester struct {
	client    *elastic.Client
	storageES *StorageES
}

func (t *CompletionSuggester) Suggest(field, value string) ([]string, error) {
	suggester := elastic.NewCompletionSuggester("suggester")
	suggester.Prefix(value)
	suggester.Field(field).SkipDuplicates(true)

	res, err := t.client.Search().
		Index(t.storageES.Index()).
		Type(t.storageES.Type()).
		Suggester(suggester).
		Size(0). // we dont' want the hits, just the suggestions
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	resSuggest := make([]string, 0)

	suggestions, found := res.Suggest["suggester"]
	if !found {
		return resSuggest, nil
	}

	for _, suggestion := range suggestions {
		for _, opt := range suggestion.Options {
			resSuggest = append(resSuggest, opt.Text)
		}
	}

	return resSuggest, nil
}

func NewCompletionSuggester(client *elastic.Client, storageES *StorageES) *CompletionSuggester {
	return &CompletionSuggester{
		client:    client,
		storageES: storageES,
	}
}
