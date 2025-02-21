package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", DbHost, DbUsername, DbPassword, DbDatabase, DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Warning("Failed to create connection")
	}

	if db != nil {
		log.WithFields(log.Fields{
			"DB": "OK",
		}).Info("Success create connection")
	}

	return db
}
