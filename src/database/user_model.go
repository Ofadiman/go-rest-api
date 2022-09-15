package database

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`

	FirstName   string         `gorm:"not null" json:"firstName"`         // https://fakerjs.dev/api/name.html#firstname
	LastName    string         `gorm:"not null" json:"lastName"`          // https://fakerjs.dev/api/name.html#lastname
	MiddleName  sql.NullString `json:"middleName"`                        // https://fakerjs.dev/api/name.html#middlename
	Email       string         `gorm:"uniqueIndex;not null" json:"email"` // https://fakerjs.dev/api/internet.html#email
	PhoneNumber sql.NullString `gorm:"uniqueIndex" json:"phoneNumber"`    // https://fakerjs.dev/api/phone.html#number

	Profile   Profile
	Posts     []Post
	Companies []*Company `gorm:"many2many:users_companies"`
}
