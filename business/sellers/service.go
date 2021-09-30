package sellers

type serviceSellers struct {
	repository Repository
}

func NewService(repoSeller Repository) Service {
	return &serviceSellers{
		repository: repoSeller,
	}
}

func (servSeller *serviceSellers) Append(seller *Domain) (*Domain, error) {
	result, err := servSeller.repository.Insert(seller)
	if err != nil {
		return &Domain{}, err
	}

	return result, nil
}

func (servSeller *serviceSellers) FindAll() ([]Domain, error) {
	result, err := servSeller.repository.FindAll()
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (servSeller *serviceSellers) FindByID(id int) (*Domain, error) {
	result, err := servSeller.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servSeller *serviceSellers) Update(seller *Domain) (*Domain, error) {
	result, err := servSeller.repository.Update(seller)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servSeller *serviceSellers) Delete(seller *Domain, id int) (string, error) {
	result, err := servSeller.repository.Delete(seller, id)
	if err != nil {
		return "Fail to delete.", err
	}
	return result, nil
}
