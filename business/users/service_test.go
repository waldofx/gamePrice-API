package users_test

import (
	"gameprice-api/app/middleware"
	businesses "gameprice-api/business"
	"gameprice-api/business/users"
	_usersMock "gameprice-api/business/users/mocks"
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
		Password:	"password",
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
	}
	os.Exit(m.Run())
}


// func TestCreateToken(t *testing.T) {
// 	t.Run("Create Token | Valid", func(t *testing.T) {
// 		usersRepository.On("FindAll").Return([]users.Domain{usersDomain}, nil).Once()

// 		result, err := usersService.FindAll()

// 		assert.Nil(t, err)
// 		assert.Equal(t, 1, len(result))
// 	})

// 	t.Run("Create Token | InValid", func(t *testing.T) {
// 		usersRepository.On("FindAll").Return([]users.Domain{}, businesses.ErrCategoryNotFound).Once()

// 		_, err := usersService.FindAll()

// 		assert.NotNil(t, err)
// 	})

// 	ctx, cancel := context.WithTimeout(ctx, servUser.contextTimeout)
// 	defer cancel()

// 	if strings.TrimSpace(username) == "" && strings.TrimSpace(password) == "" {
// 		return "", businesses.ErrUsernamePasswordNotFound
// 	}

// 	userDomain, err := servUser.repository.FindByUsername(ctx, username)
// 	if err != nil {
// 		return "", err
// 	}

// 	if !encrypt.ValidateHash(password, userDomain.Password) {
// 		return "", businesses.ErrInternalServer
// 	}

// 	token := servUser.jwtAuth.GenerateToken(userDomain.ID)
// 	return token, nil
// }

// func TestStore(t *testing.T) {
// 	t.Run("Store | Valid", func(t *testing.T) {
// 		usersRepository.On("FindByUsername", mock.AnythingOfType("int")).Return([]users.Domain{usersDomain}, nil).Once()

// 		result, err := usersService.FindByUsername

// 		assert.Nil(t, err)
// 		assert.Equal(t, 1, len(result))
// 	})

// 	t.Run("Store | InValid", func(t *testing.T) {
// 		usersRepository.On("FindByUsername").Return([]users.Domain{}, businesses.ErrCategoryNotFound).Once()

// 		_, err := usersService.FindAll()

// 		assert.NotNil(t, err)
// 	})
	// ctx, cancel := context.WithTimeout(ctx, servUser.contextTimeout)
	// defer cancel()

	// existedUser, err := servUser.repository.FindByUsername(ctx, userDomain.Username)
	// if err != nil {
	// 	if !strings.Contains(err.Error(), "not found") {
	// 		return err
	// 	}
	// }
	// if existedUser != (Domain{}) {
	// 	return businesses.ErrDuplicateData
	// }

	// userDomain.Password, err = encrypt.Hash(userDomain.Password)
	// if err != nil {
	// 	return businesses.ErrInternalServer
	// }
	// err = servUser.repository.Store(ctx, userDomain)
	// if err != nil {
	// 	return err
	// }

	// return nil
// }

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

		result, err := usersService.Update(&usersDomain, usersDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, &usersDomain, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		usersRepository.On("Update", mock.AnythingOfType("*users.Domain"), mock.AnythingOfType("int")).Return(&usersDomain, businesses.ErrInternalServer).Once()

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