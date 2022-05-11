package controller

import (
	"fmt"
	"strings"
	"test-api/model"

	"gorm.io/gorm"
)

type ProductController struct {
	db *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{
		db: db,
	}
}

//Lists all products
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
func (pf *ProductController) FindProduct(filter *model.ProductFilter) ([]model.Product, error) {
	product := []model.Product{}
	query, args := []string{}, []interface{}{}

	//generate the query from request
	if f := filter.Name; f != nil {
		query, args = append(query, "name = ?"), append(args, *f)
	}

	if f := filter.Price; f != nil {
		query, args = append(query, "price >= ?"), append(args, *f)
	}

	if f := filter.IsCampaign; f != nil {
		query, args = append(query, "is_campaign = ?"), append(args, *f)
	}
	//create the query string
	queryend := strings.Join(query, " AND ")
	err := pf.db.Where(queryend, args...).Find(&product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

//find product from queryparams/url
func (pf *ProductController) FindProductQueryParams(filter *model.ProductFilter2) ([]model.Product, error) {
	product := []model.Product{}
	query, args := []string{}, []interface{}{}

	//generate the query from paramsbinder
	if f := &filter.Name; f != nil {
		query, args = append(query, "name = ?"), append(args, f)
	}

	if f := &filter.Price; f != nil {
		query, args = append(query, "price >= ?"), append(args, f)
	}

	if f := &filter.IsCampaign; f != nil {
		query, args = append(query, "is_campaign = ?"), append(args, f)
	}

	//create the query string
	queryend := strings.Join(query, " AND ")
	err := pf.db.Where(queryend, args...).Find(&product).Error

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

//Creates multiple products from the request body
//takes array of model.Product
//**Maybe put a limit to number of insert
/*func (pf *ProductController) BatchCreateProduct(p []model.Product) error {
	err := pf.db.Create(&p).Error
	if err != nil {
		return err
	}

	return nil
}*/

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
