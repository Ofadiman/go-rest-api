package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/common"
	"net/http"
	"os"
)

func UseCookie(context *gin.Context) {
	cookie, err := context.Cookie("custom_name")
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	fmt.Printf("cookie value is %v", cookie)

	block, cipherErr := aes.NewCipher([]byte(os.Getenv("COOKIE_SECRET")))
	if cipherErr != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, common.ApplicationError{
			Message: cipherErr.Error(),
		})
		return
	}

	cipherText, decodeErr := base64.StdEncoding.DecodeString(cookie)
	if decodeErr != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, common.ApplicationError{
			Message: decodeErr.Error(),
		})
		return
	}

	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	context.JSON(http.StatusOK, gin.H{
		"Message": fmt.Sprintf("successfully used endpoint requiring cookie with \"%v\" cookie value", string(plainText)),
	})
}
