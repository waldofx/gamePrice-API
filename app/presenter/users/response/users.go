package response

import (
	"gameprice-api/business/users"
	"time"
)

type Users struct {
	ID        	int       	`json:"id"`
	Name   		string		`json:"name"`
	Username	string		`json:"username"`
	Email  	  	string		`json:"email"`
	Password	string		`json:"password"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}

func FromDomain(domain users.Domain) Users {
	return Users{
		ID:        	domain.ID,
		Name:      	domain.Name,
		Username: 	domain.Username,
		Email: 		domain.Email,
		Password:	domain.Password,
		CreatedAt: 	domain.CreatedAt,
		UpdatedAt: 	domain.UpdatedAt,
	}
}

func NewResponseArray(modelUsers []users.Domain) []Users{
	var response []Users

	for _, val := range modelUsers{
		response = append(response, FromDomain(val))
	}
	return response
}