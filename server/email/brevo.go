package email

import (
	"context"
	"fmt"
	"keyvalue-api/constants"

	brevo "github.com/getbrevo/brevo-go/lib"
)

func InitBrevo() *brevo.APIClient {
	var ctx context.Context
	cfg := brevo.NewConfiguration()
	//Configure API key authorization: api-key
	cfg.AddDefaultHeader("api-key", constants.BREVO_API_KEY)

	brevoClient := brevo.NewAPIClient(cfg)
	result, resp, err := brevoClient.AccountApi.GetAccount(ctx)
	if err != nil {
		fmt.Println("Error when calling AccountApi->get_account: ", err.Error())
		return nil
	}
	fmt.Println("GetAccount Object:", result, " GetAccount Response: ", resp)
	return brevoClient
}

type AppEmailer struct {
	BrevoClient *brevo.APIClient
}

func (emailer *AppEmailer) SendLimitReachedEmail(emails []string) (brevo.CreateSmtpEmail, error) {
	ctx := context.Background()
	// Set up the email content
	toEmails := make([]brevo.SendSmtpEmailTo, 0)
	for _, email := range emails {
		toEmails = append(toEmails, brevo.SendSmtpEmailTo{
			Email: email,
			Name:  email,
		})
	}
	emailContent := brevo.SendSmtpEmail{
		Sender: &brevo.SendSmtpEmailSender{
			Name:  "KeyValue API",
			Email: constants.FROM_EMAIL,
		},
		To:          toEmails,
		Subject:     "KeyValue API - Limit Reached",
		HtmlContent: "<html><body><h1>Limit Reached</h1><p>Your limit for the KeyValue API has been reached. Please upgrade your plan to continue using the service.</p></body></html>",
	}
	// Send the email
	created, _, err := emailer.BrevoClient.TransactionalEmailsApi.SendTransacEmail(ctx, emailContent)
	return created, err
}
