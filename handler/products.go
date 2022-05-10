package handler

import (
	"net/http"
	"strconv"
	"sync"
	"test-api/controller"
	"test-api/model"
	"time"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	ProductController controller.ProductController
	wg                sync.WaitGroup
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
	products := []model.Product{}
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Millisecond))
	done := make(chan bool, 1)
	done <- true
	//defer cancel()
	err := c.Bind(&products)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	h.wg.Add(len(products))
	for i := 0; i < len(products); i++ {
		product := products[i]
		select {
		case <-done:
			go func(product model.Product) {
				err := h.ProductController.CreateProduct(&product)
				if err != nil {
					c.Logger().Error(err)
				}
				h.wg.Done()
				done <- true
			}(product)
		case <-time.After(1000 * time.Millisecond):
			return c.JSON(http.StatusRequestTimeout, "timeout")
		}
	}
	h.wg.Wait()

	return c.JSON(http.StatusOK, products)
}
