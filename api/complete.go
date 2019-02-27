package api

import (
	"github.com/gramework/gramework"
	"github.com/yakud/yachts-test/yacht"
)

type CompleteSuggest struct {
	field     string
	suggester *yacht.CompletionSuggester
}

func (t *CompleteSuggest) Handler(ctx *gramework.Context) (interface{}, error) {
	text := ctx.QueryArgs().Peek("text")

	suggests, err := t.suggester.Suggest(t.field, string(text))
	if err != nil {
		// @TODO:
	}

	return suggests, nil
}

func NewCompleteSuggest(field string, suggester *yacht.CompletionSuggester) *CompleteSuggest {
	return &CompleteSuggest{
		field:     field,
		suggester: suggester,
	}
}
