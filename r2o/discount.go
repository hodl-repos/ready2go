package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// DiscountService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type DiscountService service

type DiscountResponse struct {
	DiscountGroupID     *int    `json:"discountGroup_id"`
	DiscountActive      *bool   `json:"discount_active"`
	DiscountCreatedAt   *string `json:"discount_created_at"`
	DiscountDescription *string `json:"discount_description"`
	DiscountID          *int    `json:"discount_id"`
	DiscountName        *string `json:"discount_name"`
	DiscountOrder       *int    `json:"discount_order"`
	DiscountUnit        *string `json:"discount_unit"`
	DiscountUpdatedAt   *string `json:"discount_updated_at"`
	DiscountValue       *string `json:"discount_value"`
}

type DiscountRequest struct {
	DiscountGroupID     *int    `json:"discountGroup_id"`
	DiscountActive      *bool   `json:"discount_active"`
	DiscountDescription *string `json:"discount_description"`
	DiscountName        *string `json:"discount_name"`
	DiscountOrder       *int    `json:"discount_order"`
	DiscountUnit        *string `json:"discount_unit"`
	DiscountValue       *string `json:"discount_value"`
}

func (as *DiscountService) GetDiscounts(ctx context.Context) (*[]DiscountResponse, error) {
	responseData := make([]DiscountResponse, 0)

	err := as.client.runHttpRequestWithContext(ctx, "discounts", http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *DiscountService) GetDiscount(ctx context.Context, id *int) (*DiscountResponse, error) {
	responseData := DiscountResponse{}

	u := fmt.Sprintf("discounts/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *DiscountService) CreateDiscount(ctx context.Context, data *DiscountRequest) (*DiscountResponse, error) {
	responseData := DiscountResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "discounts", http.MethodPut, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *DiscountService) UpdateDiscount(ctx context.Context, id *int, data *DiscountRequest) (*DiscountResponse, error) {
	responseData := DiscountResponse{}

	u := fmt.Sprintf("discounts/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *DiscountService) DeleteDiscount(ctx context.Context, id *int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	u := fmt.Sprintf("discounts/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodDelete, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
