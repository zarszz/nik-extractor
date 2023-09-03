package http

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	view2 "nik-extractor/app/view"
	"nik-extractor/domain"
	"nik-extractor/user/handler/http/view"
)

type UserHandler struct {
	userUseCase domain.UserUseCase
}

func NewUserHandler(gin *gin.Engine, userUseCase domain.UserUseCase) {
	handler := &UserHandler{
		userUseCase: userUseCase,
	}

	gin.POST("/v1/submit-ids", handler.Submit)
	gin.GET("/v1/users/province/:province_id", handler.FindByProvinceId)
	gin.GET("/v1/users/city/:city_id", handler.FindByCityId)
	gin.GET("/v1/users/district/:district_id", handler.FindByDistrictId)
	gin.GET("/v1/users/year/:year_of_birth", handler.FindByYearOfBirth)
	gin.GET("/v1/users/gender/:gender", handler.FindByGender)
	gin.GET("/v1/extract/:id", handler.Extract)
	gin.POST("/v1/validate", handler.Validate)
	gin.DELETE("v1/clean-up", handler.CleanUp)
}

func (p *UserHandler) Submit(c *gin.Context) {
	var submitUserView []view.SubmitUserView

	if err := c.ShouldBindJSON(&submitUserView); err != nil {
		view2.MakeResponse(c, 400, "Internal Server Error", gin.H{
			"error": err.Error(),
		})
		return
	}

	err := p.userUseCase.Submit(submitUserView)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			view2.MakeResponse(c, 404, "Not Found", []domain.User{})
			return
		}
		view2.MakeResponse(c, 500, "Internal Server Error", []domain.User{})
		return
	}

	view2.MakeResponse(c, 200, "Success", []domain.User{})
}

func (p *UserHandler) FindByProvinceId(c *gin.Context) {
	provinceId := c.Param("province_id")

	users, err := p.userUseCase.FindUserByProvinceId(provinceId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			view2.MakeResponse(c, 404, "province not found", []domain.User{})
			return
		}
		view2.MakeResponse(c, 500, "Internal Server Error", []domain.User{})
		return
	}

	view2.MakeResponse(c, 200, "Success", users)
}

func (p *UserHandler) FindByCityId(c *gin.Context) {
	cityId := c.Param("city_id")

	users, err := p.userUseCase.FindUserByCityId(cityId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			view2.MakeResponse(c, 404, "city not found", []domain.User{})
			return
		}
		view2.MakeResponse(c, 500, "Internal Server Error", []domain.User{})
		return
	}

	c.JSON(200, users)
}

func (p *UserHandler) FindByDistrictId(c *gin.Context) {
	districtId := c.Param("district_id")

	users, err := p.userUseCase.FindUserByDistrictId(districtId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			view2.MakeResponse(c, 404, "district not found", []domain.User{})
			return
		}
		view2.MakeResponse(c, 500, "Internal Server Error", []domain.User{})
		return
	}
	view2.MakeResponse(c, 200, "Success", users)
}

func (p *UserHandler) FindByYearOfBirth(c *gin.Context) {
	yearOfBirth := c.Param("year_of_birth")

	users, err := p.userUseCase.FindUserByYearOfBirth(yearOfBirth)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			view2.MakeResponse(c, 404, "Not Found", []domain.User{})
			return
		}
		view2.MakeResponse(c, 500, err.Error(), []domain.User{})
		return
	}

	c.JSON(200, users)
}

func (p *UserHandler) FindByGender(c *gin.Context) {
	gender := c.Param("gender")

	users, err := p.userUseCase.FindUserByGender(gender)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			view2.MakeResponse(c, 404, "Not Found", []domain.User{})
			return
		}
		view2.MakeResponse(c, 500, err.Error(), []domain.User{})
		return
	}

	c.JSON(200, users)
}

func (p *UserHandler) Extract(context *gin.Context) {
	extractedUserIdView, err := p.userUseCase.Extract(context.Param("id"))
	if len(err) > 0 {
		view2.MakeResponse(context, 400, "Bad Request", err)
		return
	}

	view2.MakeResponse(context, 200, "Success", extractedUserIdView)
}

func (p *UserHandler) Validate(context *gin.Context) {
	var validateUserDataView []view.ValidateUserDataView

	if err := context.ShouldBindJSON(&validateUserDataView); err != nil {
		view2.MakeResponse(context, 400, "Bad Request", gin.H{
			"error": err.Error(),
		})
		return
	}

	isContainError, errorViews := p.userUseCase.Validate(validateUserDataView)
	if !isContainError {
		view2.MakeResponse(context, 400, "Bad Request", errorViews)
		return
	}
	view2.MakeResponse(context, 200, "Success", []domain.User{})
}

func (p *UserHandler) CleanUp(context *gin.Context) {
	err := p.userUseCase.CleanUp()
	if err != nil {
		view2.MakeResponse(context, 500, "Internal Server Error", []domain.User{})
		return
	}
	view2.MakeResponse(context, 200, "Success", []domain.User{})
}
