package gogapis

import (
	"gameprice-api/business/gogapis"
)

type GOG struct {
	ID                         int    `json:"id"`
	Title                      string `json:"title"`
	Links struct {
		PurchaseLink string `json:"purchase_link"`
		ProductCard  string `json:"product_card"`
		Support      string `json:"support"`
		Forum        string `json:"forum"`
	} `json:"links"`
	IsSecret      bool   `json:"is_secret"`
	IsInstallable bool   `json:"is_installable"`
	GameType      string `json:"game_type"`
}

func (rec *GOG) ToDomain() gogapis.Domain {
	return gogapis.Domain{
		URL: rec.Links.ProductCard,
	}
}