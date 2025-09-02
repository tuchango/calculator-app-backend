package main

import (
	"fmt"
	"net/http"

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

func calculateExpression(expression string) string {
	expr, _ := govaluate.NewEvaluableExpression(expression)
	res, _ := expr.Evaluate(nil)
	return fmt.Sprint(res)
}

func getCalculations(c echo.Context) error {
	c.JSON(http.StatusOK, calculations)
	return nil
}

func postCalculations(c echo.Context) error {
	var req CalculationRequest
	_ = c.Bind(&req)
	res := calculateExpression(req.Expression)
	calc := Calculation{ID: uuid.NewString(), Expression: req.Expression, Result: res}
	calculations = append(calculations, calc)
	c.JSON(http.StatusCreated, calc)
	return nil
}

func patchCalculations(c echo.Context) error {
	id := c.Param("id")
	var req CalculationRequest
	_ = c.Bind(&req)
	res := calculateExpression(req.Expression)

	for i, calc := range calculations {
		if calc.ID == id {
			calculations[i].Expression = req.Expression
			calculations[i].Result = res
			return c.JSON(http.StatusOK, calculations[i])
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "calc not found"})
}

func deleteCalculations(c echo.Context) error {
	id := c.Param("id")

	for i, calc := range calculations {
		if calc.ID == id {
			calculations = remove(calculations, i)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "calc not found"})
}

func remove(s []Calculation, i int) []Calculation {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
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
