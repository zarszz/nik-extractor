package usecase

import "nik-extractor/domain"

type VillageUseCase struct {
	repo domain.VillageRepository
}

func NewVillageUseCase(repo domain.VillageRepository) domain.VillageUseCase {
	return &VillageUseCase{repo: repo}
}

func (p VillageUseCase) FindById(id string) (domain.Village, error) {
	return p.repo.FindById(id)
}
