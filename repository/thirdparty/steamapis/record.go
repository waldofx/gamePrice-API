package steamapis

import (
	"fmt"
	steamapis "gameprice-api/business/steamapi"
)

type SteamName struct {
	Appid string `json:"appid"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Logo  string `json:"logo"`
}

type SteamAPI struct {
	Num1238810 struct {
		Success bool `json:"success"`
		Data    struct {
			Type                string `json:"type"`
			Name                string `json:"name"`
			SteamAppid          int    `json:"steam_appid"`
			RequiredAge         int    `json:"required_age"`
			IsFree              bool   `json:"is_free"`
			ControllerSupport   string `json:"controller_support"`
			Dlc                 []int  `json:"dlc"`
			DetailedDescription string `json:"detailed_description"`
			AboutTheGame        string `json:"about_the_game"`
			ShortDescription    string `json:"short_description"`
			SupportedLanguages  string `json:"supported_languages"`
			HeaderImage         string `json:"header_image"`
			Website             string `json:"website"`
			PcRequirements      struct {
				Minimum     string `json:"minimum"`
				Recommended string `json:"recommended"`
			} `json:"pc_requirements"`
			MacRequirements struct {
				Minimum     string `json:"minimum"`
				Recommended string `json:"recommended"`
			} `json:"mac_requirements"`
			LinuxRequirements struct {
				Minimum     string `json:"minimum"`
				Recommended string `json:"recommended"`
			} `json:"linux_requirements"`
			LegalNotice          string `json:"legal_notice"`
			DrmNotice            string `json:"drm_notice"`
			ExtUserAccountNotice string `json:"ext_user_account_notice"`
			PriceOverview        struct {
				Currency         string `json:"currency"`
				Initial          int    `json:"initial"`
				Final            int    `json:"final"`
				DiscountPercent  int    `json:"discount_percent"`
				InitialFormatted string `json:"initial_formatted"`
				FinalFormatted   string `json:"final_formatted"`
			} `json:"price_overview"`
		} `json:"data"`
	} `json:"1238810"`
}

func (rec *SteamName) ToDomainID() steamapis.Domain{
	fmt.Printf("%+v", rec)
	return steamapis.Domain{
		AppID: rec.Appid,
	}
}

func (rec *SteamAPI) ToDomain() steamapis.Domain {
	fmt.Printf("%+v", rec)
	return steamapis.Domain{
		Name:  	   rec.Num1238810.Data.Name,
		Price:     rec.Num1238810.Data.PriceOverview.Final,
	}
}