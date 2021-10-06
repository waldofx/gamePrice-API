package sellers_test

import (
	businesses "gameprice-api/business"
	"gameprice-api/business/sellers"
	_sellersMock "gameprice-api/business/sellers/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	sellersRepository _sellersMock.Repository
	sellersService	sellers.Service
	sellersDomain sellers.Domain
)

func TestMain(m *testing.M) {
	sellersService = sellers.NewService(&sellersRepository)
	sellersDomain = sellers.Domain{
		ID:          1,
		Name:       "Test Sellers",
		URL:		"example.com",
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
	}
	os.Exit(m.Run())
}

func TestAppend(t *testing.T){
	t.Run("Append | Valid", func(t *testing.T) {
		sellersRepository.On("Insert", mock.AnythingOfType("*sellers.Domain")).Return(&sellersDomain, nil).Once()

		result, err := sellersService.Append(&sellersDomain)

		assert.Nil(t, err)
		assert.Equal(t, &sellersDomain, result)
	})

	t.Run("Append | InValid", func(t *testing.T) {
		sellersRepository.On("Insert", mock.AnythingOfType("*sellers.Domain")).Return(&sellersDomain, businesses.ErrInternalServer).Once()

		_, err := sellersService.Append(&sellersDomain)

		assert.NotNil(t, err)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("Find All | Valid", func(t *testing.T) {
		sellersRepository.On("FindAll").Return([]sellers.Domain{sellersDomain}, nil).Once()

		result, err := sellersService.FindAll()

		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("Find All | InValid", func(t *testing.T) {
		sellersRepository.On("FindAll").Return([]sellers.Domain{}, businesses.ErrCategoryNotFound).Once()

		_, err := sellersService.FindAll()

		assert.NotNil(t, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("Find By ID | Valid", func(t *testing.T) {
		sellersRepository.On("FindByID", mock.AnythingOfType("int")).Return(&sellersDomain, nil).Once()

		result, err := sellersService.FindByID(sellersDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, &sellersDomain, result)
	})

	t.Run("Find By ID | InValid", func(t *testing.T) {
		sellersRepository.On("FindByID", mock.AnythingOfType("int")).Return(&sellersDomain, businesses.ErrCategoryNotFound).Once()

		_, err := sellersService.FindByID(sellersDomain.ID)

		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		sellersRepository.On("Update", mock.AnythingOfType("*sellers.Domain"), mock.AnythingOfType("int")).Return(&sellersDomain, nil).Once()

		result, err := sellersService.Update(&sellersDomain, sellersDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, &sellersDomain, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		sellersRepository.On("Update", mock.AnythingOfType("*sellers.Domain"), mock.AnythingOfType("int")).Return(&sellersDomain, businesses.ErrInternalServer).Once()

		_, err := sellersService.Update(&sellersDomain, sellersDomain.ID)

		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		sellersRepository.On("Delete", mock.AnythingOfType("*sellers.Domain"), mock.AnythingOfType("int")).Return("Data Deleted.", nil).Once()

		result, err := sellersService.Delete(&sellersDomain, sellersDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, "Data Deleted.", result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		sellersRepository.On("Delete", mock.AnythingOfType("*sellers.Domain"), mock.AnythingOfType("int")).Return("Fail to delete.", businesses.ErrInternalServer).Once()

		_, err := sellersService.Delete(&sellersDomain, sellersDomain.ID)

		assert.NotNil(t, err)
	})
}