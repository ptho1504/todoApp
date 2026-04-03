package main

import (
	"backend/config"
	"backend/db"
	"backend/server"
)

func main() {
	cfg := config.Load()

	database := db.New(cfg)
	defer database.Close()

	srv := server.New(database)
	srv.Run(cfg.Port)
}
