package aws_services

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ses"
)

type BaseAWSService struct {
	client interface{}
}

func NewBaseAWSService(serviceName string) *BaseAWSService {
	// You can add code to handle AWS region and credentials here
	awsRegion := os.Getenv("AWS_REGION")
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	// Initialize the AWS session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
	})
	if err != nil {
		log.Printf("Error initializing AWS session: %s\n", err.Error())
		return nil
	}

	// Initialize service client
	var client interface{}
	switch serviceName {
	case "s3":
		client = s3.New(sess)
	case "ses":
		client = ses.New(sess)
	default:
		log.Printf("Unknown AWS service: %s\n", serviceName)
	}

	return &BaseAWSService{client: client}
}
