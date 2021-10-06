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

		result, err := productsService.APIDetail(&productsDomain)

		assert.Nil(t, err)
		assert.Equal(t, &steamapisDomain.Price, &result.Price)
		assert.Equal(t, &steamapisDomain.Discount, &result.Discount)
		assert.Equal(t, "https://store.steampowered.com/app/1", result.URL)
	})
	t.Run("API Detail | Case 2 Valid", func(t *testing.T)  {
		steamapisRepository.On("GetID", mock.AnythingOfType("string")).Return(steamapisDomain, nil).Once()
		steamapisRepository.On("GetData", mock.AnythingOfType("string")).Return(steamapisDomain, nil).Once()

		productsDomain2 := products.Domain{
			SellerID:  	2,
		}

		result, err := productsService.APIDetail(&productsDomain2)

		assert.Nil(t, err)
		assert.Equal(t, "Price is not available from this seller.", result.Price)
		assert.Equal(t, "https://www.gog.com/game/", result.URL)
	})
	t.Run("API Detail | Case 1 inValid", func(t *testing.T)  {
		steamapisRepository.On("GetID", mock.AnythingOfType("string")).Return(steamapisDomain, businesses.ErrInternalServer).Once()
		steamapisRepository.On("GetData", mock.AnythingOfType("string")).Return(steamapisDomain, businesses.ErrInternalServer).Once()

		_, err := productsService.APIDetail(&products.Domain{
			GameID: 20,
			Game: "aaaaaaaaaaaaaaa",
		})
	
		assert.Nil(t, err)
	})
}

func TestAppend(t *testing.T){
	t.Run("Append | Valid", func(t *testing.T) {
		productsRepository.On("Insert", mock.AnythingOfType("*products.Domain")).Return(&productsDomain, nil).Once()
		steamapisRepository.On("GetID", mock.AnythingOfType("string")).Return(steamapisDomain, nil).Once()
		steamapisRepository.On("GetData", mock.AnythingOfType("string")).Return(steamapisDomain, nil).Once()
		productsRepository.On("Update", mock.AnythingOfType("*products.Domain"), mock.AnythingOfType("int")).Return(&productsDomain, nil).Once()

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
	// t.Run("Update | Valid 1", func(t *testing.T) {
	// 	productsRepository.On("Update", mock.AnythingOfType("*products.Domain"), 2).Return(&productsDomain, nil).Once()
	// 	steamapisRepository.On("GetID", mock.AnythingOfType("string")).Return(steamapisDomain, nil).Once()
	// 	steamapisRepository.On("GetData", mock.AnythingOfType("string")).Return(steamapisDomain, nil).Once()
	// 	productsRepository.On("Update", mock.AnythingOfType("*products.Domain"), 2).Return(&productsDomain, nil).Once()
		
	// 	result, err := productsService.Update(&productsDomain, 2)
	// 	want := &products.Domain{ID:0, GameID:0, Game:"", SellerID:0, Seller:"", Price:"", Discount:false, URL:"", CreatedAt:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), UpdatedAt:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}

	// 	assert.Nil(t, err)
	// 	assert.Equal(t, want, result)
	// })

	t.Run("Update | InValid", func(t *testing.T) {
		productsRepository.On("Update", mock.AnythingOfType("*products.Domain"), mock.AnythingOfType("int")).Return(nil, businesses.ErrNotFound).Twice()
		steamapisRepository.On("GetID", mock.AnythingOfType("string")).Return(nil, businesses.ErrNotFound).Once()
		steamapisRepository.On("GetData", mock.AnythingOfType("string")).Return(nil, businesses.ErrNotFound).Once()

		_, err := productsService.Update(&productsDomain, productsDomain.ID)
		assert.NotNil(t, err)
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