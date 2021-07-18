package src

import (
	"ecommerce/api/src/core"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Server struct {
	s core.Service
}

func NewServer(s core.Service) *Server {
	return &Server{
		s: s,
	}
}

func (s *Server) Init() {
	sr := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: s.Handle(),
	}

	fmt.Println("Server started...")
	sr.ListenAndServe()
}

func (s *Server) Handle() http.Handler {
	r := mux.NewRouter()

	sr := r.PathPrefix("/api/v1").Subrouter()

	sr.HandleFunc("/categories/popular", s.HandleGetPopularCategories)
	sr.HandleFunc("/categories/filter", s.HandleGetFilterCategories)
	sr.HandleFunc("/products/trend", s.HandleProductsTrend)
	sr.HandleFunc("/products/last_seen", s.HandleProductsLastSeen)

	return r
}

func (s *Server) HandleAuthenticated(h http.HandlerFunc) {

}

func (s *Server) HandleProductsTrend(w http.ResponseWriter, r *http.Request) {
	res := s.s.GetTrendingProducts()

	s.Respond(w, r, res, http.StatusOK)
}

func (s *Server) HandleProductsLastSeen(w http.ResponseWriter, r *http.Request) {
	res := s.s.GetLastSeenProducts()

	s.Respond(w, r, res, http.StatusOK)
}

func (s *Server) HandleGetPopularCategories(w http.ResponseWriter, r *http.Request) {
	res := s.s.GetPopularCategories()

	s.Respond(w, r, res, http.StatusOK)
}

func (s *Server) HandleGetFilterCategories(w http.ResponseWriter, r *http.Request) {
	res := s.s.GetFilterCategories()

	s.Respond(w, r, res, http.StatusOK)
}

func (s *Server) Respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.WriteHeader(status)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			fmt.Println(err)
		}
	}
}
