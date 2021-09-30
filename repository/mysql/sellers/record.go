package sellers

import (
	"fmt"
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
	fmt.Printf("%+v", rec)
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
		Name:    domain.Name,
		URL:     domain.URL,
	}
}


func FromDomainUpdate(domain sellers.Domain) Sellers {
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