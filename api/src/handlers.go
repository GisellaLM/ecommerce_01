package src

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleProductsTrend(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := s.GetTrendingProducts()

		Respond(w, r, res, http.StatusOK)
	}
}

func HandleProductsLastSeen(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := s.GetLastSeenProducts()

		Respond(w, r, res, http.StatusOK)
	}
}

func HandleGetPopularCategories(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := s.GetPopularCategories()

		Respond(w, r, res, http.StatusOK)
	}
}

func HandleGetFilterCategories(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := s.GetFilterCategories()

		Respond(w, r, res, http.StatusOK)
	}
}

func Respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.WriteHeader(status)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			fmt.Println(err)
		}
	}
}
