package main

import (
	"test-api/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//routes
	e.GET("/products/:id", handler.GetProduct)
	e.GET("/products", handler.GetAllProducts)
	e.POST("/products", handler.CreateProduct)
	e.PUT("/products/:id", handler.UpdateProduct)
	e.DELETE("/products/:id", handler.DeleteProduct)

	e.Logger.Fatal(e.Start(":1323"))

}
