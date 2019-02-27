package api

import (
	"html/template"
	"log"

	"github.com/gramework/gramework"
)

type Index struct {
	template *template.Template
}

func (t *Index) Handler(ctx *gramework.Context) error {
	indexTemplate, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Fatal(err)
	}

	if err := indexTemplate.Execute(ctx.Response.BodyWriter(), nil); err != nil {
		return err
	}
	ctx.HTML()

	return nil
}

func NewIndex(template *template.Template) *Index {
	return &Index{
		template: template,
	}
}
