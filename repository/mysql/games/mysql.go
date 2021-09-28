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
	recordGame := fromDomain(*game)
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
	return &games.Domain{}, nil
}
func (repo *repoGames) FindByID(id int) (*games.Domain, error) {
	var recordGame Games

	if err := repo.DBConn.Where("games.id = ?", id).Joins("Author").Find(&recordGame).Error; err != nil {
		return &games.Domain{}, err
	}
	result := toDomain(recordGame)
	return &result, nil
}
func (repo *repoGames) FindAll(generalSearch string, availability bool) []games.Domain {
	return []games.Domain{}
}
