package controller

import (
	"delos-intern/dto"
	"delos-intern/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

var telemetryService service.Svc

func NewTelemetryController(service service.Svc) {
	telemetryService = service
}

func GetTelemetry(c echo.Context) error {
	result, err := telemetryService.GetTelemetry(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get telemetry", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Successfully get telemetry", result))
}
