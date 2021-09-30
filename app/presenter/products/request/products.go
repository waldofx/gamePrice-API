package request

import (
	"gameprice-api/business/products"
)

type Products struct {
	GameID			int	`json:"game_id"`
	SellerID		int	`json:"seller_id"`
	Price 			int	`json:"price"`
}

func ToDomain(request Products) *products.Domain {
	return &products.Domain{
		GameID:     request.GameID,
		SellerID: 	request.SellerID,
		Price: 		request.Price,
	}
}