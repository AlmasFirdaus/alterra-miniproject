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

func TestGetUsersController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// testing
	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUserController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// testing
	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateUserController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	user := models.User{
		Name:     "almas1",
		Email:    "almas1@gmail.com",
		Password: "almas1",
	}
	userJSON, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// testing
	if assert.NoError(t, CreateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// // validate response
		// var response map[string]interface{}
		// json.Unmarshal(rec.Body.Bytes(), &response)
		// assert.Equal(t, "success create new user", response["message"])

		// // validate data inserted to DB
		// var result models.User
		// config.DB.First(&result, response["users"].(map[string]interface{})["id"])
		// assert.Equal(t, user.Name, result.Name)
		// assert.Equal(t, user.Email, result.Email)
		// assert.Equal(t, user.Password, result.Password)
	}
}

func TestUpdateUserController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	user := models.User{
		Name:     "almas1",
		Email:    "almas1@gmail.com",
		Password: "almas123",
	}
	config.DB.Create(&user)
	user.Name = "almas1"
	userJSON, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPut, "/users/1"+strconv.Itoa(int(user.ID)), bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(user.ID)))

	// testing
	if assert.NoError(t, UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// // validate response
		// var response map[string]interface{}
		// json.Unmarshal(rec.Body.Bytes(), &response)
		// assert.Equal(t, "success update data user", response["message"])

		// // validate data updated in DB
		// var result models.User
		// config.DB.First(&result, user.ID)
		// assert.Equal(t, user.Name, result.Name)
		// assert.Equal(t, user.Email, result.Email)
		// assert.Equal(t, user.Password, result.Password)
	}
}

func TestDeleteUserController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	//Assertion
	if assert.NoError(t, DeleteUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
