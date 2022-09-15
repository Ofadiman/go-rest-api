package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/src/database"
	"gopkg.in/gomail.v2"
	"net/http"
	"os"
	"strconv"
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

	existingUser := database.User{}
	result := database.Gorm.First(&existingUser, "email = ?", body.Email)

	if result.RowsAffected != 0 {
		context.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"err": "user already exists",
		})
		fmt.Println("User found")
		return
	}

	newUser := database.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
	}
	database.Gorm.Create(&newUser)

	message := gomail.NewMessage()
	message.SetHeader("From", os.Getenv("SMTP_FROM_EMAIL"))
	message.SetHeader("To", body.Email)
	message.SetHeader("Subject", "Go-server account activation")
	message.SetBody("text/html", "<a href=\""+os.Getenv("SERVER_URL")+"/users/activate/"+strconv.FormatUint(uint64(newUser.ID), 10)+"\">Click here to activate your account.</a>")

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("SMTP_FROM"), os.Getenv("SMTP_PASSWORD"))

	if err := d.DialAndSend(message); err != nil {
		panic(err)
	}

	responseBody := RegisterUserResponseBody{ID: newUser.ID}
	context.JSON(http.StatusOK, &responseBody)
}
