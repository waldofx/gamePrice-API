package request

import (
	"gameprice-api/business/sellers"
)

type Sellers struct {
	Name    	string 	`json:"name"`
	URL    		string 	`json:"url"`
}

func ToDomain(request Sellers) *sellers.Domain {
	return &sellers.Domain{
		Name: request.Name,
		URL: request.URL,
	}
}