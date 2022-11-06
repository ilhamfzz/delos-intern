package service

import (
	"delos-intern/helper"
	"delos-intern/model"
	"errors"

	"github.com/labstack/echo/v4"
)

func (s *Service) CreateFarm(c echo.Context, farm model.Farm_Binding) (model.Farm, error) {
	farmModels, err := s.GetFarms(c)
	if err != nil {
		return model.Farm{}, errors.New("failed to get farms")
	}
	if nameFound := helper.IsFarmNameExist(farm.Name, farmModels); nameFound {
		return model.Farm{}, errors.New("farm name already exist")
	}

	var farmModel model.Farm
	if len(farmModels) == 0 {
		farmModel.ID = 1
	} else {
		farmModel.ID = farmModels[len(farmModels)-1].ID + 1
	}

	farmModel = model.Farm{
		Name: farm.Name,
	}

	err = s.connection.Create(&farmModel).Error
	if err != nil {
		return model.Farm{}, errors.New("failed to create farm")
	}
	return farmModel, nil
}

func (s *Service) UpdateFarm(c echo.Context, id int, farm model.Farm_Binding) (model.Farm, error) {
	var farmModel model.Farm
	farmModels, err := s.GetFarms(c)
	if err != nil {
		return model.Farm{}, errors.New("failed to get farms")
	}
	if idFound := helper.IsFarmIdExist(uint(id), farmModels); !idFound {
		if len(farmModels) == 0 || uint(id) > farmModels[len(farmModels)-1].ID {
			// if no registered farm or make sure id is not already soft deleted
			farmModel = model.Farm{
				Name: farm.Name,
			}
		} else { // if id is already soft deleted
			return model.Farm{
				Name: "register but already soft deleted",
			}, nil
		}
		err = s.connection.Create(&farmModel).Error
		if err != nil {
			return model.Farm{}, errors.New("failed to create farm")
		}
		return farmModel, nil
	}

	err = s.connection.Where("id = ?", id).First(&farmModel).Error
	if err != nil {
		return model.Farm{}, errors.New("farm not found")
	}

	farmModel.Name = farm.Name
	err = s.connection.Updates(&farmModel).Error
	if err != nil {
		return model.Farm{}, errors.New("failed to update farm")
	}
	return farmModel, nil
}

func (s *Service) DeleteFarm(c echo.Context, id int) error {
	var farmModel model.Farm
	err := s.connection.Where("id = ?", id).First(&farmModel).Error
	if err != nil {
		return errors.New("farm not found")
	}

	err = s.connection.Delete(&farmModel).Error
	if err != nil {
		return errors.New("failed to delete farm")
	}
	return nil
}

func (s *Service) GetFarms(c echo.Context) ([]model.Farm, error) {
	var farmModels []model.Farm
	err := s.connection.Find(&farmModels).Error
	if err != nil {
		return nil, errors.New("farm not found")
	}

	var pondModels []model.Pond
	err = s.connection.Find(&pondModels).Error
	if err != nil {
		return nil, errors.New("failed to get ponds")
	}

	for i := 0; i < len(farmModels); i++ {
		for j := 0; j < len(pondModels); j++ {
			if uint(farmModels[i].ID) == pondModels[j].FarmID {
				farmModels[i].Ponds = append(farmModels[i].Ponds, pondModels[j])
			}
		}
	}
	return farmModels, nil
}

func (s *Service) GetFarmById(c echo.Context, id int) (model.Farm, error) {
	var farmModel model.Farm
	err := s.connection.Where("id = ?", id).First(&farmModel).Error
	if err != nil {
		return model.Farm{}, errors.New("farm not found")
	}

	var pondModels []model.Pond
	err = s.connection.Find(&pondModels).Error
	if err != nil {
		return model.Farm{}, errors.New("failed to get ponds")
	}

	for i := 0; i < len(pondModels); i++ {
		if farmModel.ID == pondModels[i].FarmID {
			farmModel.Ponds = append(farmModel.Ponds, pondModels[i])
		}
	}
	return farmModel, nil
}
