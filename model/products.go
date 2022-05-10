package model

type Product struct {
	ID         uint    `json:"id" gorm:"primary_key,autoIncrement:1"`
	Name       string  `json:"name"`
	Detail     string  `json:"detail"`
	Price      float64 `json:"price"`
	IsCampaign bool    `json:"is_campaign"`
}

type ProductFilter struct {
	Name       *string  `form:"name" json:"name,omitempty"`
	Detail     *string  `form:"detail" json:"detail"`
	Price      *float64 `form:"price" json:"price"`
	IsCampaign *bool    `form:"is_campaign" json:"is_campaign"`
}

type ProductFilter2 struct {
	Name       string  `form:"name" json:"name"`
	Detail     string  `form:"detail" json:"detail"`
	Price      float64 `form:"price" json:"price"`
	IsCampaign bool    `form:"is_campaign" json:"is_campaign"`
}
