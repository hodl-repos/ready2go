package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// CurrencyService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type CurrencyService service

type CurrencyResponse struct {
	CurrencyID    *int    `json:"currency_id"`
	CurrencyName  *string `json:"currency_name"`
	CurrencyShort *string `json:"currency_short"`
}

func (as *CurrencyService) GetCurrencies(ctx context.Context) (*[]CurrencyResponse, error) {
	responseData := make([]CurrencyResponse, 0)

	err := as.client.runHttpRequestWithContext(ctx, "currencies", http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CurrencyService) GetCurrency(ctx context.Context, id *int) (*CurrencyResponse, error) {
	responseData := CurrencyResponse{}

	u := fmt.Sprintf("currencies/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
