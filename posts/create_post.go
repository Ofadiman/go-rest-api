package posts

import (
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/common"
	database2 "github.com/ofadiman/go-server/database"
	"net/http"
)

type CreatePostRequestBody struct {
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	TimeToRead uint64 `json:"timeToRead" binding:"required"`
	UserID     uint64 `json:"userId" binding:"required"`
}

func CreatePost(context *gin.Context) {
	body := CreatePostRequestBody{}
	if err := context.ShouldBindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	post := database2.Post{
		Title:      body.Title,
		Content:    body.Content,
		TimeToRead: body.TimeToRead,
		UserID:     body.UserID,
	}
	database2.Gorm.Create(&post)

	context.JSON(http.StatusCreated, &post)
}
