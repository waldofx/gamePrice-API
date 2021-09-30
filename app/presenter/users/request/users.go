package request

import (
	"gameprice-api/business/users"
)

type Users struct {
	Name   		string	`json:"name"`
	Email  	  	string	`json:"email"`
	Password	string	`json:"password"`
}

func ToDomain(request Users) *users.Domain {
	return &users.Domain{
		Name: request.Name,
		Email: request.Email,
		Password: request.Password,
	}
}