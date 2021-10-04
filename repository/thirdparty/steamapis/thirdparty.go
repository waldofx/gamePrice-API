package steamapis

import (
	"encoding/json"
	steamapis "gameprice-api/business/steamapis"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Steam struct {
	httpClient http.Client
}

func NewRepo() steamapis.Repository{
	return &Steam{
		httpClient: http.Client{},
	}
}

//third-party
func (steam *Steam) GetID(gname string) (steamapis.Domain, error){
	//var steamname SteamName
	name := strings.ToLower(gname)
	endpoint := "https://steamcommunity.com/actions/SearchApps/"
	req, err := http.NewRequest("GET", endpoint+url.QueryEscape(name), nil)
	if err != nil {
		return steamapis.Domain{}, err
	}

	resp, err := steam.httpClient.Do(req)
	if err != nil {
		return steamapis.Domain{}, err
	}

	defer resp.Body.Close()
	bodybytes, _ := io.ReadAll(resp.Body)
	var steamname []SteamName
	_ = json.Unmarshal(bodybytes, &steamname)

	if err != nil {
		return steamapis.Domain{}, err
	}
	return steamname[0].ToDomain(), nil
}

//get data price
func (steam *Steam) GetData(appid string) (steamapis.Domain, error){
	client := &http.Client{}
	endpoint := "https://store.steampowered.com/api/appdetails?"
	filters := "filters=price_overview&appids="
	req, err := http.NewRequest("GET", endpoint+filters+appid, nil)
	if err != nil {
		return steamapis.Domain{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return steamapis.Domain{}, err
	}

	defer resp.Body.Close()
	bodybytes, _ := ioutil.ReadAll(resp.Body)
	var steamapi map[string]SteamAPI
	_ = json.Unmarshal(bodybytes, &steamapi)


	if err != nil {
		return steamapis.Domain{}, err
	}
	if steamapi[appid].Data.PriceOverview.DiscountPercent != 0 {
		return steamapis.Domain{
			Price: steamapi[appid].Data.PriceOverview.FinalFormatted,
			Discount: true,
		}, nil
	} else{
		return steamapis.Domain{
			Price: steamapi[appid].Data.PriceOverview.FinalFormatted,
		}, nil
	}
}