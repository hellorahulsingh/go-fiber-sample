package aws_services

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Service struct {
	BaseAWSService
	bucketName string
}

func NewS3Service() *S3Service {
	baseService := NewBaseAWSService("s3")
	bucketName := os.Getenv("AWS_S3_BUCKET_NAME")
	return &S3Service{
		BaseAWSService: *baseService,
		bucketName:     bucketName,
	}
}

func (s *S3Service) UploadFile(file *os.File, fileName string) (string, error) {
	uploader := s3manager.NewUploader(session.New())

	// Upload file to S3
	fileKey := fmt.Sprintf("uploads/%s_%s", time.Now().Format("20060102150405"), fileName)
	result, err := uploader.Upload(&s3.UploadInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(fileKey),
		Body:   file,
	})
	if err != nil {
		log.Printf("Failed to upload file to S3: %s\n", err.Error())
		return "", err
	}

	return result.Location, nil
}

func (s *S3Service) GetPresignedURL(fileKey string, expiration int64) (string, error) {
	svc := s3.New(session.New())
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(fileKey),
	})
	url, err := req.Presign(time.Duration(expiration) * time.Second)
	if err != nil {
		log.Printf("Error generating presigned URL: %s\n", err.Error())
		return "", err
	}

	return url, nil
}
