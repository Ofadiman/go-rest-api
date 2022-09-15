package database

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	ID        uint64         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`

	Picture        sql.NullString `json:"Picture"`        // https://fakerjs.dev/api/image.html#avatar
	FavoriteAnimal sql.NullString `json:"favoriteAnimal"` // https://fakerjs.dev/api/animal.html#type
	FavoriteColor  sql.NullString `json:"favoriteColor"`  // https://fakerjs.dev/api/color.html#human
	FavoriteQuote  sql.NullString `json:"favoriteQuote"`  // https://fakerjs.dev/api/hacker.html#verb
	Gender         sql.NullString `json:"gender"`         // https://fakerjs.dev/api/name.html#gender
	JobTitle       sql.NullString `json:"jobTitle"`       // https://fakerjs.dev/api/name.html#jobtitle

	UserID uint `gorm:"not null;uniqueIndex" json:"userId"`
}
