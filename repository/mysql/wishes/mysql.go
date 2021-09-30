package wishes

import (
	"gameprice-api/business/wishes"

	"gorm.io/gorm"
)

type repoWishes struct {
	DBConn *gorm.DB
}

func NewRepoMySQL(db *gorm.DB) wishes.Repository {
	return &repoWishes{
		DBConn: db,
	}
}

func (repo *repoWishes) Insert(wish *wishes.Domain) (*wishes.Domain, error) {
	recordWish := FromDomain(*wish)
	if err := repo.DBConn.Create(&recordWish).Error; err != nil {
		return &wishes.Domain{}, err
	}

	record, err := repo.FindByID(int(recordWish.ID))
	if err != nil {
		return &wishes.Domain{}, err
	}
	return record, nil
}

func (repo *repoWishes) Update(wish *wishes.Domain, id int) (*wishes.Domain, error) {
	recordWish := FromDomain(*wish)
	if err := repo.DBConn.Where("id=?", id).Updates(&recordWish).Error; err != nil {
		return &wishes.Domain{}, err
	}

	record, err := repo.FindByID(int(id))
	if err != nil {
		return &wishes.Domain{}, err
	}
	return record, nil
}

func (repo *repoWishes) FindByID(id int) (*wishes.Domain, error) {
	var recordWish Wishes

	if err := repo.DBConn.Where("wishes.id = ?", id).Find(&recordWish).Error; err != nil {
		return &wishes.Domain{}, err
	}
	result := ToDomain(recordWish)
	return &result, nil
}

func (repo *repoWishes) FindAll() ([]wishes.Domain, error) {
	var recordWish []Wishes
	if err := repo.DBConn.Find(&recordWish).Error; err != nil{
		return []wishes.Domain{}, err
	}
	return ToDomainArray(recordWish), nil
}

func (repo *repoWishes) Delete(wish *wishes.Domain, id int) (string, error) {
	recordWish := FromDomain(*wish)
	if err := repo.DBConn.Delete(&recordWish).Error; err != nil{
		return "", err
	}
	return "Data Deleted.", nil
}