package controllers

import (
	"alterra-miniproject/lib/database"
	"alterra-miniproject/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// get all products
func GetProductsController(c echo.Context) error {
	products, e := database.GetProducts()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all products",
		"data":    products,
	})
}

func GetProductController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product, e := database.GetProduct(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get product",
		"data":    product,
	})
}

// create new user
func CreateProductController(c echo.Context) error {
	weight, _ := strconv.Atoi(c.FormValue("weight"))
	qty, _ := strconv.Atoi(c.FormValue("qty"))
	typeProductID, _ := strconv.Atoi(c.FormValue("type_product_id"))

	product := models.Product{
		Merk:          c.FormValue("merk"),
		Weight:        weight,
		Qty:           qty,
		TypeProductID: typeProductID,
	}
	newProduct, e := database.CreateProduct(product)
	c.Bind(&product)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new product",
		"data":    newProduct,
	})
}

// update user by id
func UpdateProductController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	weight, _ := strconv.Atoi(c.FormValue("weight"))
	qty, _ := strconv.Atoi(c.FormValue("qty"))
	typeProductID, _ := strconv.Atoi(c.FormValue("type_product_id"))

	product := models.Product{
		Model:         gorm.Model{ID: uint(id)},
		Merk:          c.FormValue("merk"),
		Weight:        weight,
		Qty:           qty,
		TypeProductID: typeProductID,
		TypeProduct:   models.TypeProduct{},
	}
	UpdateProduct, e := database.UpdateProduct(product)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update data product",
		"data":    UpdateProduct,
	})
}

// delete user by id
func DeleteProductController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	deleteProduct, e := database.DeleteProduct(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete data product",
		"data":    deleteProduct,
	})
}
