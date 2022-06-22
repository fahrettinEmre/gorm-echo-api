package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if !DB.Migrator().HasTable(&Todo{}) {
		DB.AutoMigrate(&Todo{})
	}
}
