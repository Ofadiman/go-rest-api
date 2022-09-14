package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Gorm *gorm.DB

func Connect() {
	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN: "host=golang_database user=admin password=password dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Warsaw",
			}),
		&gorm.Config{},
	)

	if err != nil {
		panic("failed to connect database")
	}

	Gorm = db
}
