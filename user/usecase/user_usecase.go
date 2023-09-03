package usecase

import (
	"errors"
	"fmt"
	"nik-extractor/domain"
	"nik-extractor/helper"
	"nik-extractor/user/handler/http/view"
	"strconv"
	"time"
)

type UserUseCase struct {
	repo         domain.UserRepository
	provinceRepo domain.ProvinceRepository
	cityRepo     domain.CityRepository
	districtRepo domain.DistrictRepository
	villageRepo  domain.VillageRepository
}

func (p UserUseCase) CleanUp() error {
	return p.repo.CleanUp()
}

func (p UserUseCase) Validate(validateUserDataView []view.ValidateUserDataView) (bool, []view.ValidateUserDataErrorView) {
	var errorViews []view.ValidateUserDataErrorView

	for _, data := range validateUserDataView {
		var mismatchedData []string
		extractedMetaData, err := p.Extract(data.Id)
		if err != nil {
			if len(err) != 0 {
				errorViews = append(errorViews, view.ValidateUserDataErrorView{
					Id:     data.Id,
					Errors: err,
				})
				continue
			}
		}

		if extractedMetaData.Gender != data.Gender {
			mismatchedData = append(mismatchedData, "gender mismatch")
		}

		if extractedMetaData.Dob != data.DoB {
			mismatchedData = append(mismatchedData, "date of birth mismatch")
		}

		if extractedMetaData.Province != data.Province {
			mismatchedData = append(mismatchedData, "province mismatch")
		}

		if extractedMetaData.City != data.City {
			mismatchedData = append(mismatchedData, "city mismatch")
		}

		if extractedMetaData.District != data.District {
			mismatchedData = append(mismatchedData, "district mismatch")
		}

		if len(mismatchedData) > 0 {
			errorViews = append(errorViews, view.ValidateUserDataErrorView{
				Id:     data.Id,
				Errors: mismatchedData,
			})
		}
	}

	return len(errorViews) == 0, errorViews
}

func CalculateYearOfBirth(twoDigitsOfBirth string) (string, error) {
	// Get the current year
	currentYear := time.Now().Year()

	// Extract the last two digits of the current year
	lastTwoDigitsCurrentYear := currentYear % 100

	// Convert the provided two digits to an integer
	twoDigitsInt, err := strconv.Atoi(twoDigitsOfBirth)
	if err != nil {
		return "", err
	}

	// Determine the full year based on the provided two digits
	var fullYear int

	if twoDigitsInt >= lastTwoDigitsCurrentYear {
		// If the provided two digits are greater or equal to the last two digits of the current year,
		// assume it's in the 20th century (e.g., "00" becomes "2000").
		fullYear = 1900 + twoDigitsInt
	} else {
		// Otherwise, assume it's in the 21st century (e.g., "01" becomes "2001").
		fullYear = 2000 + twoDigitsInt
	}

	return strconv.Itoa(fullYear), nil
}

func (p UserUseCase) Extract(id string) (*view.ExtractUserIdView, []string) {
	var extractUserIdView view.ExtractUserIdView
	var mismatches []string

	provinceCode := id[:2]
	province, err := p.provinceRepo.FindById(provinceCode)
	if err != nil {
		mismatches = append(mismatches, "invalid province id")
	} else {
		extractUserIdView.Province = province.Name
	}

	cityCode := fmt.Sprintf("%s%s", provinceCode, id[2:4])
	city, err := p.cityRepo.FindById(cityCode)
	if err != nil {
		mismatches = append(mismatches, "invalid city id")
	} else {
		extractUserIdView.City = city.Name
	}

	districtCode := fmt.Sprintf("%s%s", cityCode, id[4:6])
	district, err := p.districtRepo.FindById(districtCode)
	if err != nil {
		mismatches = append(mismatches, "invalid district id")
	} else {
		extractUserIdView.District = district.Name
	}

	var intDateOfBirth, intMonthOfBirth int
	dateOfBirth := id[6:8]
	if len(dateOfBirth) != 2 {
		mismatches = append(mismatches, "invalid date of birth")
	} else {
		if dateOfBirth[0] == '0' {
			intDateOfBirth, _ = strconv.Atoi(dateOfBirth[1:])
		} else {
			intDateOfBirth, _ = strconv.Atoi(dateOfBirth)
		}
	}

	if intDateOfBirth > 40 {
		intDateOfBirth -= 40
	}

	if (intDateOfBirth < 1) || (intDateOfBirth > 31) {
		mismatches = append(mismatches, "invalid date of birth")
	}

	if intDateOfBirth < 10 {
		dateOfBirth = fmt.Sprintf("0%d", intDateOfBirth)
	} else {
		dateOfBirth = fmt.Sprintf("%d", intDateOfBirth)
	}

	monthOfBirth := id[8:10]
	if len(monthOfBirth) != 2 {
		mismatches = append(mismatches, "invalid month of birth")
	} else {
		if monthOfBirth[0] == '0' {
			intMonthOfBirth, _ = strconv.Atoi(monthOfBirth[1:])
		} else {
			intMonthOfBirth, _ = strconv.Atoi(monthOfBirth)
		}
	}

	if (intMonthOfBirth < 1) || (intMonthOfBirth > 12) {
		mismatches = append(mismatches, "invalid month of birth")
	}

	if intMonthOfBirth < 10 {
		monthOfBirth = fmt.Sprintf("0%d", intMonthOfBirth)
	}

	lastTwoDigitYearOfBirth := id[10:12]
	yearOfBirth, err := CalculateYearOfBirth(lastTwoDigitYearOfBirth)
	if err != nil {
		mismatches = append(mismatches, "invalid year of birth")
	}
	extractUserIdView.Dob = fmt.Sprintf("%s-%s-%s", dateOfBirth, monthOfBirth, yearOfBirth)

	if intDateOfBirth <= 31 {
		extractUserIdView.Gender = "m"
	} else {
		extractUserIdView.Gender = "f"
	}
	extractUserIdView.Id = id
	return &extractUserIdView, mismatches
}

