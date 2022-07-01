package config

import (
	"fmt"
	"os"

	"github.com/MKA-Nigeria/mkanmedia-go/logger"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var Env env

const defaultPort = "8082"
const defaultMongoUrl = "mongodb://localhost:27017"

func init() {
	// load env before reading from env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Couldn't load .env file, are you in Prod?")
	}

	// load end config
	e := loadConfig()

	// validate env being loaded
	if err := e.Validate(); err != nil {
		logger.DefaultLogger.WithFields(logrus.Fields{"type": "env_error", "stack": err}).Error("Invalid environment variables")
	}

	Env = e
}

func loadConfig() env {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mongoUrl := os.Getenv("MONGO_URL")
	if mongoUrl == "" {
		mongoUrl = defaultMongoUrl
	}

	return env{
		ENV:                       os.Getenv("ENV"),
		PORT:                      port,
		MONGO_URL:                 mongoUrl,
		SOUND_CLOUD_CLIENT_ID:     os.Getenv("SOUND_CLOUD_CLIENT_ID"),
		SOUND_CLOUD_CLIENT_SECRET: os.Getenv("SOUND_CLOUD_CLIENT_SECRET"),
	}
}
