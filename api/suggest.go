package api

import (
	"github.com/gramework/gramework"
	"github.com/yakud/yachts-test/yacht"
)

type Suggest struct {
	field     string
	suggester *yacht.CompletionSuggester
}

func (t *Suggest) Handler(ctx *gramework.Context) (interface{}, error) {
	text := ctx.QueryArgs().Peek("q")

	suggests, err := t.suggester.Suggest(t.field, string(text))
	if err != nil {
		// @TODO:
	}

	return suggests, nil
}

func NewSuggest(field string, suggester *yacht.CompletionSuggester) *Suggest {
	return &Suggest{
		field:     field,
		suggester: suggester,
	}
}
