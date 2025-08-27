package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// User represents a user in our system
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Created  string `json:"created"`
}

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// In-memory storage for demo purposes
var users []User
var nextID = 1

func main() {
	// Initialize with some dummy data
	initializeData()

	// Create router
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api/v1").Subrouter()
	
	// User endpoints
	api.HandleFunc("/users", getUsers).Methods("GET")
	api.HandleFunc("/users/{id}", getUser).Methods("GET")
	api.HandleFunc("/users", createUser).Methods("POST")
	api.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	
	// Health check endpoint
	api.HandleFunc("/health", healthCheck).Methods("GET")

	// Root endpoint
	r.HandleFunc("/", rootHandler).Methods("GET")

	// Add CORS middleware
	r.Use(corsMiddleware)

	fmt.Println("🚀 Server starting on http://localhost:8080")
	fmt.Println("📋 Available endpoints:")
	fmt.Println("  GET    /                    - Root endpoint")
	fmt.Println("  GET    /api/v1/health       - Health check")
	fmt.Println("  GET    /api/v1/users        - Get all users")
	fmt.Println("  GET    /api/v1/users/{id}   - Get user by ID")
	fmt.Println("  POST   /api/v1/users        - Create new user")
	fmt.Println("  PUT    /api/v1/users/{id}   - Update user")
	fmt.Println("  DELETE /api/v1/users/{id}   - Delete user")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func initializeData() {
	users = []User{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Created: time.Now().Format(time.RFC3339)},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Created: time.Now().Format(time.RFC3339)},
		{ID: 3, Name: "Bob Johnson", Email: "bob@example.com", Created: time.Now().Format(time.RFC3339)},
	}
	nextID = 4
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
	response := Response{
		Success: true,
		Message: "API is healthy",
		Data: map[string]string{
			"status":    "ok",
			"timestamp": time.Now().Format(time.RFC3339),
		},
	}
	sendJSONResponse(w, http.StatusOK, response)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
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

	for _, user := range users {
		if user.ID == id {
			response := Response{
				Success: true,
				Message: "User found",
				Data:    user,
			}
			sendJSONResponse(w, http.StatusOK, response)
			return
		}
	}

	response := Response{
		Success: false,
		Message: "User not found",
	}
	sendJSONResponse(w, http.StatusNotFound, response)
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

	// Assign ID and timestamp
	newUser.ID = nextID
	nextID++
	newUser.Created = time.Now().Format(time.RFC3339)

	// Add to users slice
	users = append(users, newUser)

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

	// Find and update user
	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id
			updatedUser.Created = user.Created // Keep original creation time
			users[i] = updatedUser

			response := Response{
				Success: true,
				Message: "User updated successfully",
				Data:    updatedUser,
			}
			sendJSONResponse(w, http.StatusOK, response)
			return
		}
	}

	response := Response{
		Success: false,
		Message: "User not found",
	}
	sendJSONResponse(w, http.StatusNotFound, response)
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

	// Find and delete user
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			response := Response{
				Success: true,
				Message: "User deleted successfully",
			}
			sendJSONResponse(w, http.StatusOK, response)
			return
		}
	}

	response := Response{
		Success: false,
		Message: "User not found",
	}
	sendJSONResponse(w, http.StatusNotFound, response)
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
