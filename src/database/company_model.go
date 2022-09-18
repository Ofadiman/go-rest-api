package database

import (
	"gorm.io/gorm"
	"time"
)

type Company struct {
	ID        uint64         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`

	Name           string `gorm:"not null;uniqueIndex" json:"name"` // https://fakerjs.dev/api/company.html#name
	CatchPhrase    string `gorm:"not null" json:"catchPhrase"`      // https://fakerjs.dev/api/company.html#catchphrase
	BuildingNumber string `gorm:"not null" json:"buildingNumber"`   // https://fakerjs.dev/api/address.html#buildingnumber
	City           string `gorm:"not null" json:"city"`             // https://fakerjs.dev/api/address.html#city
	Street         string `gorm:"not null" json:"street"`           // https://fakerjs.dev/api/address.html#street
	Country        string `gorm:"not null" json:"country"`          // https://fakerjs.dev/api/address.html#country
	CountryCode    string `gorm:"not null" json:"countryCode"`      // https://fakerjs.dev/api/address.html#countrycode
	TimeZone       string `gorm:"not null" json:"timeZone"`         // https://fakerjs.dev/api/address.html#timezone
	ZipCode        string `gorm:"not null" json:"zipCode"`          // https://fakerjs.dev/api/address.html#zipcode

	Employees []*User `gorm:"many2many:users_companies" json:"employees,omitempty"`
	Owner     *User   `gorm:"foreignKey:OwnerID" json:"owner,omitempty"`
	OwnerID   uint64  `gorm:"not null"`
}
