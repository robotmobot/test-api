package handler

import (
	"test-api/model"
)

type Repo interface {
	GetAllProducts() ([]model.Product, error)
	GetProductByID(id int32) (*model.Product, error)
	FindProduct(filter model.ProductFilter) ([]model.Product, error)
	FindProductQueryParams(filter *model.ProductFilter2) ([]model.Product, error)
	CreateProduct(p *model.Product) error
	UpdateProduct(id int32, p *model.Product) (*model.Product, error)
	DeleteProduct(id int32) error
}
