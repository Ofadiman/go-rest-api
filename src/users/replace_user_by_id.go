package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/src/common"
	"github.com/ofadiman/go-server/src/database"
	"gopkg.in/guregu/null.v4"
	"net/http"
)

type ReplaceUserByIdRequestBody struct {
	FirstName   string  `json:"firstName" binding:"required,max=200"`
	LastName    string  `json:"lastName" binding:"required,max=200"`
	Email       string  `json:"email" binding:"required,email,max=200"`
	MiddleName  *string `json:"middleName" binding:"omitempty,max=200"`
	PhoneNumber *string `json:"phoneNumber" binding:"omitempty,max=200"`
}

type ReplaceUserByIdRequestUri struct {
	UserID uint64 `uri:"userId" binding:"required"`
}

func ReplaceUserById(context *gin.Context) {
	uri := ReplaceUserByIdRequestUri{}
	if err := context.ShouldBindUri(&uri); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	body := ReplaceUserByIdRequestBody{}
	if err := context.ShouldBindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	userByEmail := database.User{}
	getUserByEmailQueryResult := database.Gorm.First(&userByEmail, "email = ?", body.Email)
	if getUserByEmailQueryResult.RowsAffected != 0 {
		context.AbortWithStatusJSON(http.StatusConflict, common.ApplicationError{
			Message: fmt.Sprintf("user with email %v already exists", body.Email),
		})
		return
	}

	if body.PhoneNumber != nil {
		userByPhoneNumber := database.User{}
		getUserByPhoneNumberQueryResult := database.Gorm.First(&userByPhoneNumber, "phone_number = ?", body.PhoneNumber)
		if getUserByPhoneNumberQueryResult.RowsAffected != 0 {
			context.AbortWithStatusJSON(http.StatusConflict, common.ApplicationError{
				Message: fmt.Sprintf("user with phone number %v already exists", body.PhoneNumber),
			})
			return
		}
	}

	user := database.User{}
	getUserByIdQueryResult := database.Gorm.First(&user, "id = ?", uri.UserID)
	if getUserByIdQueryResult.RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, common.ApplicationError{
			Message: fmt.Sprintf("user with id %v not found", uri.UserID),
		})
		return
	}

	database.Gorm.Model(&user).
		Where(&database.User{ID: uri.UserID}).
		Updates(map[string]interface{}{
			"FirstName":   body.FirstName,
			"LastName":    body.LastName,
			"Email":       body.Email,
			"MiddleName":  null.StringFromPtr(body.MiddleName),
			"PhoneNumber": null.StringFromPtr(body.PhoneNumber),
		})

	context.JSON(http.StatusOK, user)
}
