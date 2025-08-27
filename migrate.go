package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Check if DATABASE_URL is set
	if os.Getenv("DATABASE_URL") == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Initialize database
	if err := initDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer closeDB()

	log.Println("🎉 Database migration completed successfully!")
	log.Println("📊 Users table is ready")

	// Insert initial data
	if err := insertInitialData(); err != nil {
		log.Printf("Warning: Failed to insert initial data: %v", err)
	} else {
		log.Println("✅ Sample data inserted (if table was empty)")
	}
}
