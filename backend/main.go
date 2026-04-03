package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var db *sql.DB

func createDB() string {
	dbHost := "testdb.ctcoe8e0aiiq.ap-southeast-1.rds.amazonaws.com:3306"
	dbUser := "root"
	dbPass := "password"
	dbName := "testdb"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		dbUser, dbPass, dbHost, dbName,
	)
	return dsn
}

func main() {
	var err error
	db, err = sql.Open("mysql", createDB())
	log.Print(createDB())
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/users", usersHandler)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", enableCors(http.DefaultServeMux))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rows, _ := db.Query("SELECT id, name FROM users")
		defer rows.Close()

		var users []User
		for rows.Next() {
			var u User
			rows.Scan(&u.ID, &u.Name)
			users = append(users, u)
		}

		json.NewEncoder(w).Encode(users)

	case "POST":
		var u User
		json.NewDecoder(r.Body).Decode(&u)

		_, err := db.Exec("INSERT INTO users(name) VALUES(?)", u.Name)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func enableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
