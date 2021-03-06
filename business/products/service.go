package products

import (
	"gameprice-api/business/steamapis"
)

type serviceProducts struct {
	repository 	Repository
	reposteam 	steamapis.Repository
}

func NewService(repoProduct Repository, rs steamapis.Repository) Service {
	return &serviceProducts{
		repository: repoProduct,
		reposteam: rs,
	}
}

func (servProduct *serviceProducts) APIDetail(product *Domain) (*Domain, error) {
	if product.SellerID == 1 {
		steam, err := servProduct.reposteam.GetID(product.Game)
		if err != nil {
			return &Domain{}, err
		}
		steam2, err := servProduct.reposteam.GetData(steam.AppID)
		if err != nil {
			return &Domain{}, err
		}
		product.Discount = steam2.Discount
		product.Price = steam2.Price
		product.URL = ("https://store.steampowered.com/app/" + steam.AppID)

	} else if product.SellerID == 2{
		product.Price = "Price is not available from this seller."
		product.URL = ("https://www.gog.com/game/"+product.Game)
	}
	return product, nil
}

func (servProduct *serviceProducts) Append(product *Domain) (*Domain, error) {
	result, err := servProduct.repository.Insert(product)
	if err != nil {
		return &Domain{}, err
	}
	servProduct.APIDetail(result)
	servProduct.Update(result, result.ID)
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
	result, err = servProduct.APIDetail(result)
	if err != nil {
		return &Domain{}, err
	}
	servProduct.repository.Update(result, id)
	return result, nil
}

func (servProduct *serviceProducts) Delete(product *Domain, id int) (string, error) {
	result, err := servProduct.repository.Delete(product, id)
	if err != nil {
		return "Fail to delete.", err
	}
	return result, nil
}
