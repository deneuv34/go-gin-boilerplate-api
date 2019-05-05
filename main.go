package main

import (
	"fdnBaseAPI/commons/config"
	"fdnBaseAPI/users"
	"github.com/gin-gonic/gin"
	"log"
)

var user users.UserRoute

func main() {
	// connect to DB
	db := config.InitDB()
	db.AutoMigrate(&users.Users{})
	defer db.Close()

	r := gin.Default()
	user.Route(r.Group("/user"))

	err := r.Run()
	if err != nil {
		log.Println("Cannot run app: ", err)
	}
}
