package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// DatabaseConfig represents the configuration for PostgreSQL connection
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	SSLMode  string `json:"sslmode"`
	DBName   string `json:"dbname"`
}

// ConnectToPostgres establishes a connection to PostgreSQL database
func ConnectToPostgres(config DatabaseConfig) (*sql.DB, error) {
	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
		config.SSLMode,
	)

	// Open connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error creating connection pool: %v", err)
	}

	// Verify the connection
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	return db, nil
}

// testConnectionHandler handles POST requests to test database connection
func testConnectionHandler(w http.ResponseWriter, r *http.Request) {
	var config DatabaseConfig

	// Decode JSON request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&config); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Attempt to connect
	db, err := ConnectToPostgres(config)
	if err != nil {
		http.Error(w, fmt.Sprintf("Connection failed: %v", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Test query
	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		http.Error(w, fmt.Sprintf("Query failed: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"version": version,
	})
}

func main() {
	// Create router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/test-connection", testConnectionHandler).Methods("POST")

	// Start server
	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
