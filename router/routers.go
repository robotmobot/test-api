package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"test-api/controller"
	"test-api/handler"
)

func NewEcho(db gorm.DB) *echo.Echo {
	e := echo.New()

	pc := controller.NewProductController(&db)
	pf := handler.Repo(pc)
	h := handler.NewHandler(pf)

	e.GET("/products", h.GetAllProducts)
	e.GET("/products/:id", h.GetProductByID)
	e.GET("/search", h.FindProduct)
	e.GET("/search-params", h.FindProductQueryParams)
	e.POST("/products", h.CreateProduct)
	e.POST("/batch-products", h.BatchCreateProduct)
	e.PUT("/products/:id", h.UpdateProduct)
	e.DELETE("/products/:id", h.DeleteProduct)

	return e
}
