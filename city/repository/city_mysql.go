package repository

import (
	"database/sql"
	"nik-extractor/domain"
)

type CityMysqlRepository struct {
	db *sql.DB
}

func NewCityRepository(db *sql.DB) domain.CityRepository {
	return &CityMysqlRepository{db: db}
}

func (p CityMysqlRepository) FindById(id string) (domain.City, error) {
	var province domain.City
	err := p.db.QueryRow("SELECT id, nama FROM cities WHERE id = ?", id).Scan(&province.Id, &province.Name)
	if err != nil {
		return province, err
	}
	return province, nil
}
