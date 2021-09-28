package games

import "time"

// domain layer / entity layer -> acuan utama dalam domain.
type Domain struct {
	ID        int
	Name      string
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// inteface of bussiness layer -> fungsi fungsi yang di butuhkan oleh domain (bussiness logic)
type Service interface {
	Append(game *Domain) (*Domain, error)
	Update(game *Domain, id int) (*Domain, error)
	FindByID(id int) (*Domain, error)
	Available(generalSearch string) []Domain
}

// interface of data layer -> fungsi fungsi untuk membaca / memerintah database, 3rd Party, ataupun package lain.
type Repository interface {
	Insert(game *Domain) (*Domain, error)
	Update(game *Domain, id int) (*Domain, error)
	FindByID(id int) (*Domain, error)
	FindAll(generalSearch string, availability bool) []Domain
}