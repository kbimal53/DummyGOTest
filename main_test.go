package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/health", healthCheck).Methods("GET")
	
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Could not parse response JSON")
	}

	if !response.Success {
		t.Errorf("Expected success to be true, got %v", response.Success)
	}
}

func TestGetUsers(t *testing.T) {
	// Initialize test data
	initializeData()

	req, err := http.NewRequest("GET", "/api/v1/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/users", getUsers).Methods("GET")
	
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Could not parse response JSON")
	}

	if !response.Success {
		t.Errorf("Expected success to be true, got %v", response.Success)
	}

	// Check if data is a slice (users array)
	if response.Data == nil {
		t.Error("Expected data to be present")
	}
}

func TestCreateUser(t *testing.T) {
	// Initialize test data
	initializeData()

	user := User{
		Name:  "Test User",
		Email: "test@example.com",
	}

	userJSON, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/users", createUser).Methods("POST")
	
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Could not parse response JSON")
	}

	if !response.Success {
		t.Errorf("Expected success to be true, got %v", response.Success)
	}
}

func TestCreateUserInvalidData(t *testing.T) {
	// Test with missing name
	user := User{
		Email: "test@example.com",
	}

	userJSON, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/users", createUser).Methods("POST")
	
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Could not parse response JSON")
	}

	if response.Success {
		t.Errorf("Expected success to be false for invalid data, got %v", response.Success)
	}
}