func (p UserUseCase) FindUserByYearOfBirth(yearOfBirth string) ([]domain.User, error) {
	users, err := p.repo.FindUserByYearOfBirth(yearOfBirth)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (p UserUseCase) FindUserByGender(gender string) ([]domain.User, error) {
	users, err := p.repo.FindUserByGender(gender)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (p UserUseCase) FindUserByDistrictId(districtId string) ([]domain.User, error) {
	_, err := p.districtRepo.FindById(districtId)
	if err != nil {
		return nil, err
	}
	users, err := p.repo.FindUserByDistrictId(districtId)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (p UserUseCase) FindUserByCityId(cityId string) ([]domain.User, error) {
	_, err := p.cityRepo.FindById(cityId)
	if err != nil {
		return nil, err
	}
	users, err := p.repo.FindUserByCityId(cityId)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserUseCase(repo domain.UserRepository, provinceRepo domain.ProvinceRepository,
	cityRepo domain.CityRepository,
	districtRepo domain.DistrictRepository) domain.UserUseCase {

	return &UserUseCase{repo: repo, provinceRepo: provinceRepo, cityRepo: cityRepo, districtRepo: districtRepo}
}

func (p UserUseCase) Submit(users []view.SubmitUserView) error {
	var submitUser []domain.User
	for _, user := range users {
		existingUser, err := p.repo.FindById(user.Id)
		if err != nil {
			return errors.New(fmt.Sprintf("Invalid ID %s", user.Id))
		}

		if existingUser != nil {
			return helper.ErrConflict
		}

		provinceCode := user.Id[:2]
		_, err = p.provinceRepo.FindById(provinceCode)
		if err != nil {
			return errors.New(fmt.Sprintf("Invalid ID %s", user.Id))
		}

		cityCode := fmt.Sprintf("%s%s", provinceCode, user.Id[2:4])
		_, err = p.cityRepo.FindById(cityCode)
		if err != nil {
			return errors.New(fmt.Sprintf("Invalid ID %s", user.Id))
		}

		districtCode := fmt.Sprintf("%s%s", cityCode, user.Id[4:6])
		_, err = p.districtRepo.FindById(districtCode)
		if err != nil {
			return errors.New(fmt.Sprintf("Invalid ID %s", user.Id))
		}

		submitUser = append(submitUser, domain.User{
			Id:   user.Id,
			Name: user.Name,
		})
	}
	err := p.repo.Submit(submitUser)
	if err != nil {
		return errors.New(fmt.Sprintf("Invalid ID %s", err.Error()))
	}
	return nil
}

func (p UserUseCase) FindById(id string) (*domain.User, error) {
	return p.repo.FindById(id)
}

func (p UserUseCase) FindUserByProvinceId(provinceId string) ([]domain.User, error) {
	_, err := p.provinceRepo.FindById(provinceId)
	if err != nil {
		return nil, err
	}
	users, err := p.repo.FindUserByProvinceId(provinceId)
	if err != nil {
		return nil, err
	}
	return users, nil
}
