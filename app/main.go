package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	_cityRepository "nik-extractor/city/repository"
	_districtRepository "nik-extractor/district/repository"
	_provinceRepository "nik-extractor/province/repository"
	_userRepository "nik-extractor/user/repository"
	_userUseCase "nik-extractor/user/usecase"

	_userHandler "nik-extractor/user/handler/http"
)

// Define your database connection parameters here
const (
	dbDriver   = "mysql"
	dbUser     = "root"
	dbPassword = ""
	dbName     = "nik_extractor"
)

func main() {
	db, err := sql.Open(dbDriver, fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", dbUser, dbPassword, dbName))
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Failed to close the database connection:", err)
		}
	}(db)

	r := gin.Default()

	cityRepo := _cityRepository.NewCityRepository(db)
	provinceRepo := _provinceRepository.NewProvinceRepository(db)
	districtRepo := _districtRepository.NewDistrictRepository(db)
	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo, provinceRepo, cityRepo, districtRepo)
	_userHandler.NewUserHandler(r, userUseCase)

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
