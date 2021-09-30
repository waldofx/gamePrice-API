package request

import (
	"gameprice-api/business/sellers"
)

type SellerInsert struct {
	Name    	string 	`json:"name"`
	URL    		string 	`json:"url"`
}

type SellerUpdate struct {
	ID 			int 	`json:"id"`
	Name    	string 	`json:"name"`
	URL    		string 	`json:"url"`
}

func ToDomain(request SellerInsert) *sellers.Domain {
	return &sellers.Domain{
		Name: request.Name,
		URL: request.URL,
	}
}

func ToDomainUpdate(request SellerUpdate) *sellers.Domain {
	return &sellers.Domain{
		ID: request.ID,
		Name: request.Name,
		URL: request.URL,
	}
}