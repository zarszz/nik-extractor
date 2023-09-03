package view

import "github.com/gin-gonic/gin"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func MakeResponse(c *gin.Context, status int, message string, data interface{}) {
	if status < 400 {
		c.JSON(status, Response{
			Status:  status,
			Message: message,
			Data:    data,
		})
	} else {
		c.AbortWithStatusJSON(status, Response{
			Status:  status,
			Message: message,
			Data:    data,
		})
	}

}
