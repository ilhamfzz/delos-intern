package route

import (
	"delos-intern/controller"
	"delos-intern/middleware"
	"delos-intern/service"

	"github.com/labstack/echo/v4"
)

func New(Svc service.Svc) *echo.Echo {
	controller.NewFarmController(Svc)
	controller.NewPondController(Svc)
	controller.NewTelemetryController(Svc)

	api := echo.New()
	api.Use(middleware.Telemetry())

	api.POST("/farm", controller.CreateFarm)
	api.PUT("/farm/:id", controller.UpdateFarm)
	api.DELETE("/farm/:id", controller.DeleteFarm)
	api.GET("/farm", controller.GetFarms)
	api.GET("/farm/:id", controller.GetFarmById)

	api.POST("/pond", controller.CreatePond)
	api.PUT("/pond/:id", controller.UpdatePond)
	api.DELETE("/pond/:id", controller.DeletePond)
	api.GET("/pond", controller.GetPonds)
	api.GET("/pond/:id", controller.GetPondById)

	api.GET("/telemetry", controller.GetTelemetry)

	return api
}
