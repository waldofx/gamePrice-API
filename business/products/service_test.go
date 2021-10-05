package products_test

import (
	businesses "gameprice-api/business"
	"gameprice-api/business/products"
	_productsMock "gameprice-api/business/products/mocks"
	"gameprice-api/business/steamapis"
	_steamapisMock "gameprice-api/business/steamapis/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	productsRepository _productsMock.Repository
	steamapisRepository	_steamapisMock.Repository
	productsService	products.Service
	productsDomain 	products.Domain
	steamapisDomain steamapis.Domain
)

func TestMain(m *testing.M) {
	productsService = products.NewService(&productsRepository, &steamapisRepository)
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
	steamapisDomain = steamapis.Domain{
		AppID: "1",
		Name: "Game Test",
		Price: "Rp 0",
		Discount: false,
	}
	os.Exit(m.Run())
}

func TestAPIDetail(t *testing.T){
	t.Run("API Detail | Case 1 Valid", func(t *testing.T)  {
		steamapisRepository.On("GetID", mock.AnythingOfType("string")).Return(steamapisDomain, nil).Once()
		steamapisRepository.On("GetData", mock.AnythingOfType("string")).Return(steamapisDomain, nil).Once()

		steam, err := steamapisRepository.GetID(productsDomain.Game)

		assert.Nil(t, err)

		steam2, err := steamapisRepository.GetData(steam.AppID)

		assert.Nil(t, err)
		assert.Equal(t, &steamapisDomain.AppID, &steam2.AppID)
		assert.Equal(t, &steamapisDomain.Price, &steam2.Price)
		assert.Equal(t, &steamapisDomain.Discount, &steam2.Discount)

		if productsDomain.SellerID == 2{
			productsDomain.Price = "Price is not available from this seller."
			productsDomain.URL = "https://www.gog.com/game/game"
			assert.Equal(t, "Price is not available from this seller.", productsDomain.Price)
			assert.Equal(t, "https://www.gog.com/game/game", productsDomain.URL)
		}
	})
	t.Run("API Detail | Case 1 inValid", func(t *testing.T)  {
		steamapisRepository.On("GetID", mock.AnythingOfType("string")).Return(steamapisDomain, businesses.ErrInternalServer).Once()

		_, err := steamapisRepository.GetID(productsDomain.Game)
	
		assert.NotNil(t, err)
	})

	t.Run("API Detail | Case 2 InValid", func(t *testing.T)  {
		steam := &productsDomain
		assert.NotNil(t, &productsDomain, steam)
	})
}

func TestAppend(t *testing.T){
	t.Run("Append | Valid", func(t *testing.T) {
		productsRepository.On("Insert", mock.AnythingOfType("*products.Domain")).Return(&productsDomain, nil).Once()
		steamapisRepository.On("GetID", mock.AnythingOfType("string")).Return(steamapisDomain, nil).Maybe().Twice()
		steamapisRepository.On("GetData", mock.AnythingOfType("string")).Return(steamapisDomain, nil).Maybe().Twice().Twice()
		productsRepository.On("Update", mock.AnythingOfType("*products.Domain"), mock.AnythingOfType("int")).Return(&productsDomain, nil).Maybe().Twice()

		result, err := productsService.Append(&productsDomain)

		assert.Nil(t, err)
		assert.Equal(t, &productsDomain, result)
	})

	t.Run("Append | InValid", func(t *testing.T) {
		productsRepository.On("Insert", mock.AnythingOfType("*products.Domain")).Return(&productsDomain, businesses.ErrInternalServer).Once()

		_, err := productsService.Append(&productsDomain)

		assert.NotNil(t, err)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("Find All | Valid", func(t *testing.T) {
		productsRepository.On("FindAll").Return([]products.Domain{productsDomain}, nil).Once()

		result, err := productsService.FindAll()

		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("Find All | InValid", func(t *testing.T) {
		productsRepository.On("FindAll").Return([]products.Domain{}, businesses.ErrCategoryNotFound).Once()

		_, err := productsService.FindAll()

		assert.NotNil(t, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("Find By ID | Valid", func(t *testing.T) {
		productsRepository.On("FindByID", mock.AnythingOfType("int")).Return(&productsDomain, nil).Once()

		result, err := productsService.FindByID(productsDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, &productsDomain, result)
	})

	t.Run("Find By ID | InValid", func(t *testing.T) {
		productsRepository.On("FindByID", mock.AnythingOfType("int")).Return(&productsDomain, businesses.ErrCategoryNotFound).Once()

		_, err := productsService.FindByID(productsDomain.ID)

		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid 1", func(t *testing.T) {
		productsRepository.On("Update", mock.AnythingOfType("*products.Domain"), mock.AnythingOfType("int")).Return(&productsDomain, nil).Maybe().Twice()
		steamapisRepository.On("GetID", mock.AnythingOfType("string")).Return(steamapisDomain, nil).Maybe().Twice()
		steamapisRepository.On("GetData", mock.AnythingOfType("string")).Return(steamapisDomain, nil).Maybe().Twice().Twice()
		
		result, err := productsService.Update(&productsDomain, productsDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, &productsDomain, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		productsRepository.On("Update", mock.AnythingOfType("*products.Domain"), mock.AnythingOfType("int")).Return(nil, businesses.ErrNotFound).Twice()
		steamapisRepository.On("GetID", mock.AnythingOfType("string")).Return(nil, businesses.ErrNotFound).Maybe().Twice()
		steamapisRepository.On("GetData", mock.AnythingOfType("string")).Return(nil, businesses.ErrNotFound).Maybe().Twice()
		//productsRepository.On("Update", mock.AnythingOfType("*products.Domain"), mock.AnythingOfType("int")).Return(nil, businesses.ErrInternalServer).Twice()

		_, err := productsService.Update(&productsDomain, productsDomain.ID)
		assert.NotNil(t, err)

		// _, err = productsService.APIDetail(&productsDomain)
		// assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		productsRepository.On("Delete", mock.AnythingOfType("*products.Domain"), mock.AnythingOfType("int")).Return("Data Deleted.", nil).Once()

		result, err := productsService.Delete(&productsDomain, productsDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, "Data Deleted.", result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		productsRepository.On("Delete", mock.AnythingOfType("*products.Domain"), mock.AnythingOfType("int")).Return("Fail to delete.", businesses.ErrInternalServer).Once()

		_, err := productsService.Delete(&productsDomain, productsDomain.ID)

		assert.NotNil(t, err)
	})
}