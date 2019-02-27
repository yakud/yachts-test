package api

import (
	"html/template"
	"strconv"
	"time"

	"github.com/gramework/gramework"
	"github.com/yakud/yachts-test/yacht"
)

const defaultSize = 500

type searchTemplateData struct {
	Yachts []yachtsTemplate
	Stats  string
}

type yachtsTemplate struct {
	Id             uint64
	BuilderName    string
	ModelName      string
	OwnerName      string
	IsAvailableNow bool
	AvailableFrom  string
}

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

	searchTemplate, err := template.ParseFiles("static/search.html")
	if err != nil {
		return err
	}

	now := time.Now()

	yachts, _, err := t.search.Search(int(from), int(size), filters)
	if err != nil {
		return err
	}

	templateYachts := make([]yachtsTemplate, 0, len(yachts))
	for _, y := range yachts {
		isAvailable := now.Before(y.ReservationFrom) || now.After(y.ReservationTo)

		var availableFrom string
		if !isAvailable {
			availableFrom = y.ReservationTo.Format("2006-01-02 15:04")
		}

		templateYachts = append(templateYachts, yachtsTemplate{
			Id: y.Id,

			BuilderName:    y.BuilderName,
			ModelName:      y.ModelName,
			OwnerName:      y.OwnerName,
			IsAvailableNow: isAvailable,
			AvailableFrom:  availableFrom,
		})
	}

	searchStats := time.Now().Sub(now).String()

	templateData := searchTemplateData{
		Yachts: templateYachts,
		Stats:  searchStats,
	}

	if err := searchTemplate.Execute(ctx.Response.BodyWriter(), templateData); err != nil {
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
