package tests

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitTestDB initializes a test database using SQLite in-memory database
func InitTestDB() *gorm.DB {
	// For testing purposes, we'll use an in-memory SQLite database
	// This is faster and doesn't require a separate test database setup
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}

	// Auto-migrate the schema
	// In a real application, you would import your models here
	// For now, we'll just return the database connection
	return db
}

// GetTestDBConfig gets test database configuration from environment variables
// or uses default values for testing
func GetTestDBConfig() (host, user, password, dbname, port, sslmode string) {
	host = getEnvOrDefault("TEST_DB_HOST", "localhost")
	user = getEnvOrDefault("TEST_DB_USER", "postgres")
	password = getEnvOrDefault("TEST_DB_PASSWORD", "password")
	dbname = getEnvOrDefault("TEST_DB_NAME", "test_db")
	port = getEnvOrDefault("TEST_DB_PORT", "5432")
	sslmode = getEnvOrDefault("TEST_DB_SSLMODE", "disable")

	return host, user, password, dbname, port, sslmode
}

// getEnvOrDefault gets an environment variable or returns a default value
func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
