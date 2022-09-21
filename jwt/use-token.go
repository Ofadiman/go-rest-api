package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ofadiman/go-server/common"
	"net/http"
	"os"
	"strings"
)

func UseToken(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	if tokenString == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, common.ApplicationError{
			Message: "jwt token is missing",
		})
		return
	}

	jwtClaims := jwt.RegisteredClaims{}
	decodedToken, err := jwt.ParseWithClaims(tokenString, &jwtClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token %v", token.Header["alg"])

		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		fmt.Println("aborting with parsing error")
		context.AbortWithStatusJSON(http.StatusUnauthorized, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	if decodedToken.Valid {
		context.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("successfully used token as %v", jwtClaims.Subject),
		})
		return
	} else {
		context.AbortWithStatusJSON(http.StatusUnauthorized, common.ApplicationError{
			Message: "unauthorized",
		})
		return
	}
}
