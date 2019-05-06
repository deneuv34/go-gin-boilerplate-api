package users

import (
	"fdnBaseAPI/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {

}

func (handler UserHandler) CreateUserData(c *gin.Context) {
	var request validation.CreateUserRequest
	var userService UserService
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "bad request",
		})
		c.Abort()
		return
	}
	data, err := userService.CreateNewData(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "server error",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data": data,
	})
}
