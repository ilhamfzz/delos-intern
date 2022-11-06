package testing

import (
	"delos-intern/config"
	"delos-intern/controller"
	"delos-intern/model"
	"delos-intern/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitFarmTestAPI() *echo.Echo {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
	db := config.InitDatabaseTest()
	Svc := service.NewService(db)
	controller.NewFarmController(Svc)
	e := echo.New()
	return e
}

func TestCreateFarm(t *testing.T) {
	testCase := []struct {
		Name         string
		Path         string
		ExpectedCode int
		SizeData     int
	}{
		{
			Name:         "sucessfully create farm",
			Path:         "/farm",
			ExpectedCode: http.StatusCreated,
			SizeData:     1,
		},
	}

	e := InitFarmTestAPI()
	userJSON := `{"name":"test"}`
	req := httptest.NewRequest(http.MethodPost, "/farm", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			if assert.NoError(t, controller.CreateFarm(c)) {
				assert.Equal(t, tc.ExpectedCode, rec.Code)
				var farms []model.Farm
				model.DB.Find(&farms)
				assert.Equal(t, tc.SizeData, len(farms))
			}
		})
	}
}

func InserDataFarm() {
	farm := model.Farm{
		Name: "test",
	}
	err := model.DB.Save(&farm).Error
	if err != nil {
		panic(err)
	}
}

func TestUpdateFarm(t *testing.T) {
	testCase := []struct {
		Name         string
		Path         string
		ExpectedCode int
		SizeData     int
	}{
		{
			Name:         "sucessfully update farm",
			Path:         "/farm/1",
			ExpectedCode: http.StatusOK,
			SizeData:     1,
		},
	}

	e := InitFarmTestAPI()
	InserDataFarm()
	userJSON := `{"name":"test updated"}`
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/farm/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			if assert.NoError(t, controller.UpdateFarm(c)) {
				assert.Equal(t, tc.ExpectedCode, rec.Code)
				var farms []model.Farm
				model.DB.Find(&farms)
				assert.Equal(t, tc.SizeData, len(farms))
			}
		})
	}
}

func TestDeleteFarm(t *testing.T) {
	testCase := []struct {
		Name         string
		Path         string
		ExpectedCode int
		SizeData     int
	}{
		{
			Name:         "sucessfully delete farm",
			Path:         "/farm/1",
			ExpectedCode: http.StatusOK,
			SizeData:     0,
		},
	}

	e := InitFarmTestAPI()
	InserDataFarm()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/farm/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			if assert.NoError(t, controller.DeleteFarm(c)) {
				assert.Equal(t, tc.ExpectedCode, rec.Code)
				var farms []model.Farm
				model.DB.Find(&farms)
				assert.Equal(t, tc.SizeData, len(farms))
			}
		})
	}
}

func TestGetFarm(t *testing.T) {
	testCase := []struct {
		Name         string
		Path         string
		ExpectedCode int
		SizeData     int
	}{
		{
			Name:         "sucessfully get farm",
			Path:         "/farm",
			ExpectedCode: http.StatusOK,
			SizeData:     1,
		},
	}

	e := InitFarmTestAPI()
	InserDataFarm()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			if assert.NoError(t, controller.GetFarms(c)) {
				assert.Equal(t, tc.ExpectedCode, rec.Code)
				var farms []model.Farm
				model.DB.Find(&farms)
				assert.Equal(t, tc.SizeData, len(farms))
			}
		})
	}
}

func TestGetFarmByID(t *testing.T) {
	testCase := []struct {
		Name         string
		Path         string
		ExpectedCode int
		SizeData     int
	}{
		{
			Name:         "sucessfully get farm by id",
			Path:         "/farm/1",
			ExpectedCode: http.StatusOK,
			SizeData:     1,
		},
	}

	e := InitFarmTestAPI()
	InserDataFarm()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/farm/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			if assert.NoError(t, controller.GetFarmById(c)) {
				assert.Equal(t, tc.ExpectedCode, rec.Code)
				var farms []model.Farm
				model.DB.Find(&farms)
				assert.Equal(t, tc.SizeData, len(farms))
			}
		})
	}
}
