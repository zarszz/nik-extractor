package domain

type City struct {
	Id   string
	Name string
}

type CityRepository interface {
	FindById(id string) (City, error)
}

type CityUseCase interface {
	FindById(id string) (City, error)
}
