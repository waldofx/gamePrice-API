package sellers

import "time"

// domain layer / entity layer -> acuan utama dalam domain.
type Domain struct {
	ID        int
	Name      string
	URL  	  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// inteface of bussiness layer -> fungsi fungsi yang di butuhkan oleh domain (bussiness logic)
type Service interface {
	Append(seller *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(seller *Domain, id int) (*Domain, error)
	Delete(seller *Domain, id int) (string, error)
}

// interface of data layer -> fungsi fungsi untuk membaca / memerintah database, 3rd Party, ataupun package lain.
type Repository interface {
	Insert(seller *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(seller *Domain, id int) (*Domain, error)
	Delete(seller *Domain, id int) (string, error)
}
