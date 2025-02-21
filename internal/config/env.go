package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

var DbHost string
var DbPort string
var DbUsername string
var DbPassword string
var DbDatabase string

func Initialize() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.WithFields(logrus.Fields{
			"env": err.Error(),
		}).Warning("Failed load .env file")
	}

	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUsername = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbDatabase = os.Getenv("DB_DATABASE")
}
