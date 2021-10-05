package wishes

import (
	"gameprice-api/business/products"
)

type serviceWishes struct {
	repository  Repository
	repoProduct products.Repository
}

func NewService(repoWish Repository, repoProduct products.Repository) Service {
	return &serviceWishes{
		repository: repoWish,
		repoProduct: repoProduct,
	}
}

func (servWish *serviceWishes) Append(wish *Domain) (*Domain, error) {
	wish.ProductID, wish.Price, wish.Discount, wish.URL = servWish.repoProduct.GetProduct(wish.GameID, wish.SellerID)
	result, err := servWish.repository.Insert(wish)
	if err != nil {
		return &Domain{}, err
	}

	return result, nil
}

func (servWish *serviceWishes) FindAll() ([]Domain, error) {
	result, err := servWish.repository.FindAll()
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (servWish *serviceWishes) FindByID(id int) (*Domain, error) {
	result, err := servWish.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servWish *serviceWishes) Update(wish *Domain, id int) (*Domain, error) {
	result, err := servWish.repository.Update(wish, id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servWish *serviceWishes) Delete(wish *Domain, id int) (string, error) {
	result, err := servWish.repository.Delete(wish, id)
	if err != nil {
		return "Fail to delete.", err
	}
	return result, nil
}
