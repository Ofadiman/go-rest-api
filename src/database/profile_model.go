package database

import (
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	ID        uint64         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`

	Picture        null.String `json:"Picture"`        // https://fakerjs.dev/api/image.html#avatar
	FavoriteAnimal null.String `json:"favoriteAnimal"` // https://fakerjs.dev/api/animal.html#type
	FavoriteColor  null.String `json:"favoriteColor"`  // https://fakerjs.dev/api/color.html#human
	FavoriteQuote  null.String `json:"favoriteQuote"`  // https://fakerjs.dev/api/hacker.html#verb
	Gender         null.String `json:"gender"`         // https://fakerjs.dev/api/name.html#gender
	JobTitle       null.String `json:"jobTitle"`       // https://fakerjs.dev/api/name.html#jobtitle

	UserID uint64 `gorm:"not null;uniqueIndex" json:"userId"`
}
