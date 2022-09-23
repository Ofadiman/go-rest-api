package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/auth"
	"github.com/ofadiman/go-server/companies"
	"github.com/ofadiman/go-server/database"
	"github.com/ofadiman/go-server/docs"
	"github.com/ofadiman/go-server/posts"
	"github.com/ofadiman/go-server/profiles"
	"github.com/ofadiman/go-server/users"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type PingResponseBody struct {
	Message string `json:"message"`
}

// Ping godoc
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

	r.POST("/auth/jwt", auth.GetToken)
	r.GET("/auth/jwt", auth.UseToken)
	r.GET("/auth/basic", gin.BasicAuth(gin.Accounts{"user": "password"}), auth.UseBasicAuth)

	r.Run()
}
