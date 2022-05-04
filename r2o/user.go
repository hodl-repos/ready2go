package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// UserService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type UserService service

type UserResponse struct {
	UserFirstName     *string `json:"user_firstName"`
	UserID            *int    `json:"user_id"`
	UserLastActionAt  *string `json:"user_lastActionAt"`
	UserLastLoginAt   *string `json:"user_lastLoginAt"`
	UserLastName      *string `json:"user_lastName"`
	UserPrintAccess   *int    `json:"user_printAccess"`
	UserPrinter       *int    `json:"user_printer"`
	UserTrainingsMode *bool   `json:"user_trainingsMode"`
	UserUsername      *string `json:"user_username"`
}

type UserRequest struct {
	UserFirstName     *string `json:"user_firstName"`
	UserLastName      *string `json:"user_lastName"`
	UserPassword      *string `json:"user_password"`
	UserPrintAccess   *int    `json:"user_printAccess"`
	UserPrinter       *int    `json:"user_printer"`
	UserTrainingsMode *bool   `json:"user_trainingsMode"`
	UserUsername      *string `json:"user_username"`
}

func (as *UserService) GetUsers(ctx context.Context, page *Pagination) (*[]UserResponse, error) {
	responseData := make([]UserResponse, 0)

	err := as.client.runHttpRequestWithParamsWithContext(ctx, "users", http.MethodGet, page, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *UserService) GetUser(ctx context.Context, id *int) (*UserResponse, error) {
	responseData := UserResponse{}

	u := fmt.Sprintf("users/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *UserService) CreateUser(ctx context.Context, data *UserRequest) (*UserResponse, error) {
	responseData := UserResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "users", http.MethodPut, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *UserService) UpdateUser(ctx context.Context, id *int, data *UserRequest) (*UserResponse, error) {
	responseData := UserResponse{}

	u := fmt.Sprintf("users/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *UserService) DeleteUser(ctx context.Context, id *int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	u := fmt.Sprintf("users/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodDelete, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *UserService) UpdatePassword(ctx context.Context, id *int, newPassword *string) (*UserResponse, error) {
	responseData := UserResponse{}

	request := struct {
		UserPassword *string `json:"user_password"`
	}{
		UserPassword: newPassword,
	}

	u := fmt.Sprintf("users/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPatch, &request, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *UserService) SetPrinterAccessCode(ctx context.Context, id *int, newCode *int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	request := struct {
		PrintAccessId *int `json:"printAccess_id"`
	}{
		PrintAccessId: newCode,
	}

	u := fmt.Sprintf("users/%v/setPrintAccessCode", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, &request, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

type SignInTokenResponse struct {
	Error     bool   `json:"error"`
	ExpiresAt string `json:"expiresAt"`
	IssuedAt  string `json:"issuedAt"`
	LoginURI  string `json:"loginURI"`
	Token     string `json:"token"`
}

func (as *UserService) GetSigningTokenWithGet(ctx context.Context, id *int, expiresInSeconds *int, userId *string) (*SignInTokenResponse, error) {
	responseData := SignInTokenResponse{}

	request := struct {
		CrmUserId *string `json:"crmUserId"`
	}{
		CrmUserId: userId,
	}

	params := struct {
		ExpiresIn *int `url:"expiresIn"`
	}{
		ExpiresIn: expiresInSeconds,
	}

	u := fmt.Sprintf("users/%v/signInToken", *id)
	err := as.client.runHttpRequestWithParamsWithContext(ctx, u, http.MethodGet, &params, &request, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *UserService) GetSigningTokenWithPost(ctx context.Context, id *int, expiresInSeconds *int, userId *string) (*SignInTokenResponse, error) {
	responseData := SignInTokenResponse{}

	request := struct {
		CrmUserId *string `json:"crmUserId"`
	}{
		CrmUserId: userId,
	}

	params := struct {
		ExpiresIn *int `url:"expiresIn"`
	}{
		ExpiresIn: expiresInSeconds,
	}

	u := fmt.Sprintf("users/%v/signInToken", *id)
	err := as.client.runHttpRequestWithParamsWithContext(ctx, u, http.MethodPost, &params, &request, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
