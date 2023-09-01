package domain

type Province struct {
	Id   string
	Name string
}

type ProvinceRepository interface {
	FindById(id string) (Province, error)
}

type ProvinceUseCase interface {
	FindById(id string) (Province, error)
}
