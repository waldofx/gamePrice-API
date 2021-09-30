package request

import (
	"gameprice-api/business/games"
)

type Games struct {
	Name    	string `json:"name"`
	Category    string `json:"category"`
}

func ToDomain(request Games) *games.Domain {
	return &games.Domain{
		Name: request.Name,
		Category: request.Category,
	}
}