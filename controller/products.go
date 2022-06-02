package controller

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/context"
	"test-api/model"
	"time"
)

type ProductController struct {
	Db DbRepo
	Rc redis.Client
}

var ctx = context.Background()

func NewProductController(db DbRepo, Rc redis.Client) *ProductController {
	return &ProductController{
		Db: db,
		Rc: Rc,
	}
}

//GetAllProducts
//List all products from the table
func (pf *ProductController) GetAllProducts() ([]model.Product, error) {
	var products []model.Product

	err := pf.Db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

// GetProductByID
//List product provided by the ID
func (pf *ProductController) GetProductByID(id int32) (*model.Product, error) {
	var product model.Product
	test, err := pf.Rc.Get(ctx, string(id)).Bytes()
	if err != nil {
		log.Errorf("Cache error on Set :%v", err)
	}
	if err == nil {
		product, _ = product.UnmarshalBinary(test)
		return &product, nil
	}

	err = pf.Db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// FindProduct
//find product by creating query from request body
func (pf *ProductController) FindProduct(filter model.ProductFilter) ([]model.Product, error) {
	var product []model.Product

	err := pf.Db.Where("name = ? AND price >= ? AND is_campaign = ?", *filter.Name, *filter.Price, *filter.IsCampaign).Find(&product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

// FindProductQueryParams
//find product from query params/url
func (pf *ProductController) FindProductQueryParams(filter *model.ProductFilter2) ([]model.Product, error) {
	var product []model.Product
	err := pf.Db.Where("name = ? AND price >= ? AND is_campaign = ?", filter.Name, filter.Price, filter.IsCampaign).Find(&product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

// CreateProduct
//Creates one product from the request body
func (pf *ProductController) CreateProduct(p *model.Product) error {
	err := pf.Db.Create(&p).Error
	if err != nil {
		return err
	}
	pRedis, err := p.MarshalBinary()
	if err != nil {
		log.Error(err)
	}
	err = pf.Rc.Set(ctx, string(p.ID), pRedis, 90*time.Second).Err()
	if err != nil {
		log.Errorf("Caching error :%v", err)
	}
	return nil
}

//UpdateProduct Takes the id of product and fields to update
//Updates the field of product of that id
func (pf *ProductController) UpdateProduct(id int32, p *model.Product) (*model.Product, error) {
	product := model.Product{}
	err := pf.Db.First(&product, id).Error

	if err != nil {
		return nil, err
	}

	pf.Db.Model(&product).Updates(model.Product{ID: p.ID, Name: p.Name})

	return &product, nil
}

// DeleteProduct
//Deletes the product from the request /products/:id
func (pf *ProductController) DeleteProduct(id int32) error {
	product := model.Product{}
	err := pf.Db.First(&product, id).Error
	fmt.Println(product)
	if err != nil {
		return err
	}

	pf.Db.Delete(&product)

	return err
}
