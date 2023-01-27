package app

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var templateDir = "./app/templates/"

func mustNewTemplate() *template.Template {
	return template.Must(template.ParseGlob(filepath.Join(templateDir, "*.html")))
}

type AppRenderer struct {
	temp *template.Template
}

func newAppRenderer() *AppRenderer {
	return &AppRenderer{
		temp: mustNewTemplate(),
	}
}

func (ar *AppRenderer) renderHome(w http.ResponseWriter, r *http.Request) {
	ar.temp.ExecuteTemplate(w, "index.html", nil)
}

type BookRender struct {
	ID           int
	Title        string
	AuthorName   string
	ReleasedDate string
}

func (ar *AppRenderer) renderCollection(w http.ResponseWriter, r *http.Request) {
    testBooks := make([]*BookRender, 1)
    testBooks[0] = &BookRender{
        ID: 1,
        Title: "The Go programming language",
        AuthorName: "test",
        ReleasedDate: "Yesterday",
    }
	ar.temp.ExecuteTemplate(w, "collection.html", testBooks)
}
