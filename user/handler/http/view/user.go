package view

type SubmitUserView struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type ExtractUserIdView struct {
	Id       string `json:"id"`
	Dob      string `json:"dob"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Gender   string `json:"gender"`
}

type ValidateUserDataView struct {
	DoB      string `json:"dob" binding:"required"`
	Province string `json:"province" binding:"required"`
	City     string `json:"city" binding:"required"`
	District string `json:"district" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Id       string `json:"id" binding:"required"`
}

type ValidateUserDataErrorView struct {
	Id     string   `json:"id"`
	Errors []string `json:"errors"`
}
