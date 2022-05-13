package controller_test

import (
	"test-api/controller"
	"test-api/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func TestGetAllProducts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mocks := mocks.NewMockDbRepo(mockCtrl)
	tempCall := gorm.DB{}
	mocks.EXPECT().Find(gomock.Any()).Return(&tempCall)
	controller := controller.NewProductController(mocks)

	products, err := controller.GetAllProducts()
	if err != nil {
		t.FailNow()
	}
	if len(products) < 1 {
		t.FailNow()
	}

}
