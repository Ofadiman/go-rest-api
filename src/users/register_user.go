package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/src/database"
	"net/http"
)

type RegisterUserRequestBody struct {
	FirstName string `json:"firstName" binding:"required,min=4,max=200"`
	LastName  string `json:"LastName" binding:"required,min=4,max=200"`
	Email     string `json:"email" binding:"required,email,min=4,max=200"`
}

type RegisterUserResponseBody struct {
	ID uint `json:"id"`
}

func RegisterUser(context *gin.Context) {
	body := RegisterUserRequestBody{}
	if err := context.ShouldBindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	existingUser := User{}
	result := database.Gorm.First(&existingUser, "email = ?", body.Email)

	if result.RowsAffected != 0 {
		context.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"err": "user already exists",
		})
		fmt.Println("User found")
		return
	}

	newUser := User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
	}
	database.Gorm.Create(&newUser)

	responseBody := RegisterUserResponseBody{ID: newUser.ID}
	context.JSON(http.StatusOK, &responseBody)
}
