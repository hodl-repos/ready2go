package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// StornoService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type StornoService service

type StornoResponse struct {
	StornoCreatedAt   *string `json:"storno_created_at"`
	StornoDescription *string `json:"storno_description"`
	StornoID          *int    `json:"storno_id"`
	StornoName        *string `json:"storno_name"`
}

type StornoRequest struct {
	StornoDescription *string `json:"storno_description"`
	StornoName        *string `json:"storno_name"`
}

func (as *StornoService) GetStornos(ctx context.Context, page *Pagination) (*[]StornoResponse, error) {
	responseData := make([]StornoResponse, 0)

	err := as.client.runHttpRequestWithParamsWithContext(ctx, "storno", http.MethodGet, page, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *StornoService) GetStorno(ctx context.Context, id *int) (*StornoResponse, error) {
	responseData := StornoResponse{}

	u := fmt.Sprintf("storno/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *StornoService) CreateStorno(ctx context.Context, data *StornoRequest) (*StornoResponse, error) {
	responseData := StornoResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "storno", http.MethodPut, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *StornoService) UpdateStorno(ctx context.Context, id *int, data *StornoRequest) (*StornoResponse, error) {
	responseData := StornoResponse{}

	u := fmt.Sprintf("storno/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *StornoService) DeleteStorno(ctx context.Context, id *int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	u := fmt.Sprintf("storno/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodDelete, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
