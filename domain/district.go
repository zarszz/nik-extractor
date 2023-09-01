package domain

type District struct {
	Id   string
	Name string
}

type DistrictRepository interface {
	FindById(id string) (District, error)
}

type DistrictUseCase interface {
	FindById(id string) (District, error)
}
