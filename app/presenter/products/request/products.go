package request

import (
	"gameprice-api/business/products"
)

type ProductInsert struct {
	GameID			int	`json:"game_id"`
	SellerID		int	`json:"seller_id"`
	Price 			int	`json:"price"`
}

type ProductUpdate struct {
	ID 			int		`json:"id"`
	Game		string	`json:"game"`
	Seller		string	`json:"seller"`
}

func ToDomain(request ProductInsert) *products.Domain {
	return &products.Domain{
		GameID:     request.GameID,
		SellerID: 	request.SellerID,
		Price: 		request.Price,
	}
}

func ToDomainUpdate(request ProductUpdate) *products.Domain {
	return &products.Domain{
		ID:        	request.ID,
		Game:      	request.Game,
		Seller: 	request.Seller,
	}
}