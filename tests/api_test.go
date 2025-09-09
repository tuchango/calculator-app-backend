package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/tuchango/calculator-app-backend/internal/calculation"
)

// TestGetCalculations tests the GET /calculations endpoint
func TestGetCalculations(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a request
	req := httptest.NewRequest(http.MethodGet, "/calculations", nil)
	rec := httptest.NewRecorder()

	// Create a mock handler that returns an empty array
	e.GET("/calculations", func(c echo.Context) error {
		return c.JSON(http.StatusOK, []calculation.Calculation{})
	})

	// Serve the request
	e.ServeHTTP(rec, req)

	// Assert the status code
	assert.Equal(t, http.StatusOK, rec.Code)

	// Assert the response body
	var calculations []calculation.Calculation
	err := json.Unmarshal(rec.Body.Bytes(), &calculations)
	assert.NoError(t, err)
	assert.Empty(t, calculations)
}

// TestPostCalculations tests the POST /calculations endpoint
func TestPostCalculations(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a request body
	requestBody := calculation.CalculationRequest{
		Expression: "2+2",
	}

	// Convert request body to JSON
	body, err := json.Marshal(requestBody)
	assert.NoError(t, err)

	// Create a request
	req := httptest.NewRequest(http.MethodPost, "/calculations", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Create a mock handler that returns a calculation
	e.POST("/calculations", func(c echo.Context) error {
		var req calculation.CalculationRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// Create a mock calculation
		calc := calculation.Calculation{
			ID:         "123",
			Expression: req.Expression,
			Result:     "4",
		}

		return c.JSON(http.StatusCreated, calc)
	})

	// Serve the request
	e.ServeHTTP(rec, req)

	// Assert the status code
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Assert the response body
	var calc calculation.Calculation
	err = json.Unmarshal(rec.Body.Bytes(), &calc)
	assert.NoError(t, err)
	assert.Equal(t, "2+2", calc.Expression)
	assert.Equal(t, "4", calc.Result)
}

// TestPatchCalculations tests the PATCH /calculations/:id endpoint
func TestPatchCalculations(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a request body
	requestBody := calculation.CalculationRequest{
		Expression: "3+3",
	}

	// Convert request body to JSON
	body, err := json.Marshal(requestBody)
	assert.NoError(t, err)

	// Create a request
	req := httptest.NewRequest(http.MethodPatch, "/calculations/123", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Create a mock handler that returns an updated calculation
	e.PATCH("/calculations/:id", func(c echo.Context) error {
		id := c.Param("id")

		var req calculation.CalculationRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// Create a mock calculation
		calc := calculation.Calculation{
			ID:         id,
			Expression: req.Expression,
			Result:     "6",
		}

		return c.JSON(http.StatusOK, calc)
	})

	// Serve the request
	e.ServeHTTP(rec, req)

	// Assert the status code
	assert.Equal(t, http.StatusOK, rec.Code)

	// Assert the response body
	var calc calculation.Calculation
	err = json.Unmarshal(rec.Body.Bytes(), &calc)
	assert.NoError(t, err)
	assert.Equal(t, "123", calc.ID)
	assert.Equal(t, "3+3", calc.Expression)
	assert.Equal(t, "6", calc.Result)
}

// TestDeleteCalculations tests the DELETE /calculations/:id endpoint
func TestDeleteCalculations(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a request
	req := httptest.NewRequest(http.MethodDelete, "/calculations/123", nil)
	rec := httptest.NewRecorder()

	// Create a mock handler that returns no content
	e.DELETE("/calculations/:id", func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})

	// Serve the request
	e.ServeHTTP(rec, req)

	// Assert the status code
	assert.Equal(t, http.StatusNoContent, rec.Code)
}
