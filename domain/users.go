package domain

import (
	"database/sql"
	"nik-extractor/user/handler/http/view"
)

type User struct {
	Id   string
	Name string
}

type UserRepository interface {
	FindById(id string) (*User, error)
	FindUserByProvinceId(provinceId string) ([]User, error)
	FindUserByCityId(cityId string) ([]User, error)
	FindUserByDistrictId(districtId string) ([]User, error)
	FindUserByYearOfBirth(yearOfBirth string) ([]User, error)
	FindUserByGender(gender string) ([]User, error)
	Submit(users []User) error
	SubmitWithTx(*sql.Tx, User) error
	CleanUp() error
}

type UserUseCase interface {
	FindById(id string) (*User, error)
	FindUserByProvinceId(provinceId string) ([]User, error)
	FindUserByCityId(cityId string) ([]User, error)
	FindUserByDistrictId(districtId string) ([]User, error)
	FindUserByYearOfBirth(yearOfBirth string) ([]User, error)
	FindUserByGender(gender string) ([]User, error)
	Submit(users []view.SubmitUserView) error
	Extract(id string) (*view.ExtractUserIdView, []string)
	Validate(view []view.ValidateUserDataView) (bool, []view.ValidateUserDataErrorView)
	CleanUp() error
}
