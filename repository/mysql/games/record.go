package games

import (
	"fmt"
	"gameprice-api/business/games"

	"gorm.io/gorm"
)

type Games struct {
	gorm.Model
	Name      string
	Category  string
}

func toDomain(rec Games) games.Domain {
	fmt.Printf("%+v", rec)
	return games.Domain{
		ID:        int(rec.ID),
		Name:      rec.Name,
		Category:  rec.Category,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(domain games.Domain) Games {
	return Games{
		Name:    domain.Name,
		Category:     domain.Category,
	}
}