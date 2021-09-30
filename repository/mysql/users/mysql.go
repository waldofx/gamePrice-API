package users

import (
	"gameprice-api/business/users"

	"gorm.io/gorm"
)

type repoUsers struct {
	DBConn *gorm.DB
}

func NewRepoMySQL(db *gorm.DB) users.Repository {
	return &repoUsers{
		DBConn: db,
	}
}

func (repo *repoUsers) Insert(user *users.Domain) (*users.Domain, error) {
	recordUser := FromDomain(*user)
	if err := repo.DBConn.Create(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}

	record, err := repo.FindByID(int(recordUser.ID))
	if err != nil {
		return &users.Domain{}, err
	}
	return record, nil
}

func (repo *repoUsers) Update(user *users.Domain, id int) (*users.Domain, error) {
	recordUser := FromDomainUpdate(*user)
	if err := repo.DBConn.Where("id=?", id).Updates(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}

	record, err := repo.FindByID(int(id))
	if err != nil {
		return &users.Domain{}, err
	}
	return record, nil
}

func (repo *repoUsers) FindByID(id int) (*users.Domain, error) {
	var recordUser Users

	if err := repo.DBConn.Where("users.id = ?", id).Find(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}
	result := ToDomain(recordUser)
	return &result, nil
}

func (repo *repoUsers) FindAll() ([]users.Domain, error) {
	var recordUser []Users
	if err := repo.DBConn.Find(&recordUser).Error; err != nil{
		return []users.Domain{}, err
	}
	return ToDomainArray(recordUser), nil
}

func (repo *repoUsers) Delete(user *users.Domain, id int) (string, error) {
	recordUser := FromDomainUpdate(*user)
	if err := repo.DBConn.Delete(&recordUser).Error; err != nil{
		return "", err
	}
	return "Data Deleted.", nil
}