package src

import (
	"fmt"
	"net/http"
	"os"
)

type Server struct {
}

func (s *Server) Init() {
	os.Setenv("PORT", ":3001")

	r := &ServiceImp{}
	r.Init()

	sr := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: Register(r),
	}

	fmt.Println("Server started...")
	sr.ListenAndServe()
}
