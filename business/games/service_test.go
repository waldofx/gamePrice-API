package games_test

import (
	businesses "gameprice-api/business"
	"gameprice-api/business/games"
	_gamesMock "gameprice-api/business/games/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	gamesRepository _gamesMock.Repository
	gamesService	games.Service
	gamesDomain games.Domain
)

func TestMain(m *testing.M) {
	gamesService = games.NewService(&gamesRepository)
	gamesDomain = games.Domain{
		ID:          1,
		Name:       "Test Games",
		SteamID: 	"000000",
		GOGID:		"00000",
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
	}
}

func TestAppend(t *testing.T){
	t.Run("Append | Valid", func(t *testing.T) {
		gamesRepository.On("Insert", mock.AnythingOfType("games.Domain")).Return(gamesDomain, nil).Once()

		result, err := gamesService.Append(&gamesDomain)

		assert.Nil(t, err)
		assert.Equal(t, gamesDomain, result)
	})

	t.Run("Append | InValid", func(t *testing.T) {
		gamesRepository.On("Insert", mock.AnythingOfType("games.Domain")).Return(gamesDomain, businesses.ErrInternalServer).Once()

		_, err := gamesService.Append(&gamesDomain)

		assert.NotNil(t, err)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("Find All | Valid", func(t *testing.T) {
		gamesRepository.On("FindAll").Return([]games.Domain{gamesDomain}, nil).Once()

		result, err := gamesService.FindAll()

		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("Find All | InValid", func(t *testing.T) {
		gamesRepository.On("Find").Return([]games.Domain{}, businesses.ErrCategoryNotFound).Once()

		_, err := gamesService.FindAll()

		assert.NotNil(t, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("Find By ID | Valid", func(t *testing.T) {
		gamesRepository.On("FindByID", mock.AnythingOfType("games.Domain")).Return(gamesDomain, nil).Once()

		result, err := gamesService.FindByID(gamesDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, gamesDomain, result)
	})

	t.Run("Find By ID | InValid", func(t *testing.T) {
		gamesRepository.On("FindByID", mock.AnythingOfType("games.Domain")).Return(gamesDomain, businesses.ErrCategoryNotFound).Once()

		_, err := gamesService.FindByID(gamesDomain.ID)

		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		gamesRepository.On("Update", mock.AnythingOfType("games.Domain"), mock.AnythingOfType("int")).Return(gamesDomain, nil).Once()

		result, err := gamesService.Update(&gamesDomain, gamesDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, gamesDomain, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		gamesRepository.On("Update", mock.AnythingOfType("games.Domain"), mock.AnythingOfType("int")).Return(gamesDomain, businesses.ErrInternalServer).Once()

		_, err := gamesService.Update(&gamesDomain, gamesDomain.ID)

		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		gamesRepository.On("Delete", mock.AnythingOfType("games.Domain"), mock.AnythingOfType("int")).Return(mock.AnythingOfType("string"), nil).Once()

		result, err := gamesService.Delete(&gamesDomain, gamesDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, "Data Deleted.", result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		gamesRepository.On("Detele", mock.AnythingOfType("games.Domain"), mock.AnythingOfType("int")).Return(mock.AnythingOfType("string"), businesses.ErrInternalServer).Once()

		_, err := gamesService.Delete(&gamesDomain, gamesDomain.ID)

		assert.NotNil(t, err)
	})
}