package profiles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/src/common"
	"github.com/ofadiman/go-server/src/database"
	"net/http"
)

type GetProfileByIdRequestUri struct {
	ProfileID uint64 `uri:"profileId" binding:"required"`
}

func GetProfileById(context *gin.Context) {
	uri := GetProfileByIdRequestUri{}
	if err := context.ShouldBindUri(&uri); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	profile := database.Profile{}
	getProfileByIdQueryResult := database.Gorm.First(&profile, "id = ?", uri.ProfileID)
	if getProfileByIdQueryResult.RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, common.ApplicationError{
			Message: fmt.Sprintf("profile with id %v not found", uri.ProfileID),
		})
		return
	}

	context.JSON(http.StatusOK, profile)
}
