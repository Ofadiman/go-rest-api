package profiles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/src/common"
	"github.com/ofadiman/go-server/src/database"
	"gopkg.in/guregu/null.v4"
	"net/http"
)

type CreateProfileRequestBody struct {
	UserID         uint64  `json:"userId" binding:"required"`
	Picture        *string `json:"picture" binding:"url"`
	FavoriteAnimal *string `json:"favoriteAnimal"`
	FavoriteColor  *string `json:"favoriteColor"`
	FavoriteQuote  *string `json:"favoriteQuote"`
	Gender         *string `json:"gender"`
	JobTitle       *string `json:"jobTitle"`
}

func CreateProfile(context *gin.Context) {
	body := CreateProfileRequestBody{}
	if err := context.ShouldBindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
	}

	user := database.User{}
	findUserByIdQueryResult := database.Gorm.First(&user, "id = ?", body.UserID)
	if findUserByIdQueryResult.RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, common.ApplicationError{
			Message: fmt.Sprintf("user with id %v not found", body.UserID),
		})
		return
	}

	existingProfile := database.Profile{}
	findProfileByUserIdQueryResult := database.Gorm.First(&existingProfile, "user_id = ?", body.UserID)
	if findProfileByUserIdQueryResult.RowsAffected == 1 {
		context.AbortWithStatusJSON(http.StatusConflict, common.ApplicationError{
			Message: fmt.Sprintf("profile for user with id %v already exists", body.UserID),
		})
		return
	}

	newProfile := database.Profile{
		UserID:         body.UserID,
		Picture:        null.StringFromPtr(body.Picture),
		FavoriteAnimal: null.StringFromPtr(body.FavoriteAnimal),
		FavoriteColor:  null.StringFromPtr(body.FavoriteColor),
		FavoriteQuote:  null.StringFromPtr(body.FavoriteQuote),
		Gender:         null.StringFromPtr(body.Gender),
		JobTitle:       null.StringFromPtr(body.JobTitle),
	}
	database.Gorm.Create(&newProfile)

	context.JSON(http.StatusCreated, &newProfile)
}
