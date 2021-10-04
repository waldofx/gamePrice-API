package request

import (
	"gameprice-api/business/games"
)

type Games struct {
	ID			int    	`json:"id"`
	Name    	string 	`json:"name"`
	SteamID    	string 	`json:"steam_id"`
	GOGID		string	`json:"gog_id"`
}

func ToDomain(request Games) *games.Domain {
	return &games.Domain{
		ID:			request.ID,
		Name: 		request.Name,
		SteamID: 	request.SteamID,
		GOGID: 		request.GOGID,
	}
}