package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// DiscountGroupService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type DiscountGroupService service

type DiscountGroupResponse struct {
	DiscountGroupActive      *bool   `json:"discountGroup_active"`
	DiscountGroupID          *int    `json:"discountGroup_id"`
	DiscountGroupDescription *string `json:"discountGroup_description"`
	DiscountGroupName        *string `json:"discountGroup_name"`
}

type DiscountGroupRequest struct {
	DiscountGroupActive      *bool   `json:"discountGroup_active"`
	DiscountGroupDescription *string `json:"discountGroup_description"`
	DiscountGroupName        *string `json:"discountGroup_name"`
}

func (as *DiscountGroupService) GetDiscountGroups(ctx context.Context) (*[]DiscountGroupResponse, error) {
	responseData := make([]DiscountGroupResponse, 0)

	err := as.client.runHttpRequestWithContext(ctx, "discountGroups", http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *DiscountGroupService) GetDiscountGroup(ctx context.Context, id *int) (*DiscountGroupResponse, error) {
	responseData := DiscountGroupResponse{}

	u := fmt.Sprintf("discountGroups/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *DiscountGroupService) CreateDiscountGroup(ctx context.Context, data *DiscountGroupRequest) (*DiscountGroupResponse, error) {
	responseData := DiscountGroupResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "discountGroups", http.MethodPut, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *DiscountGroupService) UpdateDiscountGroup(ctx context.Context, id *int, data *DiscountGroupRequest) (*DiscountGroupResponse, error) {
	responseData := DiscountGroupResponse{}

	u := fmt.Sprintf("discountGroups/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *DiscountGroupService) DeleteDiscountGroup(ctx context.Context, id *int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	u := fmt.Sprintf("discountGroups/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodDelete, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
