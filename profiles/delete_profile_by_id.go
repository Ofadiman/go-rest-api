package profiles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/common"
	database2 "github.com/ofadiman/go-server/database"
	"net/http"
)

type DeleteProfileByIdRequestUri struct {
	ProfileID uint64 `uri:"profileId" binding:"required"`
}

func DeleteProfileById(context *gin.Context) {
	uri := DeleteProfileByIdRequestUri{}
	if err := context.ShouldBindUri(&uri); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	profile := database2.Profile{}
	getProfileByIdQueryResult := database2.Gorm.First(&profile, "id = ?", uri.ProfileID)
	if getProfileByIdQueryResult.RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, common.ApplicationError{
			Message: fmt.Sprintf("profile with id %v not found", uri.ProfileID),
		})
		return
	}

	database2.Gorm.Unscoped().Delete(&profile)

	context.Status(http.StatusNoContent)
}
