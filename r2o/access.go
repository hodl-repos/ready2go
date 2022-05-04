package r2o

import (
	"context"
	"net/http"
)

// AccessService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type AccessService service

func (as *AccessService) Revoke(ctx context.Context) (*MessageResponse, error) {
	responseData := MessageResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "access/revoke", http.MethodPost, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
