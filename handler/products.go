package handler

import (
	"net/http"
	"strconv"
	"test-api/controller"
	"test-api/model"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	ProductController controller.ProductController
}

func NewHandler(pf controller.ProductController) *Handler {
	return &Handler{
		ProductController: pf,
	}
}

func (h *Handler) GetAllProducts(c echo.Context) error {
	products, err := h.ProductController.GetAllProducts()

	if err != nil {
		return c.JSON(http.StatusNoContent, err)
	}

	return c.JSON(http.StatusOK, products)
}

func (h *Handler) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.ProductController.GetProduct(id)

	if err != nil {
		return c.JSON(http.StatusNoContent, err)
	}

	return c.JSON(http.StatusOK, product)
}

func (h *Handler) CreateProduct(c echo.Context) error {
	product := model.Product{}

	err := c.Bind(&product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	errCreate := h.ProductController.CreateProduct(&product)
	if errCreate != nil {
		return c.JSON(http.StatusConflict, errCreate)
	}

	return c.JSON(http.StatusOK, product)
}

func (h *Handler) UpdateProduct(c echo.Context) error {
	product := model.Product{}

	err := c.Bind(&product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	response, errUpdate := h.ProductController.UpdateProduct(id, &product)
	if errUpdate != nil {
		return c.JSON(http.StatusConflict, errUpdate)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.ProductController.DeleteProduct((id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, err)
}
