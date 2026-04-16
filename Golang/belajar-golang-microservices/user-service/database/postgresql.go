package database

import (
	"user-service/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreSQL() *gorm.DB{
	gormDB, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=users port=5432"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	return gormDB
}
