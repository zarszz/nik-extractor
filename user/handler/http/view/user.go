package view

type SubmitUserView struct {
	Id   string `json:"id" binding:"required,numeric,min=16,max=16,len=16"`
	Name string `json:"name" binding:"required"`
}

type ExtractUserIdView struct {
	Id       string `json:"id" binding:"required,numeric,min=16,max=16,len=16"`
	Dob      string `json:"dob" binding:"required"`
	Province string `json:"province" binding:"required"`
	City     string `json:"city" binding:"required"`
	District string `json:"district" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
}

type ValidateUserDataViewResponse struct {
	DoB      string `json:"dob" binding:"required"`
	Province string `json:"province" binding:"required"`
	City     string `json:"city" binding:"required"`
	District string `json:"district" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Id       string `json:"id" binding:"required,numeric,min=16,max=16,len=16"`
}

type ValidateUserDataErrorViewResponse struct {
	Id     string   `json:"id"`
	Errors []string `json:"errors"`
}
