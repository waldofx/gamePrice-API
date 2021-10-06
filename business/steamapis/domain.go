package steamapis

type Domain struct {
	AppID    string
	Name     string
	Price    string
	Discount bool
}

type Repository interface {
	GetID(name string) (Domain, error)
	GetData(appid string) (Domain, error)
}