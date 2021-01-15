package mailjetmail

import (
	"context"
	"emailservice/models"

	"github.com/mailjet/mailjet-apiv3-go"
)

//SendEmailWithMailjet is used to send mail using Mailjet
func SendEmailWithMailjet(ctx context.Context, req models.From) (*mailjet.ResultsV31, error) {

	reqBody := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: req.FromEmail,
				Name:  req.FromName,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: req.ToEmail,
					Name:  req.ToName,
				},
			},
			Cc: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: req.CcEmail,
					Name:  req.CcName,
				},
			},
			Bcc: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: req.BccEmail,
					Name:  req.BccName,
				},
			},
			Subject:  req.Subject,
			TextPart: req.TextPart,
		},
	}

	mailjetClient := mailjet.NewMailjetClient("API_PUBLIC", "API_PRIVATE")

	messages := mailjet.MessagesV31{Info: reqBody}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		return &mailjet.ResultsV31{}, err
	}

	return res, nil

}
