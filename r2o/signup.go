package r2o

import (
	"context"
	"net/http"
)

// SignupService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type SignupService service

func (as *SignupService) ResendConfirmationEmail(ctx context.Context) (*MessageResponse, error) {
	responseData := MessageResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "signup/confirmation/resend-email", http.MethodPost, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

type SignupResponse struct {
	Admin *string `json:"admin"`
	Pos   *string `json:"pos"`
}

func (as *SignupService) SetPassword(ctx context.Context, password *string) (*SignupResponse, error) {
	responseData := SignupResponse{}

	requestData := struct {
		Password *string `json:"password"`
	}{
		Password: password,
	}

	err := as.client.runHttpRequestWithContext(ctx, "signup/set-password", http.MethodPost, requestData, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
