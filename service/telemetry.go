package service

import (
	"delos-intern/model"
	"errors"

	"github.com/labstack/echo/v4"
)

func (s *Service) GetTelemetry(c echo.Context) ([]model.CountRespon, error) {
	var count []model.CountRespon
	err := s.connection.Model(&model.Telemetry{}).Select("endpoint", "method", "count(*) as count", "count(DISTINCT ip) as unique_user").Group("endpoint").Group("method").Find(&count).Error
	if err != nil {
		return nil, errors.New("failed to get telemetry")
	}
	return count, nil
}
