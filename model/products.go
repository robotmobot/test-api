package model

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Detail     string  `json:"detail"`
	Price      float64 `json:"price"`
	IsCampaign bool    `json:"is_campaign"`
}

type ProductFilter struct {
	Name       string  `query:"name"`
	Detail     string  `query:"detail"`
	Price      float64 `query:"price"`
	IsCampaign bool    `query:"is_campaign"`
}
