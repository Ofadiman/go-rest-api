package users

import (
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/common"
	"github.com/ofadiman/go-server/database"
	"net/http"
	"strconv"
)

type DeleteUserByIdRequestUri struct {
	UserID uint64 `uri:"userId" binding:"required"`
}

func DeleteUserById(context *gin.Context) {
	uri := DeleteUserByIdRequestUri{}
	if err := context.ShouldBindUri(&uri); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	user := database.User{}
	getUserByIdQueryResult := database.Gorm.First(&user, "id = ?", uri.UserID)
	if getUserByIdQueryResult.RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, common.ApplicationError{
			Message: "user with id " + strconv.Itoa(int(uri.UserID)) + " not found",
		})
		return
	}

	database.Gorm.Unscoped().Delete(&user)

	context.Status(http.StatusNoContent)
}
