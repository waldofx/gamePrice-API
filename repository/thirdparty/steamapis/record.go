package steamapis

import (
	"gameprice-api/business/steamapis"
)

type SteamName struct {
	AppID string `json:"appid"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Logo  string `json:"logo"`
}

type SteamAPI struct {
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
}

func (rec *SteamName) ToDomain() steamapis.Domain{
	return steamapis.Domain{
		AppID: rec.AppID,
		Name: rec.Name,
	}
}