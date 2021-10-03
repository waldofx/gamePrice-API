package response

import (
	"gameprice-api/business/products"
	"time"
)

type Products struct {
	ID        	int       	`json:"id"`
	GameID		int			`json:"game_id"`
	Game		string		`json:"game"`
	SellerID		int		`json:"seller_id"`
	Seller		string		`json:"seller"`
	Price		int			`json:"price"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}

func FromDomain(domain products.Domain) Products {
	return Products{
		ID:        	domain.ID,
		Game:      	domain.Game,
		GameID:     domain.GameID,
		Seller: 	domain.Seller,
		SellerID: 	domain.SellerID,
		Price:		domain.Price,
		CreatedAt: 	domain.CreatedAt,
		UpdatedAt: 	domain.UpdatedAt,
	}
}

func NewResponseArray(modelProducts []products.Domain) []Products{
	var response []Products

	for _, val := range modelProducts{
		response = append(response, FromDomain(val))
	}
	return response
}