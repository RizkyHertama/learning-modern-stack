package database

import (
	"user-service/env"
	"user-service/src/common/log"
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreSQL() *gorm.DB {
	gormDB, err := gorm.Open(postgres.Open(env.Conf.Postgres.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	return gormDB
}

func ClosePostgreSQL(gormDB *gorm.DB) {
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	if err := sqlDB.Close(); err != nil {
		return err
	}

}
