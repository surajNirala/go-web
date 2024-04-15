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
	ID    int    `json:"id"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	// Add other fields as needed
}

func main() {
	fmt.Println("deded")
	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/get-users", getUsers)
	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	// Replace the connection string with your MySQL connection string
	db, err := sql.Open("mysql", "srj:12345678@tcp(localhost:3306)/go_crud1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the database connection is successful
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, phone, email FROM users limit 20")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Phone, &user.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert users slice to JSON
	jsonBytes, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set content type header and write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Replace the connection string with your MySQL connection string
	// db, err := sql.Open("mysql", "srj:12345678@tcp(localhost:3306)/go_crud1")
	db, err := sql.Open("mysql", "srj:admin@tcp(65.2.37.239:3306)/srj_crud")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error connecting to database: %v", err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error pinging database: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Database connection successful")
}
