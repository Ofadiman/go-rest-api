package main

import (
	"github.com/gin-gonic/gin"
	companies2 "github.com/ofadiman/go-server/companies"
	database2 "github.com/ofadiman/go-server/database"
	"github.com/ofadiman/go-server/docs"
	posts2 "github.com/ofadiman/go-server/posts"
	profiles2 "github.com/ofadiman/go-server/profiles"
	users2 "github.com/ofadiman/go-server/users"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	database2.Connect()
	database2.Gorm.AutoMigrate(&database2.User{})
	database2.Gorm.AutoMigrate(&database2.Company{})
	database2.Gorm.AutoMigrate(&database2.Profile{})
	database2.Gorm.AutoMigrate(&database2.Post{})

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", Ping)

	r.POST("/users", users2.CreateUser)
	r.GET("/users/:userId", users2.GetUserById)
	r.PUT("/users/:userId", users2.ReplaceUserById)
	r.DELETE("/users/:userId", users2.DeleteUserById)
	r.GET("/users", users2.PaginateUsers)

	r.POST("/profiles", profiles2.CreateProfile)
	r.GET("/profiles/:profileId", profiles2.GetProfileById)
	r.PUT("/profiles/:profileId", profiles2.ReplaceProfileById)
	r.DELETE("/profiles/:profileId", profiles2.DeleteProfileById)

	r.POST("/posts", posts2.CreatePost)
	r.GET("/posts/:postId", posts2.GetPostById)
	r.PUT("/posts/:postId", posts2.ReplacePostById)
	r.DELETE("/posts/:postId", posts2.DeletePostById)

	r.POST("/companies", companies2.CreateCompany)
	r.GET("/companies/:companyId", companies2.GetCompanyById)
	r.PUT("/companies/:companyId", companies2.ReplaceCompanyById)
	r.DELETE("/companies/:companyId", companies2.DeleteCompanyById)

	r.Run()
}
