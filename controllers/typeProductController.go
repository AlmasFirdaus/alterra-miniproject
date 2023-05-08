package controllers

import (
	"alterra-miniproject/lib/database"
	"alterra-miniproject/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// get all type products
func GetTypeProductsController(c echo.Context) error {
	typeProducts, e := database.GetTypeProducts()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all type products",
		"data":    typeProducts,
	})
}

func GetTypeProductController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	typeProduct, e := database.GetTypeProduct(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get type product",
		"data":    typeProduct,
	})
}

// create new type product
func CreateTypeProductController(c echo.Context) error {
	typeProduct := models.TypeProduct{
		Name: c.FormValue("name"),
	}
	newTypeProduct, e := database.CreateTypeProduct(typeProduct)
	c.Bind(&typeProduct)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new type product",
		"data":    newTypeProduct,
	})
}

// update type product by id
func UpdateTypeProductController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	typeProduct := models.TypeProduct{
		Model: gorm.Model{ID: uint(id)},
		Name:  c.FormValue("name"),
	}
	updateTypeProduct, e := database.UpdateTypeProduct(typeProduct)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update data type product",
		"data":    updateTypeProduct,
	})
}

// delete type product by id
func DeleteTypeProductController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	deleteTypeProduct, e := database.DeleteTypeProduct(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete data type product",
		"data":    deleteTypeProduct,
	})
}
