package cookies

import "github.com/gin-gonic/gin"

func ClearCookies(context *gin.Context) {
	context.SetCookie("custom_name", "", 0, "/", "localhost", false, true)
}
