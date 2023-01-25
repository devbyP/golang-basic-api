package main

import "net/http"

func pagination(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// implement pagination by get page query
		next.ServeHTTP(w, r)
	})
}
