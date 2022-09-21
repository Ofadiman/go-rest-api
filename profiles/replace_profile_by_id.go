package profiles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/common"
	"github.com/ofadiman/go-server/database"
	"gopkg.in/guregu/null.v4"
	"net/http"
)

type ReplaceProfileByIdRequestUri struct {
	ProfileID uint64 `uri:"profileId" binding:"required"`
}

type ReplaceProfileByIdRequestBody struct {
	UserID         uint64  `json:"userId" binding:"required"`
	Picture        *string `json:"picture" binding:"url"`
	FavoriteAnimal *string `json:"favoriteAnimal"`
	FavoriteColor  *string `json:"favoriteColor"`
	FavoriteQuote  *string `json:"favoriteQuote"`
	Gender         *string `json:"gender"`
	JobTitle       *string `json:"jobTitle"`
}

func ReplaceProfileById(context *gin.Context) {
	uri := ReplaceProfileByIdRequestUri{}
	if err := context.ShouldBindUri(&uri); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	body := ReplaceProfileByIdRequestBody{}
	if err := context.ShouldBindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	user := database.User{}
	getUserByIdQueryResult := database.Gorm.First(&user, "id = ?", body.UserID)
	if getUserByIdQueryResult.RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, common.ApplicationError{
			Message: fmt.Sprintf("user with id %v not found", body.UserID),
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

	database.Gorm.Model(&profile).
		Where(&database.Profile{ID: uri.ProfileID}).
		Updates(map[string]interface{}{
			"UserID":         body.UserID,
			"Picture":        null.StringFromPtr(body.Picture),
			"FavoriteAnimal": null.StringFromPtr(body.FavoriteAnimal),
			"FavoriteColor":  null.StringFromPtr(body.FavoriteColor),
			"FavoriteQuote":  null.StringFromPtr(body.FavoriteQuote),
			"Gender":         null.StringFromPtr(body.Gender),
			"JobTitle":       null.StringFromPtr(body.JobTitle),
		})

	context.JSON(http.StatusOK, profile)
}
