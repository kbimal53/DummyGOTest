package internal

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB initializes the database connection and creates tables if they don't exist
func InitDB() error {
	var err error
	
	// Get database URL from environment
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	// Connect to database
	DB, err = sql.Open("postgres", databaseURL)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL database")

	// Create users table if it doesn't exist
	if err = createUsersTable(); err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}

	return nil
}

// createUsersTable creates the users table if it doesn't exist
func createUsersTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := DB.Exec(query)
	if err != nil {
		return err
	}

	log.Println("✅ Users table ready")
	return nil
}

// InsertInitialData inserts sample data if the table is empty
func InsertInitialData() error {
	// Check if table has data
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return err
	}

	// If table is empty, insert sample data
	if count == 0 {
		sampleUsers := []struct {
			Name  string
			Email string
		}{
			{"John Doe", "john@example.com"},
			{"Jane Smith", "jane@example.com"},
			{"Bob Johnson", "bob@example.com"},
		}

		   for _, user := range sampleUsers {
			   _, err := DB.Exec(
				   "INSERT INTO users (name, email) VALUES ($1, $2)",
				   user.Name, user.Email,
			   )
			   if err != nil {
				   return err
			   }
		   }
		log.Println("✅ Sample data inserted")
	}

	return nil
}

// CloseDB closes the database connection
func CloseDB() {
       if DB != nil {
	       DB.Close()
       }
}
