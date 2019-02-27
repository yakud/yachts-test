package api

import (
	"time"

	"github.com/gramework/gramework"
	"github.com/yakud/yachts-test/yacht"
)

type suggestResponse struct {
	Suggest []string `json:"suggest"`
	Stats   string   `json:"stats"`
}

type Suggest struct {
	field     string
	suggester *yacht.CompletionSuggester
}

func (t *Suggest) Handler(ctx *gramework.Context) (interface{}, error) {
	text := ctx.QueryArgs().Peek("q")

	now := time.Now()
	suggests, err := t.suggester.Suggest(t.field, string(text))
	if err != nil {
		return nil, err
	}

	return suggestResponse{
		Suggest: suggests,
		Stats:   time.Now().Sub(now).String(),
	}, nil
}

func NewSuggest(field string, suggester *yacht.CompletionSuggester) *Suggest {
	return &Suggest{
		field:     field,
		suggester: suggester,
	}
}
