package usecase

import "nik-extractor/domain"

type DistrictUseCase struct {
	repo domain.DistrictRepository
}

func NewDistrictUseCase(repo domain.DistrictRepository) domain.DistrictUseCase {
	return &DistrictUseCase{repo: repo}
}

func (p DistrictUseCase) FindById(id string) (domain.District, error) {
	return p.repo.FindById(id)
}
