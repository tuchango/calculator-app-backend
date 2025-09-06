package handlers

import (
	"net/http"

	"calculator-app-backend/internal/calculationService"

	"github.com/labstack/echo/v4"
)

type calcHandler struct {
	service calculationService.CalculationService
}

func NewCalculationHandler(s calculationService.CalculationService) *calcHandler {
	return &calcHandler{service: s}
}

func (h *calcHandler) GetCalculations(c echo.Context) error {
	calculations, err := h.service.GetAllCalculations()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get calculations"})
	}

	return c.JSON(http.StatusOK, calculations)
}

func (h *calcHandler) PostCalculations(c echo.Context) error {
	var req calculationService.CalculationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	calc, err := h.service.CreateCalculation(req.Expression)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create calculation"})
	}

	return c.JSON(http.StatusCreated, calc)
}

func (h *calcHandler) PatchCalculations(c echo.Context) error {
	id := c.Param("id")
	var req calculationService.CalculationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	calc, err := h.service.UpdateCalculation(id, req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not update calculation"})
	}

	return c.JSON(http.StatusOK, calc)
}

func (h *calcHandler) DeleteCalculations(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteCalculation(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not delete calculation"})
	}

	return c.NoContent(http.StatusNoContent)
}
