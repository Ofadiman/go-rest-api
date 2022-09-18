package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ofadiman/go-server/src/companies"
	"github.com/ofadiman/go-server/src/database"
	"github.com/ofadiman/go-server/src/posts"
	"github.com/ofadiman/go-server/src/profiles"
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
	database.Gorm.AutoMigrate(&database.User{})
	database.Gorm.AutoMigrate(&database.Company{})
	database.Gorm.AutoMigrate(&database.Profile{})
	database.Gorm.AutoMigrate(&database.Post{})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/users", users.CreateUser)
	r.GET("/users/:userId", users.GetUserById)
	r.PUT("/users/:userId", users.ReplaceUserById)
	r.DELETE("/users/:userId", users.DeleteUserById)
	r.GET("/users", users.PaginateUsers)

	r.POST("/profiles", profiles.CreateProfile)
	r.GET("/profiles/:profileId", profiles.GetProfileById)
	r.PUT("/profiles/:profileId", profiles.ReplaceProfileById)
	r.DELETE("/profiles/:profileId", profiles.DeleteProfileById)

	r.POST("/posts", posts.CreatePost)
	r.GET("/posts/:postId", posts.GetPostById)
	r.PUT("/posts/:postId", posts.ReplacePostById)
	r.DELETE("/posts/:postId", posts.DeletePostById)

	r.POST("/companies", companies.CreateCompany)
	r.GET("/companies/:companyId", companies.GetCompanyById)
	r.PUT("/companies/:companyId", companies.ReplaceCompanyById)
	r.DELETE("/companies/:companyId", companies.DeleteCompanyById)

	r.Run()
}
