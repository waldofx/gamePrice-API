package steamapis

import (
	"encoding/json"
	steamapis "gameprice-api/business/steamapi"
	"io"
	"net/http"
	"net/url"
)

type steam struct {
	Client http.Client
}

//third-party
func (steam *steam) GetID(name string) (steamapis.Domain, error){
	var steamname SteamName
	endpoint := "https://steamcommunity.com/actions/SearchApps/"
	req, err := http.NewRequest("GET", endpoint+url.QueryEscape(name), nil)
	if err != nil {
		return steamapis.Domain{}, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return steamapis.Domain{}, err
	}

	defer resp.Body.Close()
	bodybytes, _ := io.ReadAll(resp.Body)

	json.Unmarshal(bodybytes, &steamname)
	return steamname.ToDomainID(), nil
}

//get data price
func (steam *steam) GetData(appid string) (steamapis.Domain, error){
	var steamapi SteamAPI
	endpoint := "https://store.steampowered.com/api/appdetails?"
	filters := "filters=basic,price_overview&appids="
	req, err := http.NewRequest("GET", endpoint+filters+appid, nil)
	if err != nil {
		return steamapis.Domain{}, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return steamapis.Domain{}, err
	}

	defer resp.Body.Close()
	bodybytes, _ := io.ReadAll(resp.Body)

	json.Unmarshal(bodybytes, &steamapi)
	return steamapi.ToDomain(), nil
}