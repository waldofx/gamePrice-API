package products

import (
	"fmt"
	"gameprice-api/business/steamapis"
)

type serviceProducts struct {
	repository Repository
	reposteam steamapis.Repository
}

func NewService(repoProduct Repository, rs steamapis.Repository) Service {
	return &serviceProducts{
		repository: repoProduct,
		reposteam: rs,
	}
}

func (servProduct *serviceProducts) Append(product *Domain) (*Domain, error) {
	result, err := servProduct.repository.Insert(product)
	if err != nil {
		return &Domain{}, err
	}
	fmt.Println(" finish init Insert")

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
