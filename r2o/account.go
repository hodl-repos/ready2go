package r2o

import (
	"context"
	"net/http"

	"github.com/hodl-repos/ready2go/helper"
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
	CompanyBusinessCity                       *string `json:"company_businessCity" validate:"required"`
	CompanyBusinessCountry                    *int64  `json:"company_businessCountry" validate:"required"`
	CompanyBusinessDateOfBirth                *string `json:"company_businessDateOfBirth" validate:"required"`
	CompanyBusinessPhoneNumber                *string `json:"company_businessPhoneNumber" validate:"required"`
	CompanyBusinessRegistrationNumber         *string `json:"company_businessRegistrationNumber" validate:"required"`
	CompanyBusinessStreet                     *string `json:"company_businessStreet" validate:"required"`
	CompanyBusinessZip                        *string `json:"company_businessZip" validate:"required"`
	CompanyCity                               *string `json:"company_city" validate:"required"`
	CompanyDisableVatReason                   *int64  `json:"company_disableVatReason" validate:"required"`
	CompanyEmail                              *string `json:"company_email" validate:"required"`
	CompanyGlobalLocationNumber               *string `json:"company_globalLocationNumber" validate:"required"`
	CompanyLegalForm                          *int64  `json:"company_legalForm" validate:"required"`
	CompanyName                               *string `json:"company_name" validate:"required"`
	CompanyPartnerData                        *string `json:"company_partnerData" validate:"required"`
	CompanyPhone                              *string `json:"company_phone" validate:"required"`
	CompanyRequireBillingMethodBeforeLiveMode *bool   `json:"company_requireBillingMethodBeforeLiveMode" validate:"required"`
	CompanyStreet                             *string `json:"company_street" validate:"required"`
	CompanyTaxIdentificationNumber            *string `json:"company_taxIdentificationNumber" validate:"required"`
	CompanyTaxOffice                          *int64  `json:"company_taxOffice" validate:"required"`
	CompanyUsername                           *string `json:"company_username" validate:"required"`
	CompanyWebsite                            *string `json:"company_website" validate:"required"`
	CompanyZip                                *string `json:"company_zip" validate:"required"`
	CurrencyID                                *int64  `json:"currency_id" validate:"required"`
	LanguageID                                *int64  `json:"language_id" validate:"required"`
	SyncToSalesforce                          *bool   `json:"syncToSalesforce" validate:"required"`
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
	err := helper.ValidateStruct(account)

	if err != nil {
		return nil, err
	}

	responseData := AccountResponse{}

	err = as.client.runHttpRequestWithContext(ctx, "company", http.MethodPost, account, &responseData)

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
