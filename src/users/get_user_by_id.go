package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/src/database"
	"net/http"
)

type GetUserByIdRequestUri struct {
	ID uint `uri:"userId" json:"id" binding:"required"`
}

func GetUserById(context *gin.Context) {
	uri := GetUserByIdRequestUri{}
	if err := context.ShouldBindUri(&uri); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	fmt.Println(uri.ID)

	existingUser := User{}
	result := database.Gorm.First(&existingUser, "id = ?", uri.ID)

	if result.RowsAffected == 0 {
		context.JSON(http.StatusConflict, gin.H{
			"err": "user not found",
		})
		return
	}

	context.JSON(http.StatusOK, &existingUser)
}
