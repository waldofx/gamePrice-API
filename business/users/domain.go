package users

import (
	"context"
	"time"
)

// domain layer / entity layer -> acuan utama dalam domain.
type Domain struct {
	ID        	int
	Username	string
	Email  	  	string
	Password	string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

// inteface of bussiness layer -> fungsi fungsi yang di butuhkan oleh domain (bussiness logic)
type Service interface {
	CreateToken(ctx context.Context, username, password string) (string, error)
	Store(ctx context.Context, data *Domain) error
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(user *Domain, id int) (*Domain, error)
	Delete(user *Domain, id int) (string, error)
}

// interface of data layer -> fungsi fungsi untuk membaca / memerintah database, 3rd Party, ataupun package lain.
type Repository interface {
	Store(ctx context.Context, data *Domain) error
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	FindByUsername(ctx context.Context, username string) (Domain, error)
	Update(user *Domain, id int) (*Domain, error)
	Delete(user *Domain, id int) (string, error)
}
