package controller

import (
	"delos-intern/dto"
	"delos-intern/model"
	"delos-intern/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var farmService service.Svc

func NewFarmController(service service.Svc) {
	farmService = service
}

func CreateFarm(c echo.Context) error {
	farm := model.Farm_Binding{}
	if err := c.Bind(&farm); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to bind farm", err))
	}

	result, err := farmService.CreateFarm(c, farm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to create farm", err))
	}
	return c.JSON(http.StatusCreated, dto.BuildResponse("Successfully created farm", result))
}

func UpdateFarm(c echo.Context) error {
	farm := model.Farm_Binding{}
	if err := c.Bind(&farm); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to bind farm", err))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to parse id", err))
	}

	result, err := farmService.UpdateFarm(c, id, farm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to update farm", err))
	}

	if result.Name == "register but already soft deleted" && result.Ponds == nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to create farm", errors.New("your id is not unique")))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Successfully updated farm", result))
}

func DeleteFarm(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to parse id", err))
	}

	err = farmService.DeleteFarm(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to delete farm", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Successfully deleted farm", nil))
}

func GetFarms(c echo.Context) error {
	result, err := farmService.GetFarms(c)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.BuildErrorResponse("Failed to get farms", err))
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, dto.BuildErrorResponse("Pond not found", nil))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Successfully get farms", result))
}

func GetFarmById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to parse id", err))
	}

	result, err := farmService.GetFarmById(c, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.BuildErrorResponse("Failed to get farm", err))
	}
	var respon model.Farm
	if result.ID == respon.ID && result.Name == respon.Name && result.Ponds == nil {
		return c.JSON(http.StatusNotFound, dto.BuildErrorResponse("Farm not found", nil))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Successfully get farm", result))
}
