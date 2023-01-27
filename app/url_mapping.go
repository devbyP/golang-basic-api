package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AppMapping struct {
    app *AppRenderer
}

func NewApp() *AppMapping {
    return &AppMapping{
        app: newAppRenderer(),
    }
}

func (am *AppMapping) Mapping() http.Handler {
    r := chi.NewRouter()

    r.Get("/", am.app.renderHome)
    r.Get("/collection", am.app.renderCollection)

    return r
}
