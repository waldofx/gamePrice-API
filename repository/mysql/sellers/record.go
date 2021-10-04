package sellers

import (
	"gameprice-api/business/sellers"

	"gorm.io/gorm"
)

type Sellers struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Name      string
	URL  	  string
}

func ToDomain(rec Sellers) sellers.Domain {
	return sellers.Domain{
		ID:        int(rec.ID),
		Name:      rec.Name,
		URL:  	   rec.URL,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func FromDomain(domain sellers.Domain) Sellers {
	return Sellers{
		ID: uint(domain.ID),
		Name:    domain.Name,
		URL:     domain.URL,
	}
}

func ToDomainArray(modelSellers []Sellers) []sellers.Domain{
	var response []sellers.Domain

	for _, val := range modelSellers{
		response = append(response, ToDomain(val))
	}
	return response
}