package posts

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/common"
	database2 "github.com/ofadiman/go-server/database"
	"net/http"
)

type DeletePostByIdRequestUri struct {
	PostID uint64 `uri:"postId" binding:"required"`
}

func DeletePostById(context *gin.Context) {
	uri := DeletePostByIdRequestUri{}
	if err := context.ShouldBindUri(&uri); err != nil {
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

	database2.Gorm.Unscoped().Delete(&post)

	context.Status(http.StatusNoContent)
}
