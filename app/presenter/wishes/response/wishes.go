package response

import (
	"gameprice-api/business/wishes"
	"time"
)

type Wishes struct {
	ID        	int       	`json:"id"`
	UserID		int			`json:"user_id"`
	User		string		`json:"user"`
	ProductID	int			`json:"product_id"`
	Product		string		`json:"product"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}

func FromDomain(domain wishes.Domain) Wishes {
	return Wishes{
		ID:        	domain.ID,
		User:      	domain.User,
		UserID:     domain.UserID,
		Product: 	domain.Product,
		ProductID: 	domain.ProductID,
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