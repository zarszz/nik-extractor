package http

import (
	"github.com/gin-gonic/gin"
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
}

func (p *UserHandler) Submit(c *gin.Context) {
	var submitUserView []view.SubmitUserView

	if err := c.ShouldBindJSON(&submitUserView); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	err := p.userUseCase.Submit(submitUserView)

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
	})
}

func (p *UserHandler) FindByProvinceId(c *gin.Context) {
	provinceId := c.Param("province_id")

	users, err := p.userUseCase.FindUserByProvinceId(provinceId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, users)
}

func (p *UserHandler) FindByCityId(c *gin.Context) {
	cityId := c.Param("city_id")

	users, err := p.userUseCase.FindUserByCityId(cityId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, users)
}

func (p *UserHandler) FindByDistrictId(c *gin.Context) {
	districtId := c.Param("district_id")

	users, err := p.userUseCase.FindUserByDistrictId(districtId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, users)
}

func (p *UserHandler) FindByYearOfBirth(c *gin.Context) {
	yearOfBirth := c.Param("year_of_birth")

	users, err := p.userUseCase.FindUserByYearOfBirth(yearOfBirth)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, users)
}

func (p *UserHandler) FindByGender(c *gin.Context) {
	gender := c.Param("gender")

	users, err := p.userUseCase.FindUserByGender(gender)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, users)
}

func (p *UserHandler) Extract(context *gin.Context) {
	extractedUserIdView, err := p.userUseCase.Extract(context.Param("id"))
	if err.Error() != "" {
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(200, extractedUserIdView)
}

func (p *UserHandler) Validate(context *gin.Context) {
	var validateUserDataView []view.ValidateUserDataView

	if err := context.ShouldBindJSON(&validateUserDataView); err != nil {
		context.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	isContainError, errorViews := p.userUseCase.Validate(validateUserDataView)
	if isContainError {
		context.JSON(400, errorViews)
		return
	}
	context.JSON(200, gin.H{"data": errorViews})
}
