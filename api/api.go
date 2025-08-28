// DUMMY COMMENT FOR GIT PUSH TEST
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// User represents a user in our system
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Created string `json:"created"`
}

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

var router *mux.Router

// init function runs before main and sets up the router for Vercel
func init() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Initialize database connection
	if err := initDB(); err != nil {
		log.Printf("Database initialization error: %v", err)
		// Don't fatal here for Vercel - let individual requests handle DB errors
	} else {
		// Insert initial data if needed //dummy comment
		if err := insertInitialData(); err != nil {
			log.Printf("Warning: Failed to insert initial data: %v", err)
		}
	}

	// Create router
	router = mux.NewRouter()

	// API routes (support both /api/v1 and /api for health)
	apiv1 := router.PathPrefix("/api/v1").Subrouter()
	api := router.PathPrefix("/api").Subrouter()

	// User endpoints (/api/v1)
	apiv1.HandleFunc("/users", getUsers).Methods("GET")
	apiv1.HandleFunc("/users/{id}", getUser).Methods("GET")
	apiv1.HandleFunc("/users", createUser).Methods("POST")
	apiv1.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	apiv1.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	apiv1.HandleFunc("/health", healthCheck).Methods("GET")


	// User endpoints (/api)

	api.HandleFunc("/users", getUsers).Methods("GET")
	api.HandleFunc("/users/{id}", getUser).Methods("GET")
	api.HandleFunc("/users", createUser).Methods("POST")
	api.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	api.HandleFunc("/health", healthCheck).Methods("GET")

	// Serve frontend static files
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./public/"))))

	// Root endpoint (API info)
	router.HandleFunc("/api", rootHandler).Methods("GET")

	// Add CORS middleware
	router.Use(corsMiddleware)
}

// Handler is the main entry point for Vercel serverless functions
func Handler(w http.ResponseWriter, r *http.Request) {
	// Ensure database is connected for each request
	if db == nil {
		if err := initDB(); err != nil {
			http.Error(w, "Database connection failed", http.StatusInternalServerError)
			return
		}
	}

	// Serve the request using the router
	router.ServeHTTP(w, r)
}

// main function for local development
func main() {
	// This is only used for local development
	// Vercel will use the Handler function instead
	
	fmt.Println("🚀 Server starting on http://localhost:8080")
	fmt.Println("🗄️  Connected to PostgreSQL database")
	fmt.Println("📋 Available endpoints:")
	fmt.Println("  GET    /                    - Root endpoint")
	fmt.Println("  GET    /api/v1/health       - Health check")
	fmt.Println("  GET    /api/v1/users        - Get all users")
	fmt.Println("  GET    /api/v1/users/{id}   - Get user by ID")
	fmt.Println("  POST   /api/v1/users        - Create new user")
	fmt.Println("  PUT    /api/v1/users/{id}   - Update user")
	fmt.Println("  DELETE /api/v1/users/{id}   - Delete user")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Success: true,
		Message: "Welcome to the Dummy Go API!",
		Data: map[string]string{
			"version": "1.0.0",
			"docs":    "Visit /api/v1/health for health check",
		},
	}
	sendJSONResponse(w, http.StatusOK, response)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	// Check database connection
	dbStatus := "ok"
	if db == nil {
		dbStatus = "disconnected"
	} else if err := db.Ping(); err != nil {
		dbStatus = "error"
	}

	response := Response{
		Success: true,
		Message: "API is healthy",
		Data: map[string]string{
			"status":    "ok",
			"database":  dbStatus,
			"timestamp": time.Now().Format(time.RFC3339),
		},
	}
	sendJSONResponse(w, http.StatusOK, response)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, email, created_at FROM users ORDER BY id")
	if err != nil {
		response := Response{
			Success: false,
			Message: "Failed to fetch users",
		}
		sendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		var createdAt time.Time
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &createdAt)
		if err != nil {
			response := Response{
				Success: false,
				Message: "Failed to scan user data",
			}
			sendJSONResponse(w, http.StatusInternalServerError, response)
			return
		}
		user.Created = createdAt.Format(time.RFC3339)
		users = append(users, user)
	}

	response := Response{
		Success: true,
		Message: "Users retrieved successfully",
		Data:    users,
	}
	sendJSONResponse(w, http.StatusOK, response)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := Response{
			Success: false,
			Message: "Invalid user ID",
		}
		sendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	var user User
	var createdAt time.Time
	err = db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &createdAt)
	
	if err != nil {
		response := Response{
			Success: false,
			Message: "User not found",
		}
		sendJSONResponse(w, http.StatusNotFound, response)
		return
	}

	user.Created = createdAt.Format(time.RFC3339)
	response := Response{
		Success: true,
		Message: "User found",
		Data:    user,
	}
	sendJSONResponse(w, http.StatusOK, response)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Invalid JSON data",
		}
		sendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	// Validate required fields
	if newUser.Name == "" || newUser.Email == "" {
		response := Response{
			Success: false,
			Message: "Name and email are required",
		}
		sendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	// Insert user into database
	var createdAt time.Time
	err = db.QueryRow(
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at",
		newUser.Name, newUser.Email,
	).Scan(&newUser.ID, &createdAt)

	if err != nil {
		response := Response{
			Success: false,
			Message: "Failed to create user",
		}
		sendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	newUser.Created = createdAt.Format(time.RFC3339)

	response := Response{
		Success: true,
		Message: "User created successfully",
		Data:    newUser,
	}
	sendJSONResponse(w, http.StatusCreated, response)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := Response{
			Success: false,
			Message: "Invalid user ID",
		}
		sendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	var updatedUser User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Invalid JSON data",
		}
		sendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	// Validate required fields
	if updatedUser.Name == "" || updatedUser.Email == "" {
		response := Response{
			Success: false,
			Message: "Name and email are required",
		}
		sendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	// Update user in database
	var createdAt time.Time
	err = db.QueryRow(
		"UPDATE users SET name = $1, email = $2 WHERE id = $3 RETURNING id, created_at",
		updatedUser.Name, updatedUser.Email, id,
	).Scan(&updatedUser.ID, &createdAt)

	if err != nil {
		response := Response{
			Success: false,
			Message: "User not found",
		}
		sendJSONResponse(w, http.StatusNotFound, response)
		return
	}

	updatedUser.Created = createdAt.Format(time.RFC3339)

	response := Response{
		Success: true,
		Message: "User updated successfully",
		Data:    updatedUser,
	}
	sendJSONResponse(w, http.StatusOK, response)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := Response{
			Success: false,
			Message: "Invalid user ID",
		}
		sendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	// Delete user from database
	result, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Failed to delete user",
		}
		sendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		response := Response{
			Success: false,
			Message: "User not found",
		}
		sendJSONResponse(w, http.StatusNotFound, response)
		return
	}

	response := Response{
		Success: true,
		Message: "User deleted successfully",
	}
	sendJSONResponse(w, http.StatusOK, response)
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
