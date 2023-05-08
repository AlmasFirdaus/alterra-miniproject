package controllers

import (
	"alterra-miniproject/lib/database"
	"alterra-miniproject/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// get all users
func GetUsersController(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	users, e := database.GetUser(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"users":   users,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{
		Name:     c.FormValue("name"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	newUser, e := database.CreateUser(user)
	c.Bind(&user)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"users":   newUser,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{
		Model:    gorm.Model{ID: uint(id)},
		Name:     c.FormValue("name"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	UpdateUser, e := database.UpdateUser(user)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update data user",
		"users":   UpdateUser,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	deleteUser, e := database.DeleteUser(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete data user",
		"users":   deleteUser,
	})
}

func GetUserDetailControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	users, err := database.GetDetailUsers(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func LoginUsersController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	users, e := database.LoginUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}
