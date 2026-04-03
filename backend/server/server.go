package server

import (
	"database/sql"
	"log"
	"net/http"
)

type Server struct {
	db *sql.DB
}

func New(db *sql.DB) *Server {
	return &Server{db: db}
}

func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", s.usersHandler)
	return enableCORS(mux)
}

func (s *Server) Run(port string) {
	addr := ":" + port
	log.Printf("🚀 Server running on %s\n", addr)

	if err := http.ListenAndServe(addr, s.routes()); err != nil {
		log.Fatal(err)
	}
}
