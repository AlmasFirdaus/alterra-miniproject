package controllers

import (
	"alterra-miniproject/config"
	"alterra-miniproject/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetProductsController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// testing
	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetProductController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// testing
	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateProductController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	product := models.Product{
		Merk:          "fantech",
		Weight:        5,
		Qty:           5,
		TypeProductID: 1,
	}
	productJSON, _ := json.Marshal(product)
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(productJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// testing
	if assert.NoError(t, CreateProductController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// validate response
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "success create new product", response["message"])

		// validate data inserted to DB
		var result models.Product
		config.DB.First(&result, response["data"].(map[string]interface{})["id"])
		assert.Equal(t, product.Merk, result.Merk)
		assert.Equal(t, product.Weight, result.Weight)
		assert.Equal(t, product.Qty, result.Qty)
		assert.Equal(t, product.TypeProductID, result.TypeProductID)
	}
}

func TestUpdateProductController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	product := models.Product{
		Merk:          "fantech",
		Weight:        10,
		Qty:           10,
		TypeProductID: 1,
	}
	config.DB.Create(&product)
	product.Merk = "fantech"
	productJSON, _ := json.Marshal(product)
	req := httptest.NewRequest(http.MethodPut, "/products/1"+strconv.Itoa(int(product.ID)), bytes.NewBuffer(productJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(product.ID)))

	// testing
	if assert.NoError(t, UpdateProductController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// validate response
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "success update data product", response["message"])

		// validate data updated in DB
		var result models.Product
		config.DB.First(&result, product.ID)
		assert.Equal(t, product.Merk, result.Merk)
		assert.Equal(t, product.Weight, result.Weight)
		assert.Equal(t, product.Qty, result.Qty)
		assert.Equal(t, product.TypeProductID, result.TypeProductID)
	}
}

func TestDeleteProductController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	//Assertion
	if assert.NoError(t, DeleteProductController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
