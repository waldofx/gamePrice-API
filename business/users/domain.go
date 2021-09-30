package users

import "time"

// domain layer / entity layer -> acuan utama dalam domain.
type Domain struct {
	ID        	int
	Name   		string
	Email  	  	string
	Password	string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

// inteface of bussiness layer -> fungsi fungsi yang di butuhkan oleh domain (bussiness logic)
type Service interface {
	Append(user *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(user *Domain, id int) (*Domain, error)
	Delete(user *Domain, id int) (string, error)
}

// interface of data layer -> fungsi fungsi untuk membaca / memerintah database, 3rd Party, ataupun package lain.
type Repository interface {
	Insert(user *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(user *Domain, id int) (*Domain, error)
	Delete(user *Domain, id int) (string, error)
}
