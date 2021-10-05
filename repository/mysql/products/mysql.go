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

func (repo *repoProducts) Insert(product *products.Domain) (*products.Domain, error) {
	recordProduct := FromDomain(*product)
	if err := repo.DBConn.Create(&recordProduct).Error; err != nil {
		return &products.Domain{}, err
	}

	record, err := repo.FindByID(int(recordProduct.ID))
	if err != nil {
		return &products.Domain{}, err
	}
	return record, nil
}

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

func (repo *repoProducts) GetProduct(gameid, sellerid int) (int, string, bool, string){
	var recordProduct Products

	if err := repo.DBConn.Joins("Game").Joins("Seller").Where("products.game_id = ?", gameid).Where("products.seller_id = ?", sellerid).Find(&recordProduct).Error; err != nil {
		return 0, "Price is not available", false, "URL is not available"
	}
	//id, price, discount, url := 
	return int(recordProduct.ID), recordProduct.Price, recordProduct.Discount, recordProduct.URL
}