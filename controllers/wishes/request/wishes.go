package request

import (
	"gameprice-api/business/wishes"
)

type Wishes struct {
	ID				int `json:"id"`
	UserID			int	`json:"user_id"`
	GameID			int	`json:"game_id"`
	SellerID		int `json:"seller_id"`
}

func ToDomain(request Wishes) *wishes.Domain {
	return &wishes.Domain{
		ID:		 	request.ID,
		UserID:     request.UserID,
		GameID:  	request.GameID,
		SellerID: 	request.SellerID,
	}
}