package request

import (
	"gameprice-api/business/users"
)

type UserInsert struct {
	Name   		string	`json:"name"`
	Email  	  	string	`json:"email"`
	Password	string	`json:"password"`
}

type UserUpdate struct {
	ID 			int		`json:"id"`
	Name   		string	`json:"name"`
	Email  	  	string	`json:"email"`
	Password	string	`json:"password"`
}

func ToDomain(request UserInsert) *users.Domain {
	return &users.Domain{
		Name: request.Name,
		Email: request.Email,
		Password: request.Password,
	}
}

func ToDomainUpdate(request UserUpdate) *users.Domain {
	return &users.Domain{
		ID: request.ID,
		Name: request.Name,
		Email: request.Email,
		Password: request.Password,
	}
}