package users

import (
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/common"
	database2 "github.com/ofadiman/go-server/database"
	"gopkg.in/guregu/null.v4"
	"net/http"
)

type CreateUserRequestBody struct {
	FirstName   string  `json:"firstName" binding:"required,max=200"`
	LastName    string  `json:"lastName" binding:"required,max=200"`
	Email       string  `json:"email" binding:"required,email,max=200"`
	MiddleName  *string `json:"middleName" binding:"omitempty,max=200"`
	PhoneNumber *string `json:"phoneNumber" binding:"omitempty,max=200"`
}

func CreateUser(context *gin.Context) {
	body := CreateUserRequestBody{}
	if err := context.ShouldBindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	userByEmail := database2.User{}
	getUserByEmailQueryResult := database2.Gorm.First(&userByEmail, "email = ?", body.Email)
	if getUserByEmailQueryResult.RowsAffected != 0 {
		context.AbortWithStatusJSON(http.StatusConflict, common.ApplicationError{
			Message: "user with email " + body.Email + " already exists",
		})
		return
	}

	if body.PhoneNumber != nil {
		userByPhoneNumber := database2.User{}
		getUserByPhoneNumberQueryResult := database2.Gorm.First(&userByPhoneNumber, "phone_number = ?", body.PhoneNumber)
		if getUserByPhoneNumberQueryResult.RowsAffected != 0 {
			context.AbortWithStatusJSON(http.StatusConflict, common.ApplicationError{
				Message: "user with phone number " + *body.PhoneNumber + " already exists",
			})
			return
		}
	}

	newUser := database2.User{
		FirstName:   body.FirstName,
		MiddleName:  null.StringFromPtr(body.MiddleName),
		LastName:    body.LastName,
		Email:       body.Email,
		PhoneNumber: null.StringFromPtr(body.PhoneNumber),
	}
	database2.Gorm.Create(&newUser)

	context.JSON(http.StatusCreated, &newUser)
}
