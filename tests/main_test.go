package tests

import (
	"os"
	"testing"
)

// TestMain is the entry point for all tests
func TestMain(m *testing.M) {
	// Set up test environment
	setup()

	// Run tests
	code := m.Run()

	// Clean up
	teardown()

	// Exit with the same code as the tests
	os.Exit(code)
}

// setup prepares the test environment
func setup() {
	// Initialize any required resources for tests
	// For now, we'll just print a message
	println("Setting up test environment...")
}

// teardown cleans up the test environment
func teardown() {
	// Clean up any resources used during tests
	// For now, we'll just print a message
	println("Tearing down test environment...")
}
