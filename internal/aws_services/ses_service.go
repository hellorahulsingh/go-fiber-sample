package aws_services

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/ses"
)

type SesService struct {
	BaseAWSService
}

func NewSesService() *SesService {
	baseService := NewBaseAWSService("ses")
	return &SesService{
		BaseAWSService: *baseService,
	}
}

func (s *SesService) SendEmail(recipient string, subject string, body string) (*ses.SendEmailOutput, error) {
	svc := ses.New(session.New())

	emailInput := &ses.SendEmailInput{
		Source: aws.String(os.Getenv("AWS_SES_SENDER_EMAIL")),
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Message: &ses.Message{
			Subject: &ses.Content{
				Data: aws.String(subject),
			},
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(body),
				},
			},
		},
	}

	result, err := svc.SendEmail(emailInput)
	if err != nil {
		log.Printf("Failed to send email: %s\n", err.Error())
		return nil, err
	}

	return result, nil
}
