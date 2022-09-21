package users

import (
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/common"
	"github.com/ofadiman/go-server/database"
	"net/http"
	"strconv"
)

type GetUserByIdRequestUri struct {
	UserID uint64 `uri:"userId" binding:"required"`
}

func GetUserById(context *gin.Context) {
	uri := GetUserByIdRequestUri{}
	if err := context.ShouldBindUri(&uri); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	existingUser := database.User{}
	getUserByIdQueryResult := database.Gorm.First(&existingUser, "id = ?", uri.UserID)
	if getUserByIdQueryResult.RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, common.ApplicationError{
			Message: "user with id " + strconv.Itoa(int(uri.UserID)) + " not found",
		})
		return
	}

	context.JSON(http.StatusOK, &existingUser)
}
