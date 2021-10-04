package response

import (
	"gameprice-api/business/sellers"
	"time"
)

type Sellers struct {
	ID        int       `json:"id"`
	Name	  string	`json:"name"`
	URL  	  string	`json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain sellers.Domain) Sellers {
	return Sellers{
		ID:        domain.ID,
		Name:      domain.Name,
		URL:  	   domain.URL,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func NewResponseArray(modelSellers []sellers.Domain) []Sellers{
	var response []Sellers

	for _, val := range modelSellers{
		response = append(response, FromDomain(val))
	}
	return response
}