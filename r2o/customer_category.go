package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// CustomerCategoryService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type CustomerCategoryService service

type CustomerCategoryResponse struct {
	CustomerCategoryDescription *string `json:"customerCategory_description"`
	CustomerCategoryID          *int    `json:"customerCategory_id"`
	CustomerCategoryName        *string `json:"customerCategory_name"`
}

type CustomerCategoryRequest struct {
	CustomerCategoryDescription *string `json:"customerCategory_description"`
	CustomerCategoryName        *string `json:"customerCategory_name"`
}

func (as *CustomerCategoryService) GetCustomerCategories(ctx context.Context) (*[]CustomerCategoryResponse, error) {
	responseData := make([]CustomerCategoryResponse, 0)

	err := as.client.runHttpRequestWithContext(ctx, "customerCategories", http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CustomerCategoryService) GetCustomerCategory(ctx context.Context, id *int) (*CustomerCategoryResponse, error) {
	responseData := CustomerCategoryResponse{}

	u := fmt.Sprintf("customerCategories/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CustomerCategoryService) CreateCustomerCategory(ctx context.Context, data *CustomerCategoryRequest) (*CustomerCategoryResponse, error) {
	responseData := CustomerCategoryResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "customerCategories", http.MethodPut, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CustomerCategoryService) UpdateCustomerCategory(ctx context.Context, id *int, data *CustomerCategoryRequest) (*CustomerCategoryResponse, error) {
	responseData := CustomerCategoryResponse{}

	u := fmt.Sprintf("customerCategories/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CustomerCategoryService) DeleteCustomerCategory(ctx context.Context, id *int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	u := fmt.Sprintf("customerCategories/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodDelete, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
