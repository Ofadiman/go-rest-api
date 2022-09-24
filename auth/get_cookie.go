package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/common"
	"net/http"
)

const cookieValue = "super secret cookie"

type GetCookieRequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func GetCookie(context *gin.Context) {
	body := GetCookieRequestBody{}
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

	block, err := aes.NewCipher([]byte("b262e3681420222898e616c9"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	plainText := []byte(cookieValue)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	context.SetCookie("custom_name", base64.StdEncoding.EncodeToString(cipherText), 3600, "/", "localhost", false, true)
}
