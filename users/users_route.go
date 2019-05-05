package users

import "github.com/gin-gonic/gin"

var handler UserHandler

type UserRoute struct {
}

func (r UserRoute) Route(route *gin.RouterGroup) {
	userRepository.Init(Users{})

	route.POST("/", handler.CreateUserData)
}
