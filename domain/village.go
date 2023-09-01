package domain

type Village struct {
	Id   string
	Name string
}

type VillageRepository interface {
	FindById(id string) (Village, error)
}

type VillageUseCase interface {
	FindById(id string) (Village, error)
}
