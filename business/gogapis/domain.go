package gogapis

type Domain struct {
	AppID string
	Name  string
	URL   string
	Price int
}

type Repository interface {
	GetData(appid string) (Domain, error)
}