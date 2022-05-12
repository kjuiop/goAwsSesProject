package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const CHARSET = "UTF-8"

func SendEmail(sess *session.Session, toAddresses []*string, emailText string, sender string, subject string) error {
	sesClient := ses.New(sess)

	_, err := sesClient.SendEmail(&ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: toAddresses,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String(CHARSET),
					Data:    aws.String(emailText),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CHARSET),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
	})

	return err
}

func SendHTMLEmail(sess *session.Session, toAddresses []*string, htmlText string, sender string, subject string) error {
	sesClient := ses.New(sess)

	_, err := sesClient.SendEmail(&ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: toAddresses,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CHARSET),
					Data:    aws.String(htmlText),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CHARSET),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
	})

	return err
}

func main() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("ap-northeast-2"),
		},
	})

	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
		return
	}

	toAddresses := []*string{
		aws.String("arneg0shua@gmail.com"),
	}
	emailText := "Amazing SES Tutorial"
	htmlText := "<html>"
	htmlText += "<head></head>"
	htmlText += "<h1 style='text-align:center'>This is the heading</h1>"
	htmlText += "<p>Hello, world</p>"
	htmlText += "</body>"
	htmlText += "</html>"
	sender := "arneg0shua@gmail.com"
	subject := "Amazing Email Tutorial!"

	err = SendEmail(sess, toAddresses, emailText, sender, subject)
	if err != nil {
		fmt.Printf("Got an error while trying to send email: %v", err)
		return
	}

	fmt.Println("Sent text email successfully")

	err = SendHTMLEmail(sess, toAddresses, htmlText, sender, subject)
	if err != nil {
		fmt.Printf("Got an error while trying to send email: %v", err)
		return
	}

	fmt.Println("Sent html email successfully")

}
