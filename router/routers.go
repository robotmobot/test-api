package router

import (
	"test-api/controller"
	"test-api/handler"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func New(db *gorm.DB) *echo.Echo {
	e := echo.New()

	pf := controller.NewProductController(db)
	h := handler.NewHandler(*pf)

	e.GET("/products", h.GetAllProducts)
	e.GET("/products/:id", h.GetProduct)
	e.POST("/products", h.CreateProduct)
	e.PUT("/products/:id", h.UpdateProduct)
	e.DELETE("/products/:id", h.DeleteProduct)

	return e
}
