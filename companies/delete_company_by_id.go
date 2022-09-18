package companies

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/common"
	database2 "github.com/ofadiman/go-server/database"
	"net/http"
)

type DeleteCompanyByIdRequestUri struct {
	CompanyID uint64 `uri:"companyId" binding:"required"`
}

func DeleteCompanyById(context *gin.Context) {
	uri := DeleteCompanyByIdRequestUri{}
	if err := context.ShouldBindUri(&uri); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	company := database2.Company{}
	getCompanyByIdQueryResult := database2.Gorm.First(&company, "id = ?", uri.CompanyID)
	if getCompanyByIdQueryResult.RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, common.ApplicationError{
			Message: fmt.Sprintf("company with id %v not found", uri.CompanyID),
		})
		return
	}

	database2.Gorm.Unscoped().Delete(&company)

	context.Status(http.StatusNoContent)
}
