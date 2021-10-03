package request

import (
	"gameprice-api/business/users"
)

type Users struct {
	Username 	string `json:"username"`
	Email  	  	string	`json:"email"`
	Password	string	`json:"password"`
}

func ToDomain(request Users) *users.Domain {
	return &users.Domain{
		Username: request.Username,
		Email: request.Email,
		Password: request.Password,
	}
}