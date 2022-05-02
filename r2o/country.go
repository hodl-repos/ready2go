package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// CountryService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type CountryService service

type Country struct {
	CountryID   *int    `json:"country_id"`
	CountryCode *string `json:"country_code"`
	CountryName *string `json:"country_name"`
}

func (as *CountryService) GetCountries(ctx context.Context) (*[]Country, error) {
	responseData := make([]Country, 0)

	err := as.client.runHttpRequestWithContext(ctx, "countries", http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CountryService) GetCountry(ctx context.Context, id *int) (*Country, error) {
	responseData := Country{}

	u := fmt.Sprintf("countries/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
