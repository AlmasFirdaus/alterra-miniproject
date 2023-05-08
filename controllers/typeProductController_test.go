package controllers

import (
	"alterra-miniproject/config"
	"alterra-miniproject/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetTypeProductsController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/typ", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// testing
	if assert.NoError(t, GetTypeProductsController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetTypeProductController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/typ/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// testing
	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateTypeProductController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	typProduct := models.TypeProduct{
		Name: "electronics",
	}
	typProductJSON, _ := json.Marshal(typProduct)
	req := httptest.NewRequest(http.MethodPost, "/typ", bytes.NewBuffer(typProductJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// testing
	if assert.NoError(t, CreateTypeProductController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// validate response
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "success create new type product", response["message"])

		// validate data inserted to DB
		var result models.Product
		config.DB.First(&result, response["data"].(map[string]interface{})["id"])
		assert.Equal(t, typProduct.Name, result.TypeProduct.Name)
	}
}

func TestUpdateTypeProductController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	typProduct := models.TypeProduct{
		Name: "snackTest",
	}
	config.DB.Create(&typProduct)
	typProduct.Name = "electronics"
	typProductJSON, _ := json.Marshal(typProduct)
	req := httptest.NewRequest(http.MethodPut, "/typ/1", bytes.NewBuffer(typProductJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// testing
	if assert.NoError(t, UpdateTypeProductController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// validate response
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "success update data type product", response["message"])

		// validate data updated in DB
		var result models.TypeProduct
		config.DB.First(&result, typProduct.ID)
		assert.Equal(t, typProduct.Name, result.Name)
	}
}

func TestDeleteTypeProductController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/typ/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	//Assertion
	if assert.NoError(t, DeleteTypeProductController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
