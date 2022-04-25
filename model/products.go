package model

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Detail     string  `json:"detail"`
	Price      float64 `json:"price"`
	IsCampaign bool    `json:"is_campaign"`
}
