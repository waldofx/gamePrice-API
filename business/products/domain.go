package products

import "time"

// domain layer / entity layer -> acuan utama dalam domain.
type Domain struct {
	ID        int
	GameID	  int
	Game      string
	SellerID  int
	Seller    string
	Price	  string
	Discount  bool
	URL		  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// inteface of bussiness layer -> fungsi fungsi yang di butuhkan oleh domain (bussiness logic)
type Service interface {
	Append(product *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(product *Domain, id int) (*Domain, error)
	Delete(product *Domain, id int) (string, error)
}

// interface of data layer -> fungsi fungsi untuk membaca / memerintah database, 3rd Party, ataupun package lain.
type Repository interface {
	Insert(product *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(product *Domain, id int) (*Domain, error)
	Delete(product *Domain, id int) (string, error)
	GetProduct(int, int) (int, string, bool, string)
}
