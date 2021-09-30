package request

import (
	"gameprice-api/business/wishes"
)

type Wishes struct {
	ID				int `json:"id"`
	UserID			int	`json:"user_id"`
	ProductID		int	`json:"product_id"`
}

func ToDomain(request Wishes) *wishes.Domain {
	return &wishes.Domain{
		ID: request.ID,
		UserID:     request.UserID,
	}
}