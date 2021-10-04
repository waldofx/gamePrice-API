package request

import (
	"gameprice-api/business/sellers"
)

type Sellers struct {
	ID			int    	`json:"id"`
	Name    	string 	`json:"name"`
	URL    		string 	`json:"url"`
}

func ToDomain(request Sellers) *sellers.Domain {
	return &sellers.Domain{
		ID: request.ID,
		Name: request.Name,
		URL: request.URL,
	}
}