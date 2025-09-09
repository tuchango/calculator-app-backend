package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"github.com/tuchango/calculator-app-backend/internal/calculation"
	"github.com/tuchango/calculator-app-backend/internal/db"
	"gorm.io/gorm"
)

// TestDB is a test database instance
var TestDB *gorm.DB

// InitializeTestDB initializes a test database
func InitializeTestDB() {
	// For simplicity in basic tests, we'll use an in-memory SQLite database
	// In a real scenario, you might want to use a separate test database
	TestDB = db.InitDB()
}

// CreateTestServer creates a test server with the application routes
func CreateTestServer() *echo.Echo {
	e := echo.New()

	// Initialize database
	InitializeTestDB()

	// Initialize services
	calcRepo := calculation.NewCalculationRepository(TestDB)
	calcService := calculation.NewCalculationService(calcRepo)
	calcHandler := calculation.NewCalculationHandler(calcService)

	// Register routes
	e.GET("/calculations", calcHandler.GetCalculations)
	e.POST("/calculations", calcHandler.PostCalculations)
	e.PATCH("/calculations/:id", calcHandler.PatchCalculations)
	e.DELETE("/calculations/:id", calcHandler.DeleteCalculations)

	return e
}

// SendRequest sends an HTTP request to the test server
func SendRequest(method, url string, body interface{}) (*http.Response, error) {
	// Create test server
	server := CreateTestServer()

	// Create request body if provided
	var reqBody []byte
	if body != nil {
		var err error
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	// Create request
	req := httptest.NewRequest(method, url, bytes.NewBuffer(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// Create response recorder
	rec := httptest.NewRecorder()

	// Serve HTTP
	server.ServeHTTP(rec, req)

	return rec.Result(), nil
}

// ParseResponse parses the response body into the provided interface
func ParseResponse(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

// CleanupTestDB cleans up the test database
func CleanupTestDB() {
	// Clean up database tables
	if TestDB != nil {
		TestDB.Exec("DELETE FROM calculations")
	}
}
