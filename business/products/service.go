package products

import (
	"fmt"
	"gameprice-api/business/gogapis"
	"gameprice-api/business/steamapis"
)

type serviceProducts struct {
	repository 	Repository
	reposteam 	steamapis.Repository
	repogog		gogapis.Repository
}

func NewService(repoProduct Repository, rs steamapis.Repository, rg gogapis.Repository) Service {
	return &serviceProducts{
		repository: repoProduct,
		reposteam: rs,
		repogog: rg,
	}
}

func (servProduct *serviceProducts) Append(product *Domain) (*Domain, error) {
	result, err := servProduct.repository.Insert(product)
	if err != nil {
		return &Domain{}, err
	}
	fmt.Println(" finish init Insert")

	if product.SellerID == 1 {
		steam, err := servProduct.reposteam.GetID(result.Game)
		if err != nil {
			return &Domain{}, err
		}
		fmt.Println(" finish GetID, result.Game: " + result.Game)
		fmt.Println(" finish GetID, steam.Name: " + steam.Name)
	
		steam2, err := servProduct.reposteam.GetData(steam.AppID)
		if err != nil {
			return &Domain{}, err
		}
		result.Price = steam2.Price
		fmt.Println(" finish GetData, steam2.Name: " + steam2.Name)
		fmt.Println(00000 + steam2.Price)

	} else if product.SellerID == 2{
		gog, err := servProduct.repogog.GetData("1447947499")
		if err != nil {
			println("ERROR GetData: ", err)
			return &Domain{}, err
		}
		result.URL = gog.URL
		fmt.Println(gog.URL)
	}

	return result, nil
}

func (servProduct *serviceProducts) FindAll() ([]Domain, error) {
	result, err := servProduct.repository.FindAll()
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (servProduct *serviceProducts) FindByID(id int) (*Domain, error) {
	result, err := servProduct.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servProduct *serviceProducts) Update(product *Domain, id int) (*Domain, error) {
	result, err := servProduct.repository.Update(product, id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servProduct *serviceProducts) Delete(product *Domain, id int) (string, error) {
	result, err := servProduct.repository.Delete(product, id)
	if err != nil {
		return "Fail to delete.", err
	}
	return result, nil
}
