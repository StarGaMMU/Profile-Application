package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

// User represents a user profile in the database.
type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Bio            string `json:"bio,omitempty"`             // Omitempty means it won't be sent if empty
	ProfilePicture string `json:"profile_picture,omitempty"` // Omitempty for optional field
	CreatedAt      string `json:"created_at"`
}

// getUsers handles GET requests to fetch all users.
func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	rows, err := db.Query("SELECT id, name, email, bio, profile_picture, created_at FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Bio, &user.ProfilePicture, &user.CreatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// createUser handles POST requests to create a new user.
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert user into the database
	_, err := db.Exec("INSERT INTO users (name, email, bio, profile_picture) VALUES (?, ?, ?, ?)", user.Name, user.Email, user.Bio, user.ProfilePicture)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:admin@tcp(localhost:3306)/profile_app")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")

	r := mux.NewRouter()
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")

	corsObj := handlers.AllowedOrigins([]string{"*"}) // Allow all origins
	corsHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	fmt.Println("Server is listening on port 8000...")
	if err := http.ListenAndServe(":8000", handlers.CORS(corsObj, corsHeaders, corsMethods)(r)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
