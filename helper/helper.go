package helper

import (
	"delos-intern/model"
	"fmt"
)

func IsFarmIdExist(id uint, ids []model.Farm) bool {
	for _, v := range ids {
		if v.ID == id {
			return true
		}
	}
	return false
}

func IsFarmNameExist(name string, names []model.Farm) bool {
	for _, v := range names {
		if v.Name == name {
			return true
		}
	}
	return false
}

func IsPondIdExist(id uint, ids []model.Pond) bool {
	for _, v := range ids {
		if v.ID == id {
			return true
		}
	}
	return false
}

func IsPondNameExist(name string, names []model.Pond) bool {
	for _, v := range names {
		if v.Name == name {
			return true
		}
	}
	return false
}

func GetPondsByFarmId(farmId uint, ponds []model.Pond) []model.Pond {
	var pondsByFarmId []model.Pond
	for _, v := range ponds {
		if v.FarmID == farmId {
			pondsByFarmId = append(pondsByFarmId, v)
		}
	}
	return pondsByFarmId
}

func CreateUnnamedFarm(id uint) model.Farm {
	return model.Farm{
		Name: "Unnamed Farm " + fmt.Sprint(id),
	}
}
