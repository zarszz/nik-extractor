package usecase

import "nik-extractor/domain"

type CityUseCase struct {
	repo domain.CityRepository
}

func NewCityUseCase(repo domain.CityRepository) domain.CityUseCase {
	return &CityUseCase{repo: repo}
}

func (p CityUseCase) FindById(id string) (domain.City, error) {
	return p.repo.FindById(id)
}
