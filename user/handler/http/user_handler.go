package http

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	view2 "nik-extractor/app/view"
	"nik-extractor/domain"
	"nik-extractor/helper"
	"nik-extractor/user/handler/http/view"
)

type UserHandler struct {
	userUseCase domain.UserUseCase
}

func NewUserHandler(gin *gin.Engine, userUseCase domain.UserUseCase) {
	handler := &UserHandler{
		userUseCase: userUseCase,
	}

	gin.POST("/api/v1/submit-ids", handler.Submit)
	gin.GET("/api/v1/users/province/:province_id", handler.FindByProvinceId)
	gin.GET("/api/v1/users/city/:city_id", handler.FindByCityId)
	gin.GET("/api/v1/users/district/:district_id", handler.FindByDistrictId)
	gin.GET("/api/v1/users/year/:year_of_birth", handler.FindByYearOfBirth)
	gin.GET("/api/v1/users/gender/:gender", handler.FindByGender)
	gin.GET("/api/v1/extract/:id", handler.Extract)
	gin.POST("/api/v1/validate", handler.Validate)
	gin.DELETE("/api/v1/clean-up", handler.CleanUp)
}

// Submit  godoc
// @Summary      Submit list of ID to database
// @Description  Submit list of ID to database
// @Tags         users
// @Accept       json
// @Produce      json
// @Param request body []view.SubmitUserView true "query params"
// @Success      200  {object}  view.Response
// @Failure      400  {object}  view.Response
// @Failure      404  {object}  view.Response
// @Failure      500  {object}  view.Response
// @Router       /submit-ids [post]
func (p *UserHandler) Submit(c *gin.Context) {
	var submitUserView []view.SubmitUserView

	if err := c.ShouldBindJSON(&submitUserView); err != nil {
		view2.MakeResponse(c, 400, err.Error(), gin.H{
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
		if errors.Is(err, helper.ErrConflict) {
			view2.MakeResponse(c, 400, "Id already used", []domain.User{})
			return
		}
		view2.MakeResponse(c, 400, err.Error(), []domain.User{})
		return
	}

	view2.MakeResponse(c, 200, "Success", []domain.User{})
}

// FindByProvinceId  godoc
// @Summary      get all users by their province id
// @Description  get all users by their province id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Province ID"
// @Success      200  {object}  view.Response
// @Failure      404  {object}  view.Response
// @Failure      500  {object}  view.Response
// @Router       /users/province/:province_id [get]
func (p *UserHandler) FindByProvinceId(c *gin.Context) {
	provinceId := c.Param("province_id")

	users, err := p.userUseCase.FindUserByProvinceId(provinceId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			view2.MakeResponse(c, 404, "province not found", []domain.User{})
			return
		}
		view2.MakeResponse(c, 400, err.Error(), []domain.User{})
		return
	}

	view2.MakeResponse(c, 200, "Success", users)
}

// FindByCityId  godoc
// @Summary      get all users by their city id
// @Description  get all users by their city id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "City ID"
// @Success      200  {object}  view.Response
// @Failure      404  {object}  view.Response
// @Failure      500  {object}  view.Response
// @Router       /users/city/:city_id [get]
func (p *UserHandler) FindByCityId(c *gin.Context) {
	cityId := c.Param("city_id")

	users, err := p.userUseCase.FindUserByCityId(cityId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			view2.MakeResponse(c, 404, "city not found", []domain.User{})
			return
		}
		view2.MakeResponse(c, 400, err.Error(), []domain.User{})
		return
	}

	c.JSON(200, users)
}

// FindByDistrictId  godoc
// @Summary      get all users by their district id
// @Description  get all users by their district id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "District ID"
// @Success      200  {object}  view.Response
// @Failure      404  {object}  view.Response
// @Failure      500  {object}  view.Response
// @Router       /users/district/:district_id [get]
func (p *UserHandler) FindByDistrictId(c *gin.Context) {
	districtId := c.Param("district_id")

	users, err := p.userUseCase.FindUserByDistrictId(districtId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			view2.MakeResponse(c, 404, "district not found", []domain.User{})
			return
		}
		view2.MakeResponse(c, 400, err.Error(), []domain.User{})
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
		view2.MakeResponse(c, 400, err.Error(), []domain.User{})
		return
	}

	c.JSON(200, users)
}

// FindByGender  godoc
// @Summary      Extract a data from id with data from database
// @Description  Extract a data from id with data from database
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  view.Response
// @Failure      400  {object}  view.Response
// @Failure      404  {object}  view.Response
// @Failure      500  {object}  view.Response
// @Router       /users/:id [get]
func (p *UserHandler) FindByGender(c *gin.Context) {
	gender := c.Param("gender")

	users, err := p.userUseCase.FindUserByGender(gender)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			view2.MakeResponse(c, 404, "Not Found", []domain.User{})
			return
		}
		view2.MakeResponse(c, 400, err.Error(), []domain.User{})
		return
	}

	c.JSON(200, users)
}

// Extract  godoc
// @Summary      Extract a data from id with data from database
// @Description  Extract a data from id with data from database
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  view.Response
// @Failure      400  {object}  view.Response
// @Failure      404  {object}  view.Response
// @Failure      500  {object}  view.Response
// @Router       /extract/:id [get]
func (p *UserHandler) Extract(context *gin.Context) {
	extractedUserIdView, err := p.userUseCase.Extract(context.Param("id"))
	if len(err) > 0 {
		view2.MakeResponse(context, 400, "Bad Request", err)
		return
	}

	view2.MakeResponse(context, 200, "Success", extractedUserIdView)
}

// Validate  godoc
// @Summary      Validate list of id with data to database
// @Description  Validate list of id with data to database
// @Tags         users
// @Accept       json
// @Produce      json
// @Param request body []	view.ValidateUserDataView true "query params"
// @Success      200  {object}  view.Response
// @Failure      400  {object}  view.Response
// @Failure      404  {object}  view.Response
// @Failure      500  {object}  view.Response
// @Router       /validate [post]
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

// CleanUp  godoc
// @Summary      CleanUp all data inside db - for test only
// @Description  CleanUp all data inside db - for test only
// @Tags         misc
// @Accept       json
// @Produce      json
// @Success      200  {object}  view.Response
// @Failure      500  {object}  view.Response
// @Router       /clean-up [post]
func (p *UserHandler) CleanUp(context *gin.Context) {
	err := p.userUseCase.CleanUp()
	if err != nil {
		view2.MakeResponse(context, 500, "Internal Server Error", []domain.User{})
		return
	}
	view2.MakeResponse(context, 200, "Success", []domain.User{})
}
