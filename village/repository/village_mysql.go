package repository

import (
	"database/sql"
	"nik-extractor/domain"
)

type VillageMysqlRepository struct {
	db *sql.DB
}

func NewVillageRepository(db *sql.DB) domain.VillageRepository {
	return &VillageMysqlRepository{db: db}
}

func (p VillageMysqlRepository) FindById(id string) (domain.Village, error) {
	var village domain.Village
	err := p.db.QueryRow("SELECT id, nama FROM villages WHERE id = ?", id).Scan(&village.Id, &village.Name)
	if err != nil {
		return village, err
	}
	return village, nil
}
