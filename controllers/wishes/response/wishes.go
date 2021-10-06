package response

import (
	"gameprice-api/business/wishes"
	"time"
)

type Wishes struct {
	ID        	int       	`json:"id"`
	UserID		int			`json:"user_id"`
	Username	string		`json:"username"`
	GameID		int			`json:"game_id"`
	GameName	string		`json:"game_name"`
	SellerID	int 		`json:"seller_id"`
	GameSeller	string		`json:"game_seller"`
	Price		string		`json:"price"`
	Discount	bool		`json:"discount"`
	URL			string		`json:"url"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}

func FromDomain(domain wishes.Domain) Wishes {
	return Wishes{
		ID:        	domain.ID,
		Username:   domain.Username,
		UserID:     domain.UserID,
		GameID:  	domain.GameID,
		GameName: 	domain.GameName,
		SellerID: 	domain.SellerID,
		GameSeller: domain.GameSeller,
		Price: 		domain.Price,
		Discount: 	domain.Discount,
		URL: 		domain.URL,
		CreatedAt: 	domain.CreatedAt,
		UpdatedAt: 	domain.UpdatedAt,
	}
}

func NewResponseArray(modelWishes []wishes.Domain) []Wishes{
	var response []Wishes

	for _, val := range modelWishes{
		response = append(response, FromDomain(val))
	}
	return response
}