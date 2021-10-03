package wishes

import (
	"fmt"
	"gameprice-api/business/wishes"
	"gameprice-api/repository/mysql/games"
	"gameprice-api/repository/mysql/sellers"
	"gameprice-api/repository/mysql/users"

	"gorm.io/gorm"
)

type Wishes struct {
	gorm.Model
	ID       		uint 				`gorm:"primaryKey"`
	UserID	 		int					
	User      		users.Users 		//`gorm:"foreignKey:users_id"`
	GameID			int
	Game			games.Games
	SellerID		int	
	Seller			sellers.Sellers
	Price	  		int
}

func ToDomain(rec Wishes) wishes.Domain {
	fmt.Printf("%+v", rec)
	return wishes.Domain{
		ID:        	int(rec.ID),
		UserID: 	rec.UserID,
		Username: 	rec.User.Username,
		GameID: 	rec.GameID,
		GameName: 	rec.Game.Name,
		SellerID: 	rec.SellerID,
		GameSeller: rec.Seller.Name,
		CreatedAt: 	rec.CreatedAt,
		UpdatedAt: 	rec.UpdatedAt,
	}
}

func FromDomain(domain wishes.Domain) Wishes {
	return Wishes{
		ID: 		uint(domain.ID),
		UserID: 	domain.UserID,
		GameID: 	domain.GameID,
		SellerID: 	domain.SellerID,
	}
}

func ToDomainArray(modelWishes []Wishes) []wishes.Domain{
	var response []wishes.Domain

	for _, val := range modelWishes{
		response = append(response, ToDomain(val))
	}
	return response
}