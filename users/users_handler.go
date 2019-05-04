package users

import (
	"fdnBaseAPI/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUserData(c *gin.Context) {
	err := c.Bind(validation.CreateUserRequest{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "bad request",
		})
		c.Abort()
	}
	data, err := UserService.CreateNewData(validation.CreateUserRequest{})
}
