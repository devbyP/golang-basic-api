package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Server struct {
	Router *chi.Mux
}

func NewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func (s Server) MountHandlers(bh BookHandler) {
	r := s.Router

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           600,
	}))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/book", func(r chi.Router) {
		r.Use(jsonContentType)
		r.Get("/", bh.handleGetBooks)
		r.Post("/", bh.handlePostBook)

		r.Route("/{bookId}", func(r chi.Router) {
			r.Use(bookCtx)
			r.Get("/", bh.handleGetBookByID)
			r.Patch("/", bh.handlePatchBook)
			r.Delete("/", bh.handleDeleteBook)
		})
	})
}

func bookCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bookID := chi.URLParam(r, "bookId")
		ctx := context.WithValue(r.Context(), "bookId", bookID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func jsonContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func Json(w http.ResponseWriter, data any, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

type HTTPError struct {
	Message string `json:"message"`
}

func JsonError(w http.ResponseWriter, mess string, status int) {
	err := HTTPError{Message: mess}
	Json(w, err, status)
}
