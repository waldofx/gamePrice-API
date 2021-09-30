package users

import (
	"fmt"
	"gameprice-api/business/users"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       	uint `gorm:"primaryKey"`
	Name    	string
	Email  	  	string
	Password  	string
}

func ToDomain(rec Users) users.Domain {
	fmt.Printf("%+v", rec)
	return users.Domain{
		ID:        	int(rec.ID),
		Name:      	rec.Name,
		Email:  	rec.Email,
		Password:	rec.Password,
		CreatedAt: 	rec.CreatedAt,
		UpdatedAt: 	rec.UpdatedAt,
	}
}

func FromDomain(domain users.Domain) Users {
	return Users{
		ID: 		uint(domain.ID),
		Name:    	domain.Name,
		Email:  	domain.Email,
		Password:	domain.Password,
	}
}

func ToDomainArray(modelUsers []Users) []users.Domain{
	var response []users.Domain

	for _, val := range modelUsers{
		response = append(response, ToDomain(val))
	}
	return response
}