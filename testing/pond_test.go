package testing

import (
	"delos-intern/config"
	"delos-intern/controller"
	"delos-intern/model"
	"delos-intern/service"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitPondTestAPI() *echo.Echo {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
	db := config.InitDatabaseTest()
	Svc := service.NewService(db)
	controller.NewPondController(Svc)
	e := echo.New()
	return e
}

func TestCreatePond(t *testing.T) {
	testCase := []struct {
		Name         string
		Path         string
		ExpectedCode int
		SizeData     int
	}{
		{
			Name:         "sucessfully create pond",
			Path:         "/pond",
			ExpectedCode: http.StatusCreated,
			SizeData:     1,
		},
	}

	e := InitPondTestAPI()
	f := make(url.Values)
	f.Set("name", "test")
	f.Set("farm_id", "1")
	req := httptest.NewRequest(http.MethodPost, "/pond", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			if assert.NoError(t, controller.CreatePond(c)) {
				assert.Equal(t, tc.ExpectedCode, rec.Code)
				var ponds []model.Pond
				model.DB.Find(&ponds)
				assert.Equal(t, tc.SizeData, len(ponds))
			}
		})
	}
}

func InserDataPond() {
	farm := model.Farm{
		Name: "test",
	}
	err := model.DB.Create(&farm).Error
	if err != nil {
		panic(err)
	}
	pond := model.Pond{
		Name:   "test",
		FarmID: 1,
	}
	err = model.DB.Create(&pond).Error
	if err != nil {
		panic(err)
	}
	farm = model.Farm{
		Ponds: []model.Pond{pond},
	}
	err = model.DB.Save(&farm).Error
	if err != nil {
		panic(err)
	}
}

func TestUpdatePond(t *testing.T) {
	testCase := []struct {
		Name         string
		Path         string
		ExpectedCode int
		SizeData     int
	}{
		{
			Name:         "sucessfully update pond",
			Path:         "/pond/1",
			ExpectedCode: http.StatusOK,
			SizeData:     1,
		},
	}

	e := InitPondTestAPI()
	InserDataPond()
	f := make(url.Values)
	f.Set("name", "test updated")
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/pond/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			if assert.NoError(t, controller.UpdatePond(c)) {
				assert.Equal(t, tc.ExpectedCode, rec.Code)
				var ponds []model.Pond
				model.DB.Find(&ponds)
				assert.Equal(t, tc.SizeData, len(ponds))
			}
		})
	}
}

func TestDeletePond(t *testing.T) {
	testCase := []struct {
		Name         string
		Path         string
		ExpectedCode int
		SizeData     int
	}{
		{
			Name:         "sucessfully delete pond",
			Path:         "/pond/1",
			ExpectedCode: http.StatusOK,
			SizeData:     0,
		},
	}

	e := InitPondTestAPI()
	InserDataPond()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/pond/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			if assert.NoError(t, controller.DeletePond(c)) {
				assert.Equal(t, tc.ExpectedCode, rec.Code)
				var ponds []model.Pond
				model.DB.Find(&ponds)
				assert.Equal(t, tc.SizeData, len(ponds))
			}
		})
	}
}

func TestGetPond(t *testing.T) {
	testCase := []struct {
		Name         string
		Path         string
		ExpectedCode int
		SizeData     int
	}{
		{
			Name:         "sucessfully get pond",
			Path:         "/pond",
			ExpectedCode: http.StatusOK,
			SizeData:     1,
		},
	}

	e := InitPondTestAPI()
	InserDataPond()
	req := httptest.NewRequest(http.MethodGet, "/pond", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			if assert.NoError(t, controller.GetPonds(c)) {
				assert.Equal(t, tc.ExpectedCode, rec.Code)
				var ponds []model.Pond
				model.DB.Find(&ponds)
				assert.Equal(t, tc.SizeData, len(ponds))
			}
		})
	}
}

func TestGetPondByID(t *testing.T) {
	testCase := []struct {
		Name         string
		Path         string
		ExpectedCode int
		SizeData     int
	}{
		{
			Name:         "sucessfully get pond by id",
			Path:         "/pond/1",
			ExpectedCode: http.StatusOK,
			SizeData:     1,
		},
	}

	e := InitPondTestAPI()
	InserDataPond()
	req := httptest.NewRequest(http.MethodGet, "/pond/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/pond/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			if assert.NoError(t, controller.GetPondById(c)) {
				assert.Equal(t, tc.ExpectedCode, rec.Code)
				var ponds []model.Pond
				model.DB.Find(&ponds)
				assert.Equal(t, tc.SizeData, len(ponds))
			}
		})
	}
}
