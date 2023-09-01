package usecase

import "nik-extractor/domain"

type ProvinceUseCase struct {
	repo domain.ProvinceRepository
}

func NewProvinceUseCase(repo domain.ProvinceRepository) domain.ProvinceUseCase {
	return &ProvinceUseCase{repo: repo}
}

func (p ProvinceUseCase) FindById(id string) (domain.Province, error) {
	return p.repo.FindById(id)
}
