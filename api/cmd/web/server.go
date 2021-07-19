package main

import (
	"context"
	"ecommerce/api/core"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Server struct {
	s   core.Service
	idp core.IdentityProvider
	r   *mux.Router
}

func NewServer(s core.Service, idp core.IdentityProvider) *Server {
	return &Server{
		s:   s,
		idp: idp,
	}
}

func (s *Server) Init() {
	sr := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: s.Handle(),
	}

	fmt.Println("Server started...")
	if err := sr.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func (s *Server) Handle() http.Handler {
	r := mux.NewRouter()

	sr := r.PathPrefix("/api/v1").Subrouter()

	sr.HandleFunc("/login", s.idp.Authenticate)
	sr.HandleFunc("/authorize", s.idp.Authorize)

	s.HandleAuthenticated(sr, "/categories/popular", s.HandleGetPopularCategories)
	s.HandleAuthenticated(sr, "/categories/filter", s.HandleGetFilterCategories)
	s.HandleAuthenticated(sr, "/products/trend", s.HandleProductsTrend)
	s.HandleAuthenticated(sr, "/products/last_seen", s.HandleProductsLastSeen)

	return r
}

func (s *Server) HandleAuthenticated(r *mux.Router, path string, h http.HandlerFunc) {
	r.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		auth := "Bearer APP_USR-2975116266796832-071902-779de2392e8be8f9d7eb7cae25bc029e-276682175" //r.Header.Get("Authorization")

		if auth == "" {
			w.Write([]byte("Authorization header missing"))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := auth[len("Bearer "):]

		if token == "" {
			w.Write([]byte("Bearer token header missing"))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "token", token))

		h(w, r)
	})
}

func (s *Server) HandleProductsTrend(w http.ResponseWriter, r *http.Request) {
	res := s.s.GetTrendingProducts(r.Context())

	s.Respond(w, r, res, http.StatusOK)
}

func (s *Server) HandleProductsLastSeen(w http.ResponseWriter, r *http.Request) {
	res := s.s.GetLastSeenProducts(r.Context())

	s.Respond(w, r, res, http.StatusOK)
}

func (s *Server) HandleGetPopularCategories(w http.ResponseWriter, r *http.Request) {
	res := s.s.GetPopularCategories(r.Context())

	s.Respond(w, r, res, http.StatusOK)
}

func (s *Server) HandleGetFilterCategories(w http.ResponseWriter, r *http.Request) {
	res := s.s.GetFilterCategories(r.Context())

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
