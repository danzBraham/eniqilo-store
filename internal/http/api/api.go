package api

import (
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type APIServer struct {
	Addr string
	DB   *pgxpool.Pool
}

func NewAPIServer(addr string, db *pgxpool.Pool) *APIServer {
	return &APIServer{
		Addr: addr,
		DB:   db,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	server := http.Server{
		Addr:    s.Addr,
		Handler: v1,
	}

	log.Printf("Server listening on %s\n", s.Addr)
	return server.ListenAndServe()
}
