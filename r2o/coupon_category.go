package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// CouponCategoryService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type CouponCategoryService service

type CouponCategoryResponse struct {
	CouponCategoryID          *int    `json:"couponCategory_id"`
	CouponCategoryName        *string `json:"couponCategory_name"`
	CouponCategoryDescription *string `json:"couponCategory_description"`
}

type CouponCategoryRequest struct {
	CouponCategoryName        *string `json:"couponCategory_name"`
	CouponCategoryDescription *string `json:"couponCategory_description"`
}

func (as *CouponCategoryService) GetCouponCategories(ctx context.Context, page *Pagination) (*[]CouponCategoryResponse, error) {
	responseData := make([]CouponCategoryResponse, 0)

	err := as.client.runHttpRequestWithParamsWithContext(ctx, "couponCategories", http.MethodGet, page, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CouponCategoryService) GetCouponCategory(ctx context.Context, id *int) (*CouponCategoryResponse, error) {
	responseData := CouponCategoryResponse{}

	u := fmt.Sprintf("couponCategories/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CouponCategoryService) CreateCouponCategory(ctx context.Context, data *CouponCategoryRequest) (*CouponCategoryResponse, error) {
	responseData := CouponCategoryResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "couponCategories", http.MethodPut, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CouponCategoryService) UpdateCouponCategory(ctx context.Context, id *int, data *CouponCategoryRequest) (*CouponCategoryResponse, error) {
	responseData := CouponCategoryResponse{}

	u := fmt.Sprintf("couponCategories/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CouponCategoryService) DeleteCouponCategory(ctx context.Context, id *int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	u := fmt.Sprintf("couponCategories/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodDelete, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
