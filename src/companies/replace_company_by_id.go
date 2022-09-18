package companies

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/src/common"
	"github.com/ofadiman/go-server/src/database"
	"net/http"
)

type ReplaceCompanyByIdRequestUri struct {
	CompanyID uint64 `uri:"companyId" binding:"required"`
}

type ReplaceCompanyByIdRequestBody struct {
	Name           string `json:"name" binding:"required"`
	CatchPhrase    string `json:"catchPhrase" binding:"required"`
	BuildingNumber string `json:"buildingNumber" binding:"required"`
	City           string `json:"city" binding:"required"`
	Street         string `json:"street" binding:"required"`
	Country        string `json:"country" binding:"required"`
	CountryCode    string `json:"countryCode" binding:"required"`
	TimeZone       string `json:"timeZone" binding:"required"`
	ZipCode        string `json:"zipCode" binding:"required"`
	OwnerID        uint64 `json:"ownerId" binding:"required"`
}

func ReplaceCompanyById(context *gin.Context) {
	uri := ReplaceCompanyByIdRequestUri{}
	if err := context.ShouldBindUri(&uri); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	body := ReplaceCompanyByIdRequestBody{}
	if err := context.ShouldBindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	user := database.User{}
	getUserByIdQueryResult := database.Gorm.First(&user, "id = ?", body.OwnerID)
	if getUserByIdQueryResult.RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, common.ApplicationError{
			Message: fmt.Sprintf("user with id %v not found", body.OwnerID),
		})
		return
	}

	company := database.Company{}
	getCompanyByIdQueryResult := database.Gorm.First(&company, "id = ?", uri.CompanyID)
	if getCompanyByIdQueryResult.RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, common.ApplicationError{
			Message: fmt.Sprintf("compnay with id %v not found", uri.CompanyID),
		})
		return
	}

	database.Gorm.Model(&company).
		Where(&database.Company{ID: uri.CompanyID}).
		Updates(map[string]interface{}{
			"Name":           body.Name,
			"CatchPhrase":    body.CatchPhrase,
			"BuildingNumber": body.BuildingNumber,
			"City":           body.City,
			"Street":         body.Street,
			"Country":        body.Country,
			"CountryCode":    body.CountryCode,
			"TimeZone":       body.TimeZone,
			"ZipCode":        body.ZipCode,
			"OwnerID":        body.OwnerID,
		})

	context.JSON(http.StatusOK, company)
}
