package middleware

import (
	"delos-intern/model"
	"errors"
	"time"

	"github.com/labstack/echo/v4"
)

func Telemetry() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			telemetry := model.Telemetry{
				Ip:       c.RealIP(),
				Method:   c.Request().Method,
				Endpoint: c.Request().URL.Path,
				Status:   c.Response().Status,
				Latency:  stop.Sub(start).Milliseconds(),
			}

			if err := model.DB.Create(&telemetry).Error; err != nil {
				return errors.New("failed to create telemetry")
			}
			return nil
		}
	}
}
