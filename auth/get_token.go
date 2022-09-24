package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ofadiman/go-server/common"
	"net/http"
	"os"
	"time"
)

type GetTokenRequestBodyDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GetToken(context *gin.Context) {
	body := GetTokenRequestBodyDto{}
	if err := context.ShouldBindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	if body.Username != "tester" || body.Password != "asdf1234" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, common.ApplicationError{
			Message: "Invalid credentials",
		})
		return
	}

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 48)),
		IssuedAt:  jwt.NewNumericDate(time.Now()), Subject: "tester",
		Issuer: "go-server",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, gin.H{
		"token": signedToken,
	})
}
