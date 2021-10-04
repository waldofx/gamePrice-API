package response

import (
	"gameprice-api/business/games"
	"time"
)

type Games struct {
	ID        	int       	`json:"id"`
	Name	  	string		`json:"name"`
	SteamID   	string 		`json:"steam_id"`
	GOGID	  	string		`json:"gog_id"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}

func FromDomain(domain games.Domain) Games {
	return Games{
		ID:       	domain.ID,
		Name:     	domain.Name,
		SteamID:  	domain.SteamID,
		GOGID: 		domain.GOGID,
		CreatedAt:	domain.CreatedAt,
		UpdatedAt:	domain.UpdatedAt,
	}
}

func NewResponseArray(modelGames []games.Domain) []Games{
	var response []Games

	for _, val := range modelGames{
		response = append(response, FromDomain(val))
	}
	return response
}