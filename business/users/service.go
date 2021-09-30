package users

type serviceUsers struct {
	repository Repository
}

func NewService(repoUser Repository) Service {
	return &serviceUsers{
		repository: repoUser,
	}
}

func (servUser *serviceUsers) Append(user *Domain) (*Domain, error) {
	result, err := servUser.repository.Insert(user)
	if err != nil {
		return &Domain{}, err
	}

	return result, nil
}

func (servUser *serviceUsers) FindAll() ([]Domain, error) {
	result, err := servUser.repository.FindAll()
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (servUser *serviceUsers) FindByID(id int) (*Domain, error) {
	result, err := servUser.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servUser *serviceUsers) Update(user *Domain) (*Domain, error) {
	result, err := servUser.repository.Update(user)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servUser *serviceUsers) Delete(user *Domain, id int) (string, error) {
	result, err := servUser.repository.Delete(user, id)
	if err != nil {
		return "Fail to delete.", err
	}
	return result, nil
}
