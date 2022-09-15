package database

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        uint64         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`

	Title      string  `gorm:"not null" json:"title"`
	Content    string  `gorm:"not null" json:"content"`
	TimeToRead float64 `gorm:"not null" json:"timeToRead"`

	UserID uint64 `gorm:"index;not null" json:"userID"`
}
