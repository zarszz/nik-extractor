package repository

import (
	"database/sql"
	"nik-extractor/domain"
)

type DistrictMysqlRepository struct {
	db *sql.DB
}

func NewDistrictRepository(db *sql.DB) domain.DistrictRepository {
	return &DistrictMysqlRepository{db: db}
}

func (p DistrictMysqlRepository) FindById(id string) (domain.District, error) {
	var district domain.District
	err := p.db.QueryRow("SELECT id, nama FROM districts WHERE id = ?", id).Scan(&district.Id, &district.Name)
	if err != nil {
		return district, err
	}
	return district, nil
}
