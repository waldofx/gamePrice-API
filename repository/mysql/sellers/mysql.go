package sellers

import (
	"gameprice-api/business/sellers"

	"gorm.io/gorm"
)

type repoSellers struct {
	DBConn *gorm.DB
}

func NewRepoMySQL(db *gorm.DB) sellers.Repository {
	return &repoSellers{
		DBConn: db,
	}
}

func (repo *repoSellers) Insert(seller *sellers.Domain) (*sellers.Domain, error) {
	recordSeller := FromDomain(*seller)
	if err := repo.DBConn.Create(&recordSeller).Error; err != nil {
		return &sellers.Domain{}, err
	}

	record, err := repo.FindByID(int(recordSeller.ID))
	if err != nil {
		return &sellers.Domain{}, err
	}
	return record, nil
}

func (repo *repoSellers) Update(seller *sellers.Domain, id int) (*sellers.Domain, error) {
	recordSeller := FromDomain(*seller)
	if err := repo.DBConn.Where("id=?", id).Updates(&recordSeller).Error; err != nil {
		return &sellers.Domain{}, err
	}

	record, err := repo.FindByID(int(id))
	if err != nil {
		return &sellers.Domain{}, err
	}
	return record, nil
}

func (repo *repoSellers) FindByID(id int) (*sellers.Domain, error) {
	var recordSeller Sellers

	if err := repo.DBConn.Where("sellers.id = ?", id).Find(&recordSeller).Error; err != nil {
		return &sellers.Domain{}, err
	}
	result := ToDomain(recordSeller)
	return &result, nil
}

func (repo *repoSellers) FindAll() ([]sellers.Domain, error) {
	var recordSeller []Sellers
	if err := repo.DBConn.Find(&recordSeller).Error; err != nil{
		return []sellers.Domain{}, err
	}
	return ToDomainArray(recordSeller), nil
}

func (repo *repoSellers) Delete(seller *sellers.Domain, id int) (string, error) {
	recordSeller := FromDomain(*seller)
	if err := repo.DBConn.Delete(&recordSeller).Error; err != nil{
		return "", err
	}
	return "Data Deleted.", nil
}