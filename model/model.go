package model

import (
	"test-api/proto"
)

type Product struct {
	ID         int32   `json:"id" gorm:"primary_key,autoIncrement:1"`
	Name       string  `json:"name"`
	Detail     string  `json:"detail"`
	Price      float32 `json:"price"`
	IsCampaign bool    `json:"is_campaign"`
}

type ProductFilter struct {
	Name       *string  `form:"name" json:"name,omitempty"`
	Detail     *string  `form:"detail" json:"detail"`
	Price      *float32 `form:"price" json:"price"`
	IsCampaign *bool    `form:"is_campaign" json:"is_campaign"`
}

type ProductFilter2 struct {
	Name       string  `form:"name" json:"name,omitempty"`
	Detail     string  `form:"detail" json:"detail"`
	Price      float32 `form:"price" json:"price"`
	IsCampaign bool    `form:"is_campaign" json:"is_campaign"`
}

func (p Product) ToProto() *productService.ProductRes {
	return &productService.ProductRes{
		Id:         p.ID,
		Name:       p.Name,
		Detail:     p.Detail,
		Price:      p.Price,
		IsCampaign: p.IsCampaign,
	}
}

func (p Product) ToProto2() *productService.ProductReq {
	return &productService.ProductReq{
		Id:         p.ID,
		Name:       p.Name,
		Detail:     p.Detail,
		Price:      p.Price,
		IsCampaign: p.IsCampaign,
	}
}
