package request

import (
	"gameprice-api/business/products"
)

type Products struct {
	ID				int 	`json:"id"`
	GameID			int		`json:"game_id"`
	SellerID		int		`json:"seller_id"`
	Price 			string	`json:"price"`
}

func ToDomain(request Products) *products.Domain {
	return &products.Domain{
		ID: request.ID,
		GameID:     request.GameID,
		SellerID: 	request.SellerID,
		Price: 		request.Price,
	}
}