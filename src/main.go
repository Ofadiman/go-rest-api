package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ofadiman/go-server/src/database"
	"github.com/ofadiman/go-server/src/users"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
