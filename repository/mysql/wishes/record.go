package wishes

import (
	"fmt"
	"gameprice-api/business/wishes"
	"gameprice-api/repository/mysql/products"
	"gameprice-api/repository/mysql/users"

	"gorm.io/gorm"
)

type Wishes struct {
	gorm.Model
	ID       		uint 				`gorm:"primaryKey"`
	UserID	 		int					//`gorm:"foreignKey:users_id"`
	User      		users.Users 		//`gorm:"foreignKey:users_id"`
	ProductID  		int					//`gorm:"foreignKey:products_id"`
	Product    		products.Products 	//`gorm:"foreignKey:products_id"`
	Price	  		int
}

func ToDomain(rec Wishes) wishes.Domain {
	fmt.Printf("%+v", rec)
	return wishes.Domain{
		ID:        	int(rec.ID),
		UserID: 	rec.UserID,
		User: 		rec.User.Name,
		ProductID: 	rec.ProductID,
		Product: 	rec.Product.Game.Name,
		CreatedAt: 	rec.CreatedAt,
		UpdatedAt: 	rec.UpdatedAt,
	}
}

func FromDomain(domain wishes.Domain) Wishes {
	return Wishes{
		ID: 		uint(domain.ID),
		UserID: 	domain.UserID,
		ProductID: 	domain.ProductID,
	}
}

func ToDomainArray(modelWishes []Wishes) []wishes.Domain{
	var response []wishes.Domain

	for _, val := range modelWishes{
		response = append(response, ToDomain(val))
	}
	return response
}