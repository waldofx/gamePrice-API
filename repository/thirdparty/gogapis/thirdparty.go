package gogapis

import (
	"encoding/json"
	"fmt"
	gogapis "gameprice-api/business/gogapis"
	"net/http"
)

type RepoGog struct {
	httpClient http.Client
}

func NewRepo() gogapis.Repository{
	return &RepoGog{
		httpClient: http.Client{},
	}
}

func (gog *RepoGog) GetData(appid string) (gogapis.Domain, error){
	var gogapi GOG
	// var appid int
	// appid,_ = strconv.Atoi(GOGID)
	endpoint := "api.gog.com/products/"
	req, err := http.NewRequest("GET", endpoint+appid, nil)
	if err != nil {
		return gogapis.Domain{}, err
	}
	fmt.Println("GOG GetData, step 1") //debug
	fmt.Println(endpoint+appid) //debug

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return gogapis.Domain{}, err
	}
	fmt.Println("GOG GetData, step 2") //debug
	fmt.Println(resp) //debug
	fmt.Println(resp.Body) //debug

	defer resp.Body.Close()
	//bodybytes, _ := io.ReadAll(resp.Body)
	//json.Unmarshal(bodybytes, &gogapi)

	err = json.NewDecoder(resp.Body).Decode(&gogapi)
	fmt.Println("GOG GetData, step 3") //debug
	fmt.Println(err, gogapi) //debug
	if err != nil {
		return gogapis.Domain{}, err
	}
	return gogapi.ToDomain(), nil
}