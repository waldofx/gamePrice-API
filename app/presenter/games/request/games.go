package request

import (
	"gameprice-api/business/games"
)

type GameInsert struct {
	Name    	string `json:"name"`
	Category    string `json:"category"`
}

type GameUpdate struct {
	ID 			int `json:"id"`
	Name    	string `json:"name"`
	Category    string `json:"category"`
}

func ToDomain(request GameInsert) *games.Domain {
	return &games.Domain{
		Name: request.Name,
		Category: request.Category,
	}
}

func ToDomainUpdate(request GameUpdate) *games.Domain {
	return &games.Domain{
		ID: request.ID,
		Name: request.Name,
		Category: request.Category,
	}
}