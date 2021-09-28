package request

import (
	"gameprice-api/business/games"
)

type GameInsert struct {
	Name    	string `json:"name"`
	Category    string `json:"category"`
}

type GameUpdate struct {
	Name    	string `json:"name"`
	Category    string `json:"category"`
}

func ToDomain(request GameInsert) *games.Domain {
	return &games.Domain{
		Name: request.Name,
		Category: request.Category,
	}
}
