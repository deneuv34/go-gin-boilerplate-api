package main

import (
	"fdnBaseAPI/commons/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// connect to DB
	db := config.InitDB()
	defer db.Close()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	err := r.Run()
	if err != nil {
		log.Println("Cannot run app: ", err)
	}
}
