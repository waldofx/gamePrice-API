package request

import (
	"gameprice-api/business/games"
)

type Games struct {
	ID			int    `json:"id"`
	Name    	string `json:"name"`
	Category    string `json:"category"`
}

func ToDomain(request Games) *games.Domain {
	return &games.Domain{
		ID: request.ID,
		Name: request.Name,
		Category: request.Category,
	}
}