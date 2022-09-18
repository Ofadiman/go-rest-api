package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ofadiman/go-server/src/companies"
	"github.com/ofadiman/go-server/src/database"
	docs "github.com/ofadiman/go-server/src/docs"
	"github.com/ofadiman/go-server/src/posts"
	"github.com/ofadiman/go-server/src/profiles"
	"github.com/ofadiman/go-server/src/users"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)

type PingResponseBody struct {
	Message string `json:"message"`
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description ping server
// @Tags ping
// @Accept json
// @Produce json
// @success 200 {object} PingResponseBody
// @Router /ping [get]
func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, PingResponseBody{Message: "pong"})
}

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

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", Ping)

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
