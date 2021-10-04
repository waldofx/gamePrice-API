package users

import (
	"context"
	"gameprice-api/app/middleware"
	businesses "gameprice-api/business"
	"gameprice-api/helpers/encrypt"
	"strings"
	"time"
)

type serviceUsers struct {
	repository     Repository
	jwtAuth        *middleware.ConfigJWT
	contextTimeout time.Duration
}

func NewService(repoUser Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Service {
	return &serviceUsers{
		repository: repoUser,
		jwtAuth:        jwtauth,
		contextTimeout: timeout,
	}
}

func (servUser *serviceUsers) CreateToken(ctx context.Context, username, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, servUser.contextTimeout)
	defer cancel()

	if strings.TrimSpace(username) == "" && strings.TrimSpace(password) == "" {
		return "", businesses.ErrUsernamePasswordNotFound
	}

	userDomain, err := servUser.repository.FindByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, userDomain.Password) {
		return "", businesses.ErrInternalServer
	}

	token := servUser.jwtAuth.GenerateToken(userDomain.ID)
	return token, nil
}

func (servUser *serviceUsers) Store(ctx context.Context, userDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, servUser.contextTimeout)
	defer cancel()

	existedUser, err := servUser.repository.FindByUsername(ctx, userDomain.Username)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}
	if existedUser != (Domain{}) {
		return businesses.ErrDuplicateData
	}

	userDomain.Password, err = encrypt.Hash(userDomain.Password)
	if err != nil {
		return businesses.ErrInternalServer
	}
	err = servUser.repository.Store(ctx, userDomain)
	if err != nil {
		return err
	}

	return nil
}

func (servUser *serviceUsers) FindAll() ([]Domain, error) {
	result, err := servUser.repository.FindAll()
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (servUser *serviceUsers) FindByID(id int) (*Domain, error) {
	result, err := servUser.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servUser *serviceUsers) Update(user *Domain, id int) (*Domain, error) {
	result, err := servUser.repository.Update(user, id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servUser *serviceUsers) Delete(user *Domain, id int) (string, error) {
	result, err := servUser.repository.Delete(user, id)
	if err != nil {
		return "Fail to delete.", err
	}
	return result, nil
}
