package src

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Register(s Service) http.Handler {
	r := mux.NewRouter()

	sr := r.PathPrefix("/api/v1").Subrouter()

	sr.HandleFunc("/categories/popular", HandleGetPopularCategories(s))
	sr.HandleFunc("/categories/filter", HandleGetFilterCategories(s))
	sr.HandleFunc("/products/trend", HandleProductsTrend(s))
	sr.HandleFunc("/products/last_seen", HandleProductsLastSeen(s))

	return r
}
