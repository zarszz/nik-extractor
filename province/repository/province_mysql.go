package repository

import (
	"database/sql"
	"nik-extractor/domain"
)

type ProvinceMysqlRepository struct {
	db *sql.DB
}

func NewProvinceRepository(db *sql.DB) domain.ProvinceRepository {
	return &ProvinceMysqlRepository{db: db}
}

func (p ProvinceMysqlRepository) FindById(id string) (domain.Province, error) {
	var province domain.Province
	err := p.db.QueryRow("SELECT id, nama FROM provinces WHERE id = ?", id).Scan(&province.Id, &province.Name)
	if err != nil {
		return province, err
	}
	return province, nil
}
