package controller

import (
	"fmt"
	"test-api/model"

	"github.com/jinzhu/gorm"
)

type ProductController struct {
	db *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
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

func (pf *ProductController) GetProduct(id int) (*model.Product, error) {
	product := model.Product{}
	err := pf.db.First(&product, id).Error
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (pf *ProductController) CreateProduct(p *model.Product) error {
	err := pf.db.Create(&p).Error
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}

func (pf *ProductController) UpdateProduct(id int, p *model.Product) (*model.Product, error) {
	product := model.Product{}
	err := pf.db.First(&product, id).Error
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	pf.db.Model(&product).Updates(model.Product{ID: p.ID, Name: p.Name})

	return &product, nil
}

func (pf *ProductController) DeleteProduct(id int) error {
	product := model.Product{}
	err := pf.db.First(&product, id).Error
	fmt.Println(err)
	if err != nil {
		return err
	}

	pf.db.Delete(&product)

	return err
}
