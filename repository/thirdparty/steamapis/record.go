package steamapis

import (
	"fmt"
	"gameprice-api/business/steamapis"
)

type SteamName struct {
	AppID string `json:"appid"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Logo  string `json:"logo"`
}

type SteamAPI struct { //[error] invalid field found for struct ()'s field PriceOverview: define a valid foreign key for relations or implement the Valuer/Scanner interface 

	Num1238810 struct {
		Success bool `json:"success"`
		Data    struct {
			PriceOverview struct {
				Currency         string `json:"currency"`
				Initial          int    `json:"initial"`
				Final            int    `json:"final"`
				DiscountPercent  int    `json:"discount_percent"`
				InitialFormatted string `json:"initial_formatted"`
				FinalFormatted   string `json:"final_formatted"`
			} `json:"price_overview,omitempty"`
		} `json:"data"`
	} `json:"1238810"`
}

func (rec *SteamName) ToDomainID() steamapis.Domain{
	fmt.Printf("%+v", rec)
	return steamapis.Domain{
		AppID: rec.AppID,
		Name: rec.Name,
	}
}

func (rec *SteamAPI) ToDomain() steamapis.Domain {
	fmt.Printf("%+v", rec)
	return steamapis.Domain{
		Price:     rec.Num1238810.Data.PriceOverview.Final,
	}
}