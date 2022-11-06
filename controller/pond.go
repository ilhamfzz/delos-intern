package controller

import (
	"delos-intern/dto"
	"delos-intern/model"
	"delos-intern/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var pondService service.Svc

func NewPondController(service service.Svc) {
	pondService = service
}

func CreatePond(c echo.Context) error {
	pond := model.Pond_Binding{}
	if err := c.Bind(&pond); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to bind pond", err))
	}

	result, err := pondService.CreatePond(c, pond)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to create pond", err))
	}
	return c.JSON(http.StatusCreated, dto.BuildResponse("Successfully create pond", result))
}

func UpdatePond(c echo.Context) error {
	pond := model.Pond_Binding{}
	if err := c.Bind(&pond); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to bind pond", err))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to parse id", err))
	}

	result, err := pondService.UpdatePond(c, id, pond)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to update pond", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Successfully updated pond", result))
}

func DeletePond(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to parse id", err))
	}

	err = pondService.DeletePond(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to delete pond", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Successfully deleted pond", nil))
}

func GetPonds(c echo.Context) error {
	result, err := pondService.GetPonds(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get ponds", err))
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, dto.BuildErrorResponse("Pond not found", nil))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Successfully get ponds", result))
}

func GetPondById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to parse id", err))
	}

	result, err := pondService.GetPondById(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get pond", err))
	}
	var empty model.Pond
	if result == empty {
		return c.JSON(http.StatusNotFound, dto.BuildErrorResponse("Pond not found", nil))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Successfully get pond", result))
}
