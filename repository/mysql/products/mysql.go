package products

import (
	"gameprice-api/business/products"

	"gorm.io/gorm"
)

type repoProducts struct {
	DBConn *gorm.DB
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

	record, err := repo.FindByID(int(recordProduct.ID))
	if err != nil {
		return &products.Domain{}, err
	}
	return record, nil
}

func (repo *repoProducts) FindByID(id int) (*products.Domain, error) {
	var recordProduct Products

	if err := repo.DBConn.Where("products.id = ?", id).Find(&recordProduct).Error; err != nil {
		return &products.Domain{}, err
	}
	result := ToDomain(recordProduct)
	return &result, nil
}

func (repo *repoProducts) FindAll() ([]products.Domain, error) {
	var recordProduct []Products
	if err := repo.DBConn.Find(&recordProduct).Error; err != nil{
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