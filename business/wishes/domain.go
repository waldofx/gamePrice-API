package wishes

import (
	"time"
)

// domain layer / entity layer -> acuan utama dalam domain.
type Domain struct {
	ID        		int
	UserID	  		int
	Username      	string
	GameID	  		int
	SellerID		int
	GameName      	string //products.Domain
	GameSeller		string
	Price			int
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}

// inteface of bussiness layer -> fungsi fungsi yang di butuhkan oleh domain (bussiness logic)
type Service interface {
	Append(wish *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(wish *Domain, id int) (*Domain, error)
	Delete(wish *Domain, id int) (string, error)
}

// interface of data layer -> fungsi fungsi untuk membaca / memerintah database, 3rd Party, ataupun package lain.
type Repository interface {
	Insert(wish *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(wish *Domain, id int) (*Domain, error)
	Delete(wish *Domain, id int) (string, error)
}
