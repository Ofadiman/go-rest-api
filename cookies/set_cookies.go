package cookies

import (
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/common"
	"net/http"
)

type SetCookiesRequestBody struct {
	Value string `json:"value" binding:"required"`
}

func SetCookies(context *gin.Context) {
	body := SetCookiesRequestBody{}
	if err := context.ShouldBindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	context.SetCookie("custom_name", body.Value, 3600, "/", "localhost", false, true)
}
