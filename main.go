package main

import (
	"github.com/tuchango/calculator-app-backend/internal/calculation"
	"github.com/tuchango/calculator-app-backend/internal/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database := db.InitDB()
	calcRepo := calculation.NewCalculationRepository(database)
	calcService := calculation.NewCalculationService(calcRepo)
	calcHandler := calculation.NewCalculationHandler(calcService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/calculations", calcHandler.GetCalculations)
	e.POST("/calculations", calcHandler.PostCalculations)
	e.PATCH("/calculations/:id", calcHandler.PatchCalculations)
	e.DELETE("/calculations/:id", calcHandler.DeleteCalculations)

	e.Start("localhost:8080")
}
