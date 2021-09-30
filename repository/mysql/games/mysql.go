package games

import (
	"gameprice-api/business/games"

	"gorm.io/gorm"
)

type repoGames struct {
	DBConn *gorm.DB
}

func NewRepoMySQL(db *gorm.DB) games.Repository {
	return &repoGames{
		DBConn: db,
	}
}

func (repo *repoGames) Insert(game *games.Domain) (*games.Domain, error) {
	recordGame := FromDomain(*game)
	if err := repo.DBConn.Create(&recordGame).Error; err != nil {
		return &games.Domain{}, err
	}

	record, err := repo.FindByID(int(recordGame.ID))
	if err != nil {
		return &games.Domain{}, err
	}
	return record, nil
}

func (repo *repoGames) Update(game *games.Domain, id int) (*games.Domain, error) {
	recordGame := FromDomain(*game)
	if err := repo.DBConn.Where("id=?", id).Updates(&recordGame).Error; err != nil {
		return &games.Domain{}, err
	}

	record, err := repo.FindByID(int(recordGame.ID))
	if err != nil {
		return &games.Domain{}, err
	}
	return record, nil
}

func (repo *repoGames) FindByID(id int) (*games.Domain, error) {
	var recordGame Games

	if err := repo.DBConn.Where("games.id = ?", id).Find(&recordGame).Error; err != nil {
		return &games.Domain{}, err
	}
	result := ToDomain(recordGame)
	return &result, nil
}

func (repo *repoGames) FindAll() ([]games.Domain, error) {
	var recordGame []Games
	if err := repo.DBConn.Find(&recordGame).Error; err != nil{
		return []games.Domain{}, err
	}
	return ToDomainArray(recordGame), nil
}

func (repo *repoGames) Delete(game *games.Domain, id int) (string, error) {
	recordGame := FromDomain(*game)
	if err := repo.DBConn.Delete(&recordGame).Error; err != nil{
		return "", err
	}
	return "Data Deleted.", nil
}