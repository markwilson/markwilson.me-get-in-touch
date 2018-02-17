package mail

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"log"

	"github.com/markwilson/markwilson.me-get-in-touch/api"
)

const (
	Recipient = "Mark Wilson <hello@markwilson.me>"
	Subject   = "Website Contact Submission"
	Sender    = "Contact Form <hello@markwilson.me>"
	CharSet   = "UTF-8"
	AWSRegion = "eu-west-1"
)

func SendEmail(request api.Request) error {
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(AWSRegion),
	})

	sesSession := ses.New(awsSession)

	var replyTo string
	if len(request.Name) < 1 {
		replyTo = fmt.Sprintf("%s <%s>", request.Email, request.Email)
	} else {
		replyTo = fmt.Sprintf("%s <%s>", request.Name, request.Email)
	}

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(fmt.Sprintf("Message: %s", request.Message)),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject),
			},
		},
		Source:           aws.String(Sender),
		ReplyToAddresses: aws.StringSlice([]string{replyTo}),
	}

	_, err = sesSession.SendEmail(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				log.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				log.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				log.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				log.Println(aerr.Error())
			}
		} else {
			log.Println(err.Error())
		}

		return err
	}

	return nil
}
