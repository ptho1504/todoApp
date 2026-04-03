package server

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s *Server) usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		rows, err := s.db.Query("SELECT id, name FROM users")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Name); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			users = append(users, u)
		}

		json.NewEncoder(w).Encode(users)

	case "POST":
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		_, err := s.db.Exec("INSERT INTO users(name) VALUES(?)", u.Name)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(http.StatusCreated)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
