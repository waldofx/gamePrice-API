package steamapis

type Domain struct {
	AppID string
	Name  string
	Price int
}

type Repository interface {
	GetID(name string) (Domain, error)
	GetData(appid string) (Domain, error)
}