package config

import (
	"log"
	"os"

	"github.com/MKA-Nigeria/mkanmedia-go/logger"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var Env env

const defaultMongoUrl = "mongodb://localhost:27017"

func init() {
	// load env before reading from env
	err := godotenv.Load()

	// load end config
	e := loadConfig()

	if err != nil && e.ENV != "prod" {
		log.Fatalf("err loading: %v", err)
	}

	// validate env being loaded
	if err := e.Validate(); err != nil {
		logger.DefaultLogger.WithFields(logrus.Fields{"type": "env_error", "stack": err}).Error("Invalid environment variables")
	}

	Env = e
}

func loadConfig() env {
	port := os.Getenv("PORT")

	mongoUrl := os.Getenv("MONGO_URL")
	if mongoUrl == "" {
		mongoUrl = defaultMongoUrl
	}

	redisUrl := os.Getenv("REDIS_URL")

	return env{
		ENV:                       os.Getenv("ENV"),
		PORT:                      port,
		MONGO_URL:                 mongoUrl,
		REDIS_URL:                 redisUrl,
		SOUND_CLOUD_CLIENT_ID:     os.Getenv("SOUND_CLOUD_CLIENT_ID"),
		SOUND_CLOUD_CLIENT_SECRET: os.Getenv("SOUND_CLOUD_CLIENT_SECRET"),
	}
}
