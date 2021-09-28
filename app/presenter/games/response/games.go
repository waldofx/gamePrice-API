package response

import (
	"gameprice-api/business/games"
	"time"
)

type Games struct {
	ID        int       `json:"id"`
	Name	  string	`json:"name"`
	Category  string	`json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain games.Domain) Games {
	return Games{
		ID:        domain.ID,
		Name:      domain.Name,
		Category:  domain.Category,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
