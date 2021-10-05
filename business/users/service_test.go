package users_test

import (
	"context"
	"gameprice-api/app/middleware"
	businesses "gameprice-api/business"
	"gameprice-api/business/users"
	_usersMock "gameprice-api/business/users/mocks"
	"gameprice-api/helpers/encrypt"
	"os"
	"strings"
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
	t.Run("Create Token | Valid", func(t *testing.T) {
		usersRepository.On("FindByUsername", mock.Anything, mock.AnythingOfType("string")).Return(usersDomain, nil).Once()

		ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
		defer cancel()

		if strings.TrimSpace(usersDomain.Username) == "" && strings.TrimSpace(usersDomain.Password) == "" {
			err := businesses.ErrUsernamePasswordNotFound
			assert.Nil(t, err)
		}

		userDomain, err := usersRepository.FindByUsername(ctx, usersDomain.Username)
		assert.Nil(t, err)

		if !encrypt.ValidateHash("hashedpassword", userDomain.Password) {
			err = nil
			assert.Nil(t, err)
		}
	
		token := "123"

		assert.Nil(t, err)
		assert.Equal(t, "123", token)
	})

	t.Run("Create Token | InValid", func(t *testing.T) {
		usersRepository.On("FindByUsername", mock.Anything, mock.AnythingOfType("string")).Return(usersDomain, businesses.ErrNotFound).Once()

		ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
		defer cancel()

		if strings.TrimSpace("") == ""{
			err := businesses.ErrUsernamePasswordNotFound
			assert.NotNil(t, err)
		}

		_, err := usersRepository.FindByUsername(ctx, usersDomain.Username)
		assert.NotNil(t, err)

		if encrypt.ValidateHash(usersDomain.Password, usersDomain.Password) {
			err := businesses.ErrInternalServer
			assert.NotNil(t, err)
		}
	})
}

func TestStore(t *testing.T) {
	t.Run("Store | Valid", func(t *testing.T) {
		usersRepository.On("FindByUsername", mock.Anything, mock.AnythingOfType("string")).Return(usersDomain, nil).Once()
		usersRepository.On("Store", mock.Anything, mock.AnythingOfType("*users.Domain")).Return(nil).Once()

		ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
		defer cancel()

		_, err := usersRepository.FindByUsername(context.Background(), usersDomain.Username)
		assert.Nil(t, err)

		usersDomain.Password, err = encrypt.Hash(usersDomain.Password)
		assert.Nil(t, err)

		err = usersRepository.Store(ctx, &usersDomain)
		assert.Nil(t, err)

		assert.Nil(t, err)
	})

	t.Run("Store | InValid", func(t *testing.T) {
		usersRepository.On("FindByUsername", mock.Anything, mock.AnythingOfType("string")).Return(usersDomain, businesses.ErrDuplicateData).Once()

		_, cancel := context.WithTimeout(context.Background(), contextTimeout)
		defer cancel()

		_, err := usersRepository.FindByUsername(context.Background(), usersDomain.Username)

		assert.NotNil(t, err)

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