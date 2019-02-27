package api

import (
	"html/template"
	"log"
	"strconv"

	"github.com/gramework/gramework"
	"github.com/yakud/yachts-test/yacht"
)

const defaultSize = 10

type SearchResponse struct {
	TotalRows int           `json:"total"`
	Yachts    []yacht.Model `json:"yachts"`
}

type Search struct {
	field  string
	search *yacht.Search
}

func (t *Search) Handler(ctx *gramework.Context) error {
	from, err := strconv.ParseInt(string(ctx.QueryArgs().Peek("from")), 10, 32)
	if err != nil {
		from = 0
	}
	size, err := strconv.ParseInt(string(ctx.QueryArgs().Peek("size")), 10, 32)
	if err != nil {
		size = defaultSize
	}

	modelName := ctx.QueryArgs().Peek("model_name")
	builderName := ctx.QueryArgs().Peek("builder_name")

	filters := make([]yacht.SearchFilter, 0)
	if len(modelName) > 0 {
		filters = append(filters, yacht.Filter("model_name", string(modelName)))
	}
	if len(builderName) > 0 {
		filters = append(filters, yacht.Filter("builder_name", string(builderName)))
	}

	yachts, _, err := t.search.Search(int(from), int(size), filters)
	if err != nil {
		// @TODO:

	}

	//if bytes.HasPrefix(ctx.Request.Header.ContentType(), []byte("application/json")) {
	//	return SearchResponse{
	//		Yachts:    yachts,
	//		TotalRows: totalRows,
	//	}, nil
	//}

	searchTemplate, err := template.ParseFiles("static/search.html")
	if err != nil {
		log.Fatal(err)
	}

	if err := searchTemplate.Execute(ctx.Response.BodyWriter(), yachts); err != nil {
		return err
	}
	ctx.HTML()

	return nil
}

func NewSearch(search *yacht.Search) *Search {
	return &Search{
		search: search,
	}
}
