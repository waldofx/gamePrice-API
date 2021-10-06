package wishes

import (
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
	User      		users.Users
	GameID			int
	Game			games.Games
	SellerID		int	
	Seller			sellers.Sellers
	ProductID		int
	Price2			string
	Discount		bool
	URL				string
}

func ToDomain(rec Wishes) wishes.Domain {
	return wishes.Domain{
		ID:        	int(rec.ID),
		UserID: 	rec.UserID,
		Username: 	rec.User.Username,
		GameID: 	rec.GameID,
		GameName: 	rec.Game.Name,
		SellerID: 	rec.SellerID,
		GameSeller: rec.Seller.Name,
		ProductID: 	rec.ProductID,
		Price: 		rec.Price2,
		Discount: 	rec.Discount,
		URL: 		rec.URL,	
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
		ProductID:  domain.ProductID,
		Price2: 	domain.Price,
		Discount: 	domain.Discount,
		URL: 		domain.URL,
	}
}

func ToDomainArray(modelWishes []Wishes) []wishes.Domain{
	var response []wishes.Domain

	for _, val := range modelWishes{
		response = append(response, ToDomain(val))
	}
	return response
}