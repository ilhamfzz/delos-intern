package service

import (
	"delos-intern/helper"
	"delos-intern/model"
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
)

func (s *Service) CreatePond(c echo.Context, pond model.Pond_Binding) (model.Pond, error) {
	pondModels, err := s.GetPonds(c)
	if err != nil {
		return model.Pond{}, errors.New("failed to check ponds")
	}

	if nameFound := helper.IsPondNameExist(pond.Name, pondModels); nameFound {
		return model.Pond{}, errors.New("pond name already exist")
	}

	var farmModels []model.Farm
	if err := s.connection.Find(&farmModels).Error; err != nil {
		return model.Pond{}, errors.New("failed to check farms")
	}
	if pond.FarmID != 0 || pond.FarmID > farmModels[len(farmModels)-1].ID {
		var farm model.Farm
		if idFound := helper.IsFarmIdExist(pond.FarmID, farmModels); idFound {
			farm = model.Farm{
				Name:      "Unnamed Farm" + fmt.Sprint(pond.FarmID),
			}
		}

		if err := s.connection.Save(&farm).Error; err != nil {
			return model.Pond{}, errors.New("failed to create farm")
		}
	} else {
		return model.Pond{}, errors.New("farm id not unique")
	}

	pondModel := model.Pond{
		Name:   pond.Name,
		FarmID: pond.FarmID,
	}
	if err := s.connection.Save(&pondModel).Error; err != nil {
		return model.Pond{}, errors.New("failed to create pond, please input valid farm id")
	}

	return pondModel, nil
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
				Name:      "Unnamed Farm",
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
