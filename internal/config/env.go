package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	MongoURI        string
	MongoDatabase   string
	AWSRegion       string
	AWSAccountID    string
	AWSPrefix       string
	AWSS3BucketName string
	AWSSenderEmail  string
}

var AppConfig EnvConfig

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	AppConfig = EnvConfig{
		MongoURI:        getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		MongoDatabase:   getEnv("MONGODB_NAME", "go-app"),
		AWSRegion:       getEnv("AWS_REGION", "ap-south-1"),
		AWSAccountID:    getEnv("AWS_ACCOUNT_ID", ""),
		AWSPrefix:       getEnv("APP_NAME", "PawStoriesServer") + "-" + getEnv("ENV", "development"),
		AWSS3BucketName: getEnv("AWS_S3_BUCKET_NAME", "pawstories-resources"),
		AWSSenderEmail:  getEnv("SES_SENDER_EMAIL", "support@neurally.in"),
	}
}
