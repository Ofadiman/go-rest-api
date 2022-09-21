package posts

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/common"
	"github.com/ofadiman/go-server/database"
	"net/http"
)

type GetPostByIdRequestUri struct {
	PostID uint64 `uri:"postId" binding:"required"`
}

func GetPostById(context *gin.Context) {
	uri := GetPostByIdRequestUri{}
	if err := context.ShouldBindUri(&uri); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	post := database.Post{}
	getPostByIdQueryResult := database.Gorm.First(&post, "id = ?", uri.PostID)
	if getPostByIdQueryResult.RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, common.ApplicationError{
			Message: fmt.Sprintf("post with id %v not found", uri.PostID),
		})
		return
	}

	context.JSON(http.StatusOK, post)
}
