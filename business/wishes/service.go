package wishes

type serviceWishes struct {
	repository Repository
}

func NewService(repoProduct Repository) Service {
	return &serviceWishes{
		repository: repoProduct,
	}
}

func (servProduct *serviceWishes) Append(wish *Domain) (*Domain, error) {
	result, err := servProduct.repository.Insert(wish)
	if err != nil {
		return &Domain{}, err
	}

	return result, nil
}

func (servProduct *serviceWishes) FindAll() ([]Domain, error) {
	result, err := servProduct.repository.FindAll()
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (servProduct *serviceWishes) FindByID(id int) (*Domain, error) {
	result, err := servProduct.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servProduct *serviceWishes) Update(wish *Domain, id int) (*Domain, error) {
	result, err := servProduct.repository.Update(wish, id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servProduct *serviceWishes) Delete(wish *Domain, id int) (string, error) {
	result, err := servProduct.repository.Delete(wish, id)
	if err != nil {
		return "Fail to delete.", err
	}
	return result, nil
}
