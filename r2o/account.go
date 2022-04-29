package r2o

import (
	"context"
	"net/http"
)

// AccountService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type AccountService service

type AccountGetResponse struct {
	CompanyID                  *int    `json:"company_id"`
	CompanyBranch              *int    `json:"company_branch"`
	CompanyUsername            *string `json:"company_username"`
	CompanyAccountNumber       *string `json:"company_accountNumber"`
	CompanyName                *string `json:"company_name"`
	CompanyBusiness            *string `json:"company_business"`
	CompanyFirstName           *string `json:"company_firstName"`
	CompanyLastName            *string `json:"company_lastName"`
	CompanyVatID               *string `json:"company_vatId"`
	CompanyStreet              *string `json:"company_street"`
	CompanyCity                *string `json:"company_city"`
	CompanyZip                 *string `json:"company_zip"`
	CompanyPhone               *string `json:"company_phone"`
	CompanyBusinessPhoneNumber *string `json:"company_businessPhoneNumber"`
	CompanyWebsite             *string `json:"company_website"`
	CompanyReferrer            *string `json:"company_referrer"`
	CompanyFailedLoginAttempts *int    `json:"company_failedLoginAttempts"`
	CompanyTrainingsMode       *bool   `json:"company_trainingsMode"`
	CompanyLiveMode            *bool   `json:"company_liveMode"`
	CompanyLiveModeStartedAt   *string `json:"company_liveModeStartedAt"`
	LanguageID                 *int    `json:"language_id"`
	CurrencyID                 *int    `json:"currency_id"`
	CountryID                  *string `json:"country_id"`
	CompanyPartnerData         *string `json:"company_partnerData"`
}

type AccountPostRequest struct {
	CompanyAccountNumber       string `json:"company_accountNumber"`
	CompanyBranch              int64  `json:"company_branch"`
	CompanyBusiness            string `json:"company_business"`
	CompanyBusinessPhoneNumber string `json:"company_businessPhoneNumber"`
	CompanyCity                string `json:"company_city"`
	CompanyFailedLoginAttempts int64  `json:"company_failedLoginAttempts"`
	CompanyFirstName           string `json:"company_firstName"`
	CompanyID                  int64  `json:"company_id"`
	CompanyLastName            string `json:"company_lastName"`
	CompanyLiveMode            bool   `json:"company_liveMode"`
	CompanyLiveModeStartedAt   string `json:"company_liveModeStartedAt"`
	CompanyName                string `json:"company_name"`
	CompanyPartnerData         string `json:"company_partnerData"`
	CompanyPhone               string `json:"company_phone"`
	CompanyReferrer            string `json:"company_referrer"`
	CompanyStreet              string `json:"company_street"`
	CompanyTrainingsMode       bool   `json:"company_trainingsMode"`
	CompanyUsername            string `json:"company_username"`
	CompanyVatID               string `json:"company_vatId"`
	CompanyWebsite             string `json:"company_website"`
	CompanyZip                 string `json:"company_zip"`
	CountryID                  string `json:"country_id"`
	CurrencyID                 int64  `json:"currency_id"`
	LanguageID                 int64  `json:"language_id"`
}

func (as *AccountService) GetAccountInfo(ctx context.Context) (*AccountGetResponse, error) {
	responseData := AccountGetResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "company", http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *AccountService) UpdateAccountInfo(ctx context.Context, account *AccountPostRequest) error {
	return as.client.runHttpRequestWithContext(ctx, "company", http.MethodPost, account, nil)
}

func (as *AccountService) MapCustomerResponseToRequest(account *AccountGetResponse) AccountPostRequest {
	return AccountPostRequest{}
}
