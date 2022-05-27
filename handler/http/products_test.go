package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"test-api/mocks"
	"test-api/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllProducts_Success(t *testing.T) {
	var products []model.Product

	ctrl := gomock.NewController(t)
	mockCtrl := mocks.NewMockRepo(ctrl)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/products")
	mockCtrl.EXPECT().GetAllProducts().Return(products, nil)

	handler := NewHandler(mockCtrl)

	assert.NoError(t, handler.GetAllProducts(c))
	assert.Equal(t, http.StatusOK, rec.Code)

}

func TestGetProductByID_Success(t *testing.T) {
	var product model.Product
	var err error

	ctrl := gomock.NewController(t)
	mockCtrl := mocks.NewMockRepo(ctrl)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/products/:id", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/products/")
	c.SetParamNames("id")
	c.SetParamValues("0")

	mockCtrl.EXPECT().GetProductByID(0).Return(&product, err)

	handler := NewHandler(mockCtrl)

	assert.NoError(t, handler.GetProductByID(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetProductByID_Error(t *testing.T) {
	var product model.Product
	err := errors.New("bad request")

	ctrl := gomock.NewController(t)
	mockCtrl := mocks.NewMockRepo(ctrl)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/products/:id", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/products/")
	c.SetParamNames("id")
	c.SetParamValues("0")

	mockCtrl.EXPECT().GetProductByID(0).Return(&product, err)
	handler := NewHandler(mockCtrl)

	assert.NoError(t, handler.GetProductByID(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

//TestCreateProduct_JSONBody
//Send JSON body with request through echo context,check error.
func TestCreateProduct_JSONBody(t *testing.T) {
	product := model.Product{}
	var err error
	ctrl := gomock.NewController(t)
	mockCtrl := mocks.NewMockRepo(ctrl)
	e := echo.New()

	body, err := json.Marshal(product)
	if err != nil {
		t.Error(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetRequest(req)
	mockCtrl.EXPECT().CreateProduct(&product).Return(nil)
	handler := NewHandler(mockCtrl)

	assert.NoError(t, handler.CreateProduct(c))
	assert.Equal(t, http.StatusOK, rec.Code)

}

//TestCreateProduct_ControllerError
//Controller returns error, check it against expected behaviour
func TestCreateProduct_ControllerError(t *testing.T) {
	product := model.Product{}
	err := errors.New("bad request")

	ctrl := gomock.NewController(t)
	mockCtrl := mocks.NewMockRepo(ctrl)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/products", nil)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	mockCtrl.EXPECT().CreateProduct(&product).Return(err)
	handler := NewHandler(mockCtrl)

	assert.NoError(t, handler.CreateProduct(c))
	assert.Equal(t, http.StatusConflict, rec.Code)
}
