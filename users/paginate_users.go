package users

import (
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/common"
	database2 "github.com/ofadiman/go-server/database"
	"net/http"
)

type PaginateUsersQuery struct {
	Page    uint64 `form:"page" binding:"required"`
	PerPage uint64 `form:"perPage" binding:"required"`
}

func PaginateUsers(context *gin.Context) {
	query := PaginateUsersQuery{}
	if err := context.ShouldBindQuery(&query); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	var users []database2.User
	database2.Gorm.Offset(int((query.Page - 1) * query.PerPage)).Limit(int(query.PerPage)).Find(&users)
	var count int64
	database2.Gorm.Model(&database2.User{}).Count(&count)

	context.JSON(http.StatusOK, gin.H{
		"rows": users,
		"meta": gin.H{
			"page":    query.Page,
			"perPage": query.PerPage,
			"total":   count,
		},
	})
}
