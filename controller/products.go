package controller

import (
	"fmt"
	"test-api/model"
)

type ProductController struct {
	db DbRepo
}

func NewProductController(db DbRepo) *ProductController {
	return &ProductController{
		db: db,
	}
}

func (pf *ProductController) GetAllProducts() ([]model.Product, error) {
	products := []model.Product{}
	response := pf.db.Find(&products)

	if response.Error != nil {
		return nil, response.Error
	}

	return products, nil
}

//List product provided by the ID
func (pf *ProductController) GetProductByID(id int) (*model.Product, error) {
	product := model.Product{}
	err := pf.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

//find product by creating query from request body
func (pf *ProductController) FindProduct(filter model.ProductFilter) ([]model.Product, error) {
	product := []model.Product{}

	err := pf.db.Where("name = ? AND price >= ? AND is_campaign = ?", *filter.Name, *filter.Price, *filter.IsCampaign).Find(&product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

//find product from queryparams/url
func (pf *ProductController) FindProductQueryParams(filter *model.ProductFilter2) ([]model.Product, error) {
	product := []model.Product{}
	err := pf.db.Where("name = ? AND price >= ? AND is_campaign = ?", filter.Name, filter.Price, filter.IsCampaign).Find(&product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

//Creates one product from the request body
func (pf *ProductController) CreateProduct(p *model.Product) error {
	err := pf.db.Create(&p).Error
	if err != nil {
		return err
	}

	return nil
}

//Takes the id of product and fields to update
//Updates the field of product of that id
func (pf *ProductController) UpdateProduct(id int, p *model.Product) (*model.Product, error) {
	product := model.Product{}
	err := pf.db.First(&product, id).Error

	if err != nil {
		return nil, err
	}

	pf.db.Model(&product).Updates(model.Product{ID: p.ID, Name: p.Name})

	return &product, nil
}

//Deletes the product from the request /products/:id
func (pf *ProductController) DeleteProduct(id int) error {
	product := model.Product{}
	err := pf.db.First(&product, id).Error
	fmt.Println(product)
	if err != nil {
		fmt.Println("check")
		return err
	}

	pf.db.Delete(&product)

	return err
}
