package posts

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/common"
	database2 "github.com/ofadiman/go-server/database"
	"net/http"
)

type ReplacePostByIdRequestUri struct {
	PostID uint64 `uri:"postId" binding:"required"`
}

type ReplacePostByIdRequestBody struct {
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	TimeToRead uint64 `json:"timeToRead" binding:"required"`
	UserID     uint64 `json:"userId" binding:"required"`
}

func ReplacePostById(context *gin.Context) {
	uri := ReplacePostByIdRequestUri{}
	if err := context.ShouldBindUri(&uri); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	body := ReplacePostByIdRequestBody{}
	if err := context.ShouldBindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	post := database2.Post{}
	getPostByIdQueryResult := database2.Gorm.First(&post, "id = ?", uri.PostID)
	if getPostByIdQueryResult.RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, common.ApplicationError{
			Message: fmt.Sprintf("post with id %v not found", uri.PostID),
		})
		return
	}

	user := database2.User{}
	getUserByIdQueryResult := database2.Gorm.First(&user, "id = ?", body.UserID)
	if getUserByIdQueryResult.RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, common.ApplicationError{
			Message: fmt.Sprintf("user with id %v not found", body.UserID),
		})
		return
	}

	database2.Gorm.Model(&post).
		Where(&database2.Post{ID: uri.PostID}).
		Updates(map[string]interface{}{
			"Title":      body.Title,
			"Content":    body.Content,
			"TimeToRead": body.TimeToRead,
			"UserID":     body.UserID,
		})

	context.JSON(http.StatusOK, post)
}
