package companies

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ofadiman/go-server/common"
	"github.com/ofadiman/go-server/database"
	"net/http"
)

type CreateCompanyRequestBody struct {
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

func CreateCompany(context *gin.Context) {
	body := CreateCompanyRequestBody{}
	if err := context.ShouldBindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, common.ApplicationError{
			Message: err.Error(),
		})
		return
	}

	owner := database.User{}
	getUserByIdQueryResult := database.Gorm.First(&owner, "id = ?", body.OwnerID)
	if getUserByIdQueryResult.RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, common.ApplicationError{
			Message: fmt.Sprintf("user with id %v not found", body.OwnerID),
		})
		return
	}

	company := database.Company{
		Name:           body.Name,
		CatchPhrase:    body.CatchPhrase,
		BuildingNumber: body.BuildingNumber,
		City:           body.City,
		Street:         body.Street,
		Country:        body.Country,
		CountryCode:    body.CountryCode,
		TimeZone:       body.TimeZone,
		ZipCode:        body.ZipCode,
		OwnerID:        body.OwnerID,
	}
	database.Gorm.Create(&company)

	context.JSON(http.StatusCreated, &company)
}
