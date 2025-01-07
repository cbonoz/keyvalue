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
