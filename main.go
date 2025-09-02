package main

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Calculation struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationRequest struct {
	Expression string `json:"expression"`
}

var calculations = []Calculation{}

func calculateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}

	res, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(res), nil
}

func getCalculations(c echo.Context) error {
	return c.JSON(http.StatusOK, calculations)
}

func postCalculations(c echo.Context) error {
	var req CalculationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	res, err := calculateExpression(req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprint("Invalid expression: ", err)})
	}

	calc := Calculation{ID: uuid.NewString(), Expression: req.Expression, Result: res}
	calculations = append(calculations, calc)
	return c.JSON(http.StatusCreated, calc)
}

func patchCalculations(c echo.Context) error {
	id := c.Param("id")
	var req CalculationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	res, err := calculateExpression(req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprint("Invalid expression: ", err)})
	}

	for i, calc := range calculations {
		if calc.ID == id {
			calculations[i].Expression = req.Expression
			calculations[i].Result = res
			return c.JSON(http.StatusOK, calculations[i])
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Calculation not found"})
}

func deleteCalculations(c echo.Context) error {
	id := c.Param("id")

	for i, calc := range calculations {
		if calc.ID == id {
			calculations = slices.Delete(calculations, i, i+1)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Calculation not found"})
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/calculations", getCalculations)
	e.POST("/calculations", postCalculations)
	e.PATCH("/calculations/:id", patchCalculations)
	e.DELETE("/calculations/:id", deleteCalculations)

	e.Start("localhost:8080")
}
