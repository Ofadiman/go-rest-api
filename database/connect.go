package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Gorm *gorm.DB

func Connect() {
	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN: os.Getenv("DATABASE_DSN"),
			}),
		&gorm.Config{},
	)

	if err != nil {
		panic("failed to connect database")
	}

	Gorm = db
}
