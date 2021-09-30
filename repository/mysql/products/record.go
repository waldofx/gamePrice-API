package products

import (
	"fmt"
	"gameprice-api/business/products"
	"gameprice-api/repository/mysql/games"
	"gameprice-api/repository/mysql/sellers"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	ID       		uint 			`gorm:"primaryKey"`
	GameID	 		int				//`gorm:"foreignKey:games_id"`
	Game      		games.Games 	//`gorm:"foreignKey:games_id"`
	SellerID  		int				//`gorm:"foreignKey:sellers_id"`
	Seller    		sellers.Sellers //`gorm:"foreignKey:sellers_id"`
	Price	  		int
}

func ToDomain(rec Products) products.Domain {
	fmt.Printf("%+v", rec)
	return products.Domain{
		ID:        	int(rec.ID),
		GameID: 	rec.GameID,
		Game: 		rec.Game.Name,
		SellerID: 	rec.SellerID,
		Seller: 	rec.Seller.Name,
		Price: 		rec.Price,
		CreatedAt: 	rec.CreatedAt,
		UpdatedAt: 	rec.UpdatedAt,
	}
}

func FromDomain(domain products.Domain) Products {
	return Products{
		ID: 		uint(domain.ID),
		GameID: 	domain.GameID,
		SellerID: 	domain.SellerID,
		Price: domain.Price,
	}
}

func ToDomainArray(modelProducts []Products) []products.Domain{
	var response []products.Domain

	for _, val := range modelProducts{
		response = append(response, ToDomain(val))
	}
	return response
}