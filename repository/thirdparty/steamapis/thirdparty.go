package steamapis

import (
	"encoding/json"
	"fmt"
	steamapis "gameprice-api/business/steamapis"
	"io"
	"net/http"
	"net/url"
)

type Steam struct {
	Client http.Client
}

func NewRepo() *Steam{
	return &Steam{
		Client: http.Client{},
	}
}

//third-party
func (steam *Steam) GetID(name string) (steamapis.Domain, error){
	var steamname SteamName
	endpoint := "https://steamcommunity.com/actions/SearchApps/"
	req, err := http.NewRequest("GET", endpoint+url.QueryEscape(name), nil)
	if err != nil {
		return steamapis.Domain{}, err
	}
	fmt.Println("test1")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return steamapis.Domain{}, err
	}
	fmt.Println("test2")

	bodybytes, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println("test3")

	json.Unmarshal(bodybytes, &steamname)
	return steamname.ToDomainID(), nil
}

//get data price
func (steam *Steam) GetData(appid string) (steamapis.Domain, error){
	var steamapi SteamAPI
	endpoint := "https://store.steampowered.com/api/appdetails?"
	//filters := "filters=basic,price_overview&appids="
	filters := "filters=price_overview&appids="
	req, err := http.NewRequest("GET", endpoint+filters+appid, nil)
	if err != nil {
		return steamapis.Domain{}, err
	}
	fmt.Println("test4")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return steamapis.Domain{}, err
	}
	fmt.Println("test5")

	bodybytes, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	json.Unmarshal(bodybytes, &steamapi)
	return steamapi.ToDomain(), nil
}