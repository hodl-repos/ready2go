package r2o

import (
	"context"
	"net/http"
)

// TssService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type TssService service

func (as *TssService) DisableTss(ctx context.Context, accountId *int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	requestData := struct {
		AccountId *int `json:"accountId"`
	}{
		AccountId: accountId,
	}

	err := as.client.runHttpRequestWithContext(ctx, "company/disableTss", http.MethodPut, requestData, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
