package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() error {
	var err error
	var logLevel logger.LogLevel

	switch Config.LogLevel {
	case "Info":
		logLevel = logger.Info
	case "Warn":
		logLevel = logger.Warn
	default:
		logLevel = logger.Silent
	}

	db, err = gorm.Open(postgres.Open(Config.DatabaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	return err
}

func GetDB() *gorm.DB {
	return db
}
