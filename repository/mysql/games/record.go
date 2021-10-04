package games

import (
	"gameprice-api/business/games"

	"gorm.io/gorm"
)

type Games struct {
	gorm.Model
	ID        	uint `gorm:"primaryKey"`
	Name      	string
	SteamID  	string
	GOGID		string
}

func ToDomain(rec Games) games.Domain {
	// fmt.Printf("%+v", rec)
	return games.Domain{
		ID:        int(rec.ID),
		Name:      rec.Name,
		SteamID:   rec.SteamID,
		GOGID:     rec.GOGID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func FromDomain(domain games.Domain) Games {
	return Games{
		ID: 		uint(domain.ID),
		Name:    	domain.Name,
		SteamID:    domain.SteamID,
		GOGID:     	domain.GOGID,
	}
}

func ToDomainArray(modelGames []Games) []games.Domain{
	var response []games.Domain

	for _, val := range modelGames{
		response = append(response, ToDomain(val))
	}
	return response
}