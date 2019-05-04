package validation

import (
	"fdnBaseAPI/users"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
}

type UserValidation struct {
	CreateUser CreateUserRequest
	Users      users.Users
}

func (self UserValidation) Bind(c *gin.Context) error {
	err := c.Bind(c, self)
}
