package controller_test

import (
	"test-api/controller"
	"test-api/mocks"
	"test-api/model"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetAllProducts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	products := []model.Product{
		{
			ID:   1,
			Name: "test",
		},
	}
	mockProductRepo := mocks.NewMockProductsRepo(mockCtrl)

	mockProductRepo.EXPECT().GetAllProducts().Return(products, nil)
	controller.GetAllProducts(mockProductRepo)

}
