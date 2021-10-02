package steamapis

import (
	"encoding/json"
	"fmt"
	steamapis "gameprice-api/business/steamapis"
	"net/http"
	"net/url"
	"strings"
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
func (steam *Steam) GetID(gname string) (steamapis.Domain, error){
	name := strings.ToLower(gname)
	endpoint := "https://steamcommunity.com/actions/SearchApps/"
	req, err := http.NewRequest("GET", endpoint+url.QueryEscape(name), nil)
	if err != nil {
		return steamapis.Domain{}, err
	}
	fmt.Println(endpoint+url.QueryEscape(name)) //debug

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return steamapis.Domain{}, err
	}
	fmt.Println(resp) //debug
	fmt.Println(resp.Body) //debug
	fmt.Println("Finish respponse!") //debug

	defer resp.Body.Close()
	// bodybytes, _ := io.ReadAll(resp.Body)
	// json.Unmarshal(bodybytes, &steamname)

	steamname := SteamName{}
	err = json.NewDecoder(resp.Body).Decode(&steamname)
	fmt.Println(err, steamname) //debug
	if err != nil {
		return steamapis.Domain{}, err
	}
	fmt.Println(" GetID sucess! 1") //debug
	return steamname.ToDomainID(), nil
}

//get data price
func (steam *Steam) GetData(appid string) (steamapis.Domain, error){
	//var steamapi SteamAPI
	endpoint := "https://store.steampowered.com/api/appdetails?"
	//filters := "filters=basic,price_overview&appids="
	filters := "filters=price_overview&appids="
	req, err := http.NewRequest("GET", endpoint+filters+appid, nil)
	if err != nil {
		return steamapis.Domain{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return steamapis.Domain{}, err
	}

	defer resp.Body.Close()
	//bodybytes, _ := io.ReadAll(resp.Body)
	//json.Unmarshal(bodybytes, &steamapi)

	steamapi := SteamAPI{}
	err = json.NewDecoder(resp.Body).Decode(&steamapi)
	fmt.Println(err, steamapi) //debug
	if err != nil {
		return steamapis.Domain{}, err
	}
	return steamapi.ToDomain(), nil
}