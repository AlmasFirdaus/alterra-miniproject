package routes

import (
	"alterra-miniproject/constants"
	"alterra-miniproject/controllers"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Route / to handler function
	// login
	e.POST("/login", controllers.LoginUsersController)

	e.POST("/users", controllers.CreateUserController)

	//jwt group
	r := e.Group("/jwt")
	r.Use(echojwt.JWT([]byte(constants.SECRET_JWT)))
	// users
	r.GET("/users", controllers.GetUsersController)
	r.GET("/users/:id", controllers.GetUserDetailControllers)
	r.POST("/users", controllers.CreateUserController)
	r.DELETE("/users/:id", controllers.DeleteUserController)
	r.PUT("/users/:id", controllers.UpdateUserController)

	// products
	r.GET("/products", controllers.GetProductsController)
	r.GET("/products/:id", controllers.GetProductController)
	r.POST("/products", controllers.CreateProductController)
	r.DELETE("/products/:id", controllers.DeleteProductController)
	r.PUT("/products/:id", controllers.UpdateProductController)

	// type products
	r.GET("/typ", controllers.GetTypeProductsController)
	r.GET("/typ/:id", controllers.GetTypeProductController)
	r.POST("/typ", controllers.CreateTypeProductController)
	r.DELETE("/typ/:id", controllers.DeleteTypeProductController)
	r.PUT("/typ/:id", controllers.UpdateTypeProductController)

	return e
}
