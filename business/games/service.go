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
func (servGame *serviceGames) Update(game *Domain, id int) (*Domain, error) {
	return &Domain{}, nil
}
func (servGame *serviceGames) FindByID(id int) (*Domain, error) {
	return &Domain{}, nil
}
func (servGame *serviceGames) Available(generalSearch string) []Domain {
	return []Domain{}
}
