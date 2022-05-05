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

func (h *Handler) GetProductByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.ProductController.GetProductByID(id)

	if err != nil {
		return c.JSON(http.StatusNoContent, err)
	}

	return c.JSON(http.StatusOK, product)
}

func (h *Handler) FindProduct(c echo.Context) error {
	query := new(model.ProductFilter)
	err := c.Bind(&query)
	result, _ := h.ProductController.FindProduct(query)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, result)

}
func (h *Handler) FindProductQueryParams(c echo.Context) error {
	filter := new(model.ProductFilter2)
	c.Bind(filter)
	err := echo.QueryParamsBinder(c).String("name", &filter.Name).String("detail", &filter.Detail).Float64("price", &filter.Price).Bool("is_campaign", &filter.IsCampaign).BindError()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, _ := h.ProductController.FindProductQueryParams(filter)
	return c.JSON(http.StatusOK, result)
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
func (h *Handler) BatchCreateProduct(c echo.Context) error {
	var product = []model.Product{}

	err := c.Bind(&product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	errCreate := h.ProductController.BatchCreateProduct(product)
	if errCreate != nil {
		return c.JSON(http.StatusConflict, errCreate)
	}

	return c.JSON(http.StatusOK, product)
}
