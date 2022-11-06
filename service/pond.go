package service

import (
	"delos-intern/helper"
	"delos-intern/model"
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
)

func (s *Service) CreatePond(c echo.Context, pond model.Pond_Binding) (model.Pond, error) {
	var flag bool = false
	pondModels, err := s.GetPonds(c)
	if err != nil {
		flag = true
	}

	idFound := helper.IsPondNameExist(pond.Name, pondModels)
	if !idFound {
		flag = true
	}

	if flag {
		pondModel := model.Pond{
			Name:   pond.Name,
			FarmID: pond.FarmID,
		}

		farmModels, err := s.GetFarms(c)
		if err != nil {
			return model.Pond{}, errors.New("farm not found")
		}
		idFound := helper.IsFarmIdExist(pond.FarmID, farmModels)
		if !idFound {
			farmModels := model.Farm{
				Name: "Unnamed Farm " + fmt.Sprint(pond.FarmID),
			}
			err = s.connection.Create(&farmModels).Error
			if err != nil {
				return model.Pond{}, errors.New("failed to create farm")
			}
			pondModel.FarmID = farmModels.ID
		}

		err = s.connection.Create(&pondModel).Error
		if err != nil {
			return model.Pond{}, errors.New("failed to create pond")
		}
		return pondModel, nil
	} else {
		return model.Pond{}, errors.New("pond already exist")
	}
}

func (s *Service) UpdatePond(c echo.Context, id int, pond model.Pond_Binding) (model.Pond, error) {
	pondModels, err := s.GetPonds(c)
	if err != nil {
		return model.Pond{}, errors.New("failed to check ponds")
	}

	idFound := helper.IsPondIdExist(uint(id), pondModels)
	if !idFound { // if pond not found, then create new pond
		var farmModels []model.Farm
		farmModels, err = s.GetFarms(c)
		if err != nil {
			return model.Pond{}, errors.New("failed to check farms")
		}
		if idFound := helper.IsFarmIdExist(pond.FarmID, farmModels); !idFound {
			// if farm not found, then create new farm
			farm := model.Farm{
				Name: "Unnamed Farm",
			}
			err = s.connection.Create(&farm).Error
			if err != nil {
				return model.Pond{}, errors.New("failed to create farm")
			}
		}

		pondModel := model.Pond{
			Name:   pond.Name,
			FarmID: pond.FarmID,
		}
		err = s.connection.Save(&pondModel).Error
		if err != nil {
			return model.Pond{}, errors.New("failed to create pond")
		}
		return pondModel, nil
	} else { // else update pond
		pondModel, err := s.GetPondById(c, id)
		if err != nil {
			return model.Pond{}, errors.New("failed to get pond")
		}

		pondModel.Name = pond.Name
		if pond.FarmID != 0 {
			pondModel.FarmID = pond.FarmID
		}

		err = s.connection.Model(&pondModel).Updates(pondModel).Error
		if err != nil {
			return model.Pond{}, errors.New("failed to update pond")
		}

		return pondModel, nil
	}
}

func (s *Service) DeletePond(c echo.Context, id int) error {
	var pondModel model.Pond
	err := s.connection.Where("id = ?", id).First(&pondModel).Error
	if err != nil {
		return errors.New("pond not found")
	}

	err = s.connection.Delete(&pondModel).Error
	if err != nil {
		return errors.New("failed to delete pond")
	}
	return nil
}

func (s *Service) GetPonds(c echo.Context) ([]model.Pond, error) {
	var ponds []model.Pond
	err := s.connection.Find(&ponds).Error
	if err != nil {
		return nil, errors.New("pond not found")
	}

	return ponds, nil
}

func (s *Service) GetPondById(c echo.Context, id int) (model.Pond, error) {
	var pond model.Pond
	err := s.connection.Where("id = ?", id).First(&pond).Error
	if err != nil {
		return model.Pond{}, errors.New("pond not found")
	}

	return pond, nil
}
