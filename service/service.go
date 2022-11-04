package service

import (
	"delos-intern/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Svc interface {
	CreateFarm(c echo.Context, farm model.Farm_Binding) (model.Farm, error)
	UpdateFarm(c echo.Context, id int, farm model.Farm_Binding) (model.Farm, error)
	DeleteFarm(c echo.Context, id int) error
	GetFarms(c echo.Context) ([]model.Farm, error)
	GetFarmById(c echo.Context, id int) (model.Farm, error)

	CreatePond(c echo.Context, pond model.Pond_Binding) (model.Pond, error)
	UpdatePond(c echo.Context, id int, pond model.Pond_Binding) (model.Pond, error)
	DeletePond(c echo.Context, id int) error
	GetPonds(c echo.Context) ([]model.Pond, error)
	GetPondById(c echo.Context, id int) (model.Pond, error)

	GetTelemetry(c echo.Context) ([]model.CountRespon, error)
}

type Service struct {
	connection *gorm.DB
}

func NewService(db *gorm.DB) Svc {
	return &Service{
		connection: db,
	}
}
