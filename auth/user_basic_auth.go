package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UseBasicAuth(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "successfully used endpoint with basic auth",
	})
}
