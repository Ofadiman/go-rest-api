package database

import (
	"errors"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
	"time"
)

type Gender struct {
	value string
}

func (g *Gender) String() string {
	return g.value
}

func (g *Gender) FromString(value string) (*Gender, error) {
	switch value {
	case "male":
		return &Male, nil
	case "female":
		return &Female, nil
	case "non_binary":
		return &NonBinary, nil
	default:
		return nil, errors.New("unsupported enum type " + value)
	}
}

var Male = Gender{value: "male"}
var Female = Gender{value: "female"}
var NonBinary = Gender{value: "non_binary"}

type User struct {
	ID        uint64         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`

	FirstName   string      `gorm:"not null" json:"firstName"`         // https://fakerjs.dev/api/name.html#firstname
	LastName    string      `gorm:"not null" json:"lastName"`          // https://fakerjs.dev/api/name.html#lastname
	MiddleName  null.String `json:"middleName"`                        // https://fakerjs.dev/api/name.html#middlename
	Email       string      `gorm:"uniqueIndex;not null" json:"email"` // https://fakerjs.dev/api/internet.html#email
	PhoneNumber null.String `gorm:"uniqueIndex" json:"phoneNumber"`    // https://fakerjs.dev/api/phone.html#number

	Profile   *Profile   `json:"profile,omitempty"`
	Posts     []*Post    `json:"posts,omitempty"`
	Companies []*Company `gorm:"many2many:users_companies" json:"companies,omitempty"`
}
