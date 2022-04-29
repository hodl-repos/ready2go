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

type AccountResponse struct {
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

type AccountRequest struct {
	CompanyBusinessCity                       *string `json:"company_businessCity"`
	CompanyBusinessCountry                    *int64  `json:"company_businessCountry"`
	CompanyBusinessDateOfBirth                *string `json:"company_businessDateOfBirth"`
	CompanyBusinessPhoneNumber                *string `json:"company_businessPhoneNumber"`
	CompanyBusinessRegistrationNumber         *string `json:"company_businessRegistrationNumber"`
	CompanyBusinessStreet                     *string `json:"company_businessStreet"`
	CompanyBusinessZip                        *string `json:"company_businessZip"`
	CompanyCity                               *string `json:"company_city"`
	CompanyDisableVatReason                   *int64  `json:"company_disableVatReason"`
	CompanyEmail                              *string `json:"company_email"`
	CompanyGlobalLocationNumber               *string `json:"company_globalLocationNumber"`
	CompanyLegalForm                          *int64  `json:"company_legalForm"`
	CompanyName                               *string `json:"company_name"`
	CompanyPartnerData                        *string `json:"company_partnerData"`
	CompanyPhone                              *string `json:"company_phone"`
	CompanyRequireBillingMethodBeforeLiveMode *bool   `json:"company_requireBillingMethodBeforeLiveMode"`
	CompanyStreet                             *string `json:"company_street"`
	CompanyTaxIdentificationNumber            *string `json:"company_taxIdentificationNumber"`
	CompanyTaxOffice                          *int64  `json:"company_taxOffice"`
	CompanyUsername                           *string `json:"company_username"`
	CompanyWebsite                            *string `json:"company_website"`
	CompanyZip                                *string `json:"company_zip"`
	CurrencyID                                *int64  `json:"currency_id"`
	LanguageID                                *int64  `json:"language_id"`
	SyncToSalesforce                          *bool   `json:"syncToSalesforce"`
}

func (as *AccountService) GetAccountInfo(ctx context.Context) (*AccountResponse, error) {
	responseData := AccountResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "company", http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *AccountService) UpdateAccountInfo(ctx context.Context, account *AccountRequest) (*AccountResponse, error) {
	responseData := AccountResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "company", http.MethodPost, account, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

//maps the matching properties from the response to a new request object
func (as *AccountService) MapCustomerResponseToRequest(account *AccountResponse) AccountRequest {
	return AccountRequest{
		CompanyBusinessPhoneNumber: account.CompanyBusinessPhoneNumber,
		CompanyCity:                account.CompanyCity,
		CompanyName:                account.CompanyName,
		CompanyPartnerData:         account.CompanyPartnerData,
		CompanyPhone:               account.CompanyPhone,
		CompanyStreet:              account.CompanyStreet,
		CompanyUsername:            account.CompanyUsername,
		CompanyWebsite:             account.CompanyWebsite,
		CompanyZip:                 account.CompanyZip,
	}
}
