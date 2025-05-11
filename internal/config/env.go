package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	MongoURI          string
	MongoDatabase     string
	AWSRegion         string
	AWSAccountID      string
	AWSPrefix         string
	AWSS3BucketName   string
	AWSSenderEmail    string
}

var AppConfig EnvConfig

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	AppConfig = EnvConfig{
		MongoURI:        GetEnv("MONGODB_URI", "mongodb://localhost:27017"),
		MongoDatabase:   GetEnv("MONGODB_NAME", "go-app"),
		AWSRegion:       GetEnv("AWS_REGION", "ap-south-1"),
		AWSAccountID:    GetEnv("AWS_ACCOUNT_ID", ""),
		AWSPrefix:       GetEnv("APP_NAME", "PawStoriesServer") + "-" + GetEnv("ENV", "development"),
		AWSS3BucketName: GetEnv("AWS_S3_BUCKET_NAME", "pawstories-resources"),
		AWSSenderEmail:  GetEnv("SES_SENDER_EMAIL", "support@neurally.in"),
	}
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
