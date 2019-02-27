package yacht

import (
	"context"

	"github.com/olivere/elastic"
)

const (
	SuggestFieldBuilderName = "builder_name_suggest"
	SuggestFieldModelName   = "model_name_suggest"
)

type CompletionSuggester struct {
	client    *elastic.Client
	storageES *StorageES
}

func (t *CompletionSuggester) Suggest(field, q string) ([]string, error) {
	suggester := elastic.NewCompletionSuggester("suggester")
	suggester.Text(q).
		Field(field).
		SkipDuplicates(true).
		FuzzyOptions(
			elastic.NewFuzzyCompletionSuggesterOptions().EditDistance(2),
		)

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
			//fmt.Printf("%s\n", string(*opt.Source))
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
