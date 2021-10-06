package users_test

import (
	"context"
	"gameprice-api/app/middleware"
	businesses "gameprice-api/business"
	"gameprice-api/business/users"
	_usersMock "gameprice-api/business/users/mocks"
	"gameprice-api/helpers/encrypt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	usersRepository _usersMock.Repository
	usersService	users.Service
	usersDomain 	users.Domain
	jwtAuth        	*middleware.ConfigJWT
	contextTimeout 	time.Duration
)

func TestMain(m *testing.M) {
	usersService = users.NewService(&usersRepository, jwtAuth, contextTimeout)
	usersDomain = users.Domain{
		ID:          1,
		Username:	"Test Users",
		Email: 		"test@email.com",
		Password:	"hashedpassword",
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
	}
	os.Exit(m.Run())
}

func TestCreateToken(t *testing.T) {
	t.Run("Create Token | InValid 1", func(t *testing.T) {
		usersRepository.On("FindByUsername", mock.Anything, mock.AnythingOfType("string")).Return(usersDomain, nil).Once()

		password2, _ := encrypt.Hash(usersDomain.Password)
		_, err := usersService.CreateToken(context.Background(), usersDomain.Username, password2)

		assert.NotNil(t, err)
	})

	t.Run("Create Token | InValid 2", func(t *testing.T) {
		usersRepository.On("FindByUsername", mock.Anything, mock.AnythingOfType("string")).Return(usersDomain, businesses.ErrNotFound).Once()

		_, err := usersService.CreateToken(context.Background(), usersDomain.Username, usersDomain.Password)

		assert.NotNil(t, err)
	})
	// t.Run("Create Token | InValid 3", func(t *testing.T) {
	// 	//usersRepository.On("FindByUsername", mock.Anything, mock.AnythingOfType("string")).Return(usersDomain, nil).Once()

	// 	_, err := usersService.CreateToken(context.Background(),"", "")

	// 	assert.NotNil(t, err)
	// })
}

func TestStore(t *testing.T) {
	t.Run("Store | Valid", func(t *testing.T) {
		usersRepository.On("FindByUsername", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, nil).Once()
		usersRepository.On("Store", mock.Anything, mock.AnythingOfType("*users.Domain")).Return(nil).Once()

		usersDomain2 := users.Domain{
			ID: 		2,
			Username:	"JOKO",
			Email: 		"test2@email.com",
			Password:	"hashedpassword2",
			CreatedAt: 	time.Now(),
			UpdatedAt: 	time.Now(),
		}

		err := usersService.Store(context.Background(), &usersDomain2)
		assert.Nil(t, err)

	})
	// t.Run("Store | Valid", func(t *testing.T) {
	// 	usersRepository.On("FindByUsername", mock.Anything, mock.AnythingOfType("string")).Return(usersDomain, nil).Once()
	// 	usersRepository.On("Store", mock.Anything, mock.AnythingOfType("*users.Domain")).Return(nil).Once()

	// 	err := usersService.Store(context.Background(), &usersDomain)
	// 	assert.NotNil(t, err)

	// })

	t.Run("Store | InValid", func(t *testing.T) {
		usersRepository.On("FindByUsername", mock.Anything, mock.AnythingOfType("string")).Return(usersDomain, businesses.ErrDuplicateData).Once()
		usersRepository.On("Store", mock.Anything, mock.AnythingOfType("*users.Domain")).Return(businesses.ErrDuplicateData).Once()

		err := usersService.Store(context.Background(), &usersDomain)
		assert.NotNil(t, err)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("Find All | Valid", func(t *testing.T) {
		usersRepository.On("FindAll").Return([]users.Domain{usersDomain}, nil).Once()

		result, err := usersService.FindAll()

		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("Find All | InValid", func(t *testing.T) {
		usersRepository.On("FindAll").Return([]users.Domain{}, businesses.ErrCategoryNotFound).Once()

		_, err := usersService.FindAll()

		assert.NotNil(t, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("Find By ID | Valid", func(t *testing.T) {
		usersRepository.On("FindByID", mock.AnythingOfType("int")).Return(&usersDomain, nil).Once()

		result, err := usersService.FindByID(usersDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, &usersDomain, result)
	})

	t.Run("Find By ID | InValid", func(t *testing.T) {
		usersRepository.On("FindByID", mock.AnythingOfType("int")).Return(&usersDomain, businesses.ErrCategoryNotFound).Once()

		_, err := usersService.FindByID(usersDomain.ID)

		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		usersRepository.On("Update", mock.AnythingOfType("*users.Domain"), mock.AnythingOfType("int")).Return(&usersDomain, nil).Once()

		usersDomain.Password, _ = encrypt.Hash(usersDomain.Password)
		result, err := usersService.Update(&usersDomain, usersDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, &usersDomain, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		usersRepository.On("Update", mock.AnythingOfType("*users.Domain"), mock.AnythingOfType("int")).Return(&usersDomain, businesses.ErrInternalServer).Once()

		usersDomain.Password, _ = encrypt.Hash(usersDomain.Password)
		_, err := usersService.Update(&usersDomain, usersDomain.ID)

		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		usersRepository.On("Delete", mock.AnythingOfType("*users.Domain"), mock.AnythingOfType("int")).Return("Data Deleted.", nil).Once()

		result, err := usersService.Delete(&usersDomain, usersDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, "Data Deleted.", result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		usersRepository.On("Delete", mock.AnythingOfType("*users.Domain"), mock.AnythingOfType("int")).Return("Fail to delete.", businesses.ErrInternalServer).Once()

		_, err := usersService.Delete(&usersDomain, usersDomain.ID)

		assert.NotNil(t, err)
	})
}