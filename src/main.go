package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/src/database"
	"github.com/ofadiman/go-server/src/users"
	"net/http"
)

func main() {
	database.Connect()
	database.Gorm.AutoMigrate(&users.User{})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/users/register", users.RegisterUser)
	r.GET("/users/:userId", users.GetUserById)

	r.Run()
}
