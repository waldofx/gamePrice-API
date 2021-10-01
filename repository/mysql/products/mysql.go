package products

import (
	"gameprice-api/business/products"
	"gameprice-api/business/steamapis"

	"gorm.io/gorm"
)

type repoProducts struct {
	DBConn *gorm.DB
	RepoAPI      steamapis.Repository
}

func NewRepoMySQL(db *gorm.DB) products.Repository {
	return &repoProducts{
		DBConn: db,
	}
}

// func (repo *repoProducts) Insert(product *products.Domain) (*products.Domain, error) {
// 	recordProduct := FromDomain(*product)
// 	if err := repo.DBConn.Create(&recordProduct).Error; err != nil {
// 		return &products.Domain{}, err
// 	}

// 	record, err := repo.FindByID(int(recordProduct.ID))
// 	if err != nil {
// 		return &products.Domain{}, err
// 	}
// 	return record, nil
// }

func (repo *repoProducts) Update(product *products.Domain, id int) (*products.Domain, error) {
	recordProduct := FromDomain(*product)
	if err := repo.DBConn.Where("id=?", id).Updates(&recordProduct).Error; err != nil {
		return &products.Domain{}, err
	}

	record, err := repo.FindByID(int(id))
	if err != nil {
		return &products.Domain{}, err
	}
	return record, nil
}

func (repo *repoProducts) FindByID(id int) (*products.Domain, error) {
	var recordProduct Products

	if err := repo.DBConn.Where("products.id = ?", id).Joins("Game").Joins("Seller").Find(&recordProduct).Error; err != nil {
		return &products.Domain{}, err
	}
	result := ToDomain(recordProduct)
	return &result, nil
}

func (repo *repoProducts) FindAll() ([]products.Domain, error) {
	var recordProduct []Products
	if err := repo.DBConn.Joins("Game").Joins("Seller").Find(&recordProduct).Error; err != nil{
		return []products.Domain{}, err
	}
	return ToDomainArray(recordProduct), nil
}

func (repo *repoProducts) Delete(product *products.Domain, id int) (string, error) {
	recordProduct := FromDomain(*product)
	if err := repo.DBConn.Delete(&recordProduct).Error; err != nil{
		return "", err
	}
	return "Data Deleted.", nil
}

//GET PRICE
func (repo *repoProducts) Insert(product *products.Domain) (*products.Domain, error) {
	recordProduct := FromDomain(*product)
	if err := repo.DBConn.Create(&recordProduct).Error; err != nil {
		return &products.Domain{}, err
	}

	// appid, err := repo.RepoAPI.GetID(recordProduct.Game.Name)
	// if err != nil {
	// 	return &products.Domain{}, err
	// }
	// newprice, err := repo.RepoAPI.GetData(appid.AppID)
	// if err != nil {
	// 	return &products.Domain{}, err
	// }
	// recordProduct.Price = newprice.Price

	record, err := repo.FindByID(int(recordProduct.ID))
	if err != nil {
		return &products.Domain{}, err
	}
	return record, nil
}

// //GET PRICE
// func (repo *repoProducts) Insert(product *products.Domain) (*products.Domain, error) {
// 	recordProduct := FromDomain(*product)
// 	if err := repo.DBConn.Create(&recordProduct).Error; err != nil {
// 		return &products.Domain{}, err
// 	}

// 	newprice, err := repo.GetPrice(recordProduct.Game.Name)
// 	if err != nil {
// 		return &products.Domain{}, err
// 	}
// 	recordProduct.Price = newprice

// 	record, err := repo.FindByID(int(recordProduct.ID))
// 	if err != nil {
// 		return &products.Domain{}, err
// 	}
// 	return record, nil
// }

// //third-party
// func (repo *repoProducts) GetPrice(name string) (int, error){
// 	var steamname SteamName
// 	endpoint := "https://steamcommunity.com/actions/SearchApps/"
// 	resp, err := http.NewRequest("GET", endpoint+url.QueryEscape(name), nil)
// 	if err != nil {
// 		return 0, err
// 	}

// 	bodybytes, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return 0, err
// 	}

// 	json.Unmarshal(bodybytes, &steamname)
// 	appid := steamname.Appid
// 	price, err := repo.GetData(appid)
// 	defer resp.Body.Close()
// 	return price, err
// }

// //get data price
// func (repo *repoProducts) GetData(appid string) (int, error){
// 	var steamapi SteamAPI
// 	endpoint := "https://store.steampowered.com/api/appdetails?"
// 	filters := "filters=basic,price_overview&appids="
// 	resp, err := http.NewRequest("GET", endpoint+filters+appid, nil)
// 	if err != nil {
// 		return 0, err
// 	}

// 	bodybytes, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return 0, err
// 	}

// 	json.Unmarshal(bodybytes, &steamapi)
// 	price := steamapi.Num1238810.Data.PriceOverview.Final
// 	defer resp.Body.Close()
// 	return price, err
// }