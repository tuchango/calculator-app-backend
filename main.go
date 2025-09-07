package main

import (
	"calculator-app-backend/internal/calculation"
	"calculator-app-backend/internal/db"
	"calculator-app-backend/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database := db.InitDB()
	calcRepo := calculation.NewCalculationRepository(database)
	calcService := calculation.NewCalculationService(calcRepo)
	calcHandler := handlers.NewCalculationHandler(calcService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/calculations", calcHandler.GetCalculations)
	e.POST("/calculations", calcHandler.PostCalculations)
	e.PATCH("/calculations/:id", calcHandler.PatchCalculations)
	e.DELETE("/calculations/:id", calcHandler.DeleteCalculations)

	e.Start("localhost:8080")
}
