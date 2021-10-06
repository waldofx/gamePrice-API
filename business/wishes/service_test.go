package wishes_test

import (
	businesses "gameprice-api/business"
	"gameprice-api/business/products"
	_productsMock "gameprice-api/business/products/mocks"
	"gameprice-api/business/wishes"
	_wishesMock "gameprice-api/business/wishes/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	wishesRepository _wishesMock.Repository
	productsRepository	_productsMock.Repository
	wishesService	wishes.Service
	wishesDomain 	wishes.Domain
	productsDomain products.Domain
)

func TestMain(m *testing.M) {
	wishesService = wishes.NewService(&wishesRepository, &productsRepository)
	wishesDomain = wishes.Domain{
		ID:         1,
		UserID: 	1,
		Username: 	"Username Test",
		GameID:		1,
		GameName:   "Game Test",
		SellerID:  	1,
		GameSeller: "Seller Test",
		Price:	 	"Rp 0",
		Discount:  	false,
		URL:		"example.com",
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
	}
	productsDomain = products.Domain{
		ID:         1,
		GameID:		1,
		Game:     	"Game Test",
		SellerID:  	1,
		Seller:    	"Seller Test",
		Price:	 	"Rp 0",
		Discount:  	false,
		URL:		"example.com",
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
	}
	os.Exit(m.Run())
}
func TestAppend(t *testing.T){
	t.Run("Append | Valid", func(t *testing.T) {
		productsRepository.On("GetProduct", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(1, "1", false, "1").Once()
		wishesRepository.On("Insert", mock.AnythingOfType("*wishes.Domain")).Return(&wishesDomain, nil).Once()

		result, err := wishesService.Append(&wishesDomain)

		assert.Nil(t, err)
		assert.Equal(t, &wishesDomain, result)
	})

	t.Run("Append | InValid", func(t *testing.T) {
		productsRepository.On("GetProduct", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(1, "1", false, "1").Once()
		wishesRepository.On("Insert", mock.AnythingOfType("*wishes.Domain")).Return(&wishesDomain, businesses.ErrInternalServer).Once()

		_, err := wishesService.Append(&wishesDomain)

		assert.NotNil(t, err)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("Find All | Valid", func(t *testing.T) {
		wishesRepository.On("FindAll").Return([]wishes.Domain{wishesDomain}, nil).Once()

		result, err := wishesService.FindAll()

		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("Find All | InValid", func(t *testing.T) {
		wishesRepository.On("FindAll").Return([]wishes.Domain{}, businesses.ErrCategoryNotFound).Once()

		_, err := wishesService.FindAll()

		assert.NotNil(t, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("Find By ID | Valid", func(t *testing.T) {
		wishesRepository.On("FindByID", mock.AnythingOfType("int")).Return(&wishesDomain, nil).Once()

		result, err := wishesService.FindByID(wishesDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, &wishesDomain, result)
	})

	t.Run("Find By ID | InValid", func(t *testing.T) {
		wishesRepository.On("FindByID", mock.AnythingOfType("int")).Return(&wishesDomain, businesses.ErrCategoryNotFound).Once()

		_, err := wishesService.FindByID(wishesDomain.ID)

		assert.NotNil(t, err)
	})
}

func TestFindUserByID(t *testing.T) {
	t.Run("Find By UserID | Valid", func(t *testing.T) {
		wishesRepository.On("FindByUserID", mock.AnythingOfType("int")).Return([]wishes.Domain{wishesDomain}, nil).Once()

		result, err := wishesService.FindByUserID(wishesDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("Find By UserID | InValid", func(t *testing.T) {
		wishesRepository.On("FindByUserID", mock.AnythingOfType("int")).Return([]wishes.Domain{wishesDomain}, businesses.ErrCategoryNotFound).Once()

		_, err := wishesService.FindByUserID(wishesDomain.ID)

		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid 1", func(t *testing.T) {
		productsRepository.On("GetProduct", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(1, "1", false, "1").Once()
		wishesRepository.On("Update", mock.AnythingOfType("*wishes.Domain"), mock.AnythingOfType("int")).Return(&wishesDomain, nil).Maybe().Twice()
		
		result, err := wishesService.Update(&wishesDomain, wishesDomain.ID)
		
		assert.Nil(t, err)
		assert.Equal(t, &wishesDomain, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		productsRepository.On("GetProduct", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(1, "1", false, "1").Once()
		wishesRepository.On("Update", mock.AnythingOfType("*wishes.Domain"), mock.AnythingOfType("int")).Return(&wishesDomain, businesses.ErrInternalServer).Twice()

		_, err := wishesService.Update(&wishesDomain, wishesDomain.ID)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		wishesRepository.On("Delete", mock.AnythingOfType("*wishes.Domain"), mock.AnythingOfType("int")).Return("Data Deleted.", nil).Once()

		result, err := wishesService.Delete(&wishesDomain, wishesDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, "Data Deleted.", result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		wishesRepository.On("Delete", mock.AnythingOfType("*wishes.Domain"), mock.AnythingOfType("int")).Return("Fail to delete.", businesses.ErrInternalServer).Once()

		_, err := wishesService.Delete(&wishesDomain, wishesDomain.ID)

		assert.NotNil(t, err)
	})
}