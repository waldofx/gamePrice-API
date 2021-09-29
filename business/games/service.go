package games

type serviceGames struct {
	repository Repository
}

func NewService(repoGame Repository) Service {
	return &serviceGames{
		repository: repoGame,
	}
}

func (servGame *serviceGames) Append(game *Domain) (*Domain, error) {
	result, err := servGame.repository.Insert(game)
	if err != nil {
		return &Domain{}, err
	}

	return result, nil
}

func (servGame *serviceGames) FindAll() ([]Domain, error) {
	result, err := servGame.repository.FindAll()
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (servGame *serviceGames) FindByID(id int) (*Domain, error) {
	result, err := servGame.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servGame *serviceGames) Update(game *Domain) (*Domain, error) {
	//return &Domain{}, nil
	result, err := servGame.repository.Update(game)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servGame *serviceGames) Delete(game *Domain, id int) (string, error) {
	result, err := servGame.repository.Delete(game, id)
	if err != nil {
		return "Fail to delete.", err
	}
	return result, nil
}
