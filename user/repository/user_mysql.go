package repository

import (
	"database/sql"
	"nik-extractor/domain"
)

type UserMysqlRepository struct {
	db *sql.DB
}

// CleanUp Only for testing purpose
func (p UserMysqlRepository) CleanUp() error {
	stmt, err := p.db.Prepare("DELETE FROM users WHERE TRUE")
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (p UserMysqlRepository) FindUserByYearOfBirth(yearOfBirth string) ([]domain.User, error) {
	var users []domain.User
	stmt, _ := p.db.Query("SELECT id, name FROM users WHERE SUBSTRING(id, 11, 2) = ?", yearOfBirth)
	for stmt.Next() {
		var user domain.User
		err := stmt.Scan(&user.Id, &user.Name)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (p UserMysqlRepository) FindUserByGender(gender string) ([]domain.User, error) {
	var users []domain.User
	query := "SELECT id, name FROM users where CAST(SUBSTRING(id, 7, 2) as UNSIGNED) "
	if gender == "m" {
		query += "< 40"
	} else {
		query += "> 40"
	}
	stmt, _ := p.db.Query(query)
	for stmt.Next() {
		var user domain.User
		err := stmt.Scan(&user.Id, &user.Name)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (p UserMysqlRepository) FindUserByDistrictId(districtId string) ([]domain.User, error) {
	var users []domain.User
	stmt, _ := p.db.Query("SELECT id, name FROM users WHERE SUBSTRING(id, 1, 6) = ?", districtId)
	for stmt.Next() {
		var user domain.User
		err := stmt.Scan(&user.Id, &user.Name)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (p UserMysqlRepository) FindUserByCityId(cityId string) ([]domain.User, error) {
	var users []domain.User
	stmt, _ := p.db.Query("SELECT id, name FROM users WHERE SUBSTRING(id, 1, 4) = ?", cityId)
	for stmt.Next() {
		var user domain.User
		err := stmt.Scan(&user.Id, &user.Name)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &UserMysqlRepository{db: db}
}

func (p UserMysqlRepository) FindById(id string) (domain.User, error) {
	var province domain.User
	err := p.db.QueryRow("SELECT id, name FROM provinces WHERE id = ?", id).Scan(&province.Id, &province.Name)
	if err != nil {
		return province, err
	}
	return province, nil
}

func (p UserMysqlRepository) Submit(users []domain.User) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO users(id, name) VALUES(?, ?)")
	if err != nil {
		return err
	}
	for _, user := range users {
		_, err := stmt.Exec(user.Id, user.Name)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return err
			}
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (p UserMysqlRepository) SubmitWithTx(tx *sql.Tx, user domain.User) error {
	stmt, _ := tx.Prepare("INSERT INTO users(id, name) VALUES(?, ?)")
	_, err := stmt.Exec(user.Id, user.Name)
	if err != nil {
		if err != nil {
			return err
		}
		return err
	}

	if err != nil {
		return err
	}
	return nil
}

func (p UserMysqlRepository) FindUserByProvinceId(provinceId string) ([]domain.User, error) {
	var users []domain.User
	stmt, _ := p.db.Query("SELECT id, name FROM users WHERE SUBSTRING(id, 1, 2) = ?", provinceId)
	for stmt.Next() {
		var user domain.User
		err := stmt.Scan(&user.Id, &user.Name)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
