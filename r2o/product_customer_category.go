package r2o

import (
	"context"
	"fmt"
	"net/http"
)

type ProductPriceCustomerCategoryResponse struct {
	CustomerCategoryID *int64  `json:"customerCategory_id"`
	PriceID            *int64  `json:"price_id"`
	PriceTimestamp     *string `json:"price_timestamp"`
	PriceValue         *string `json:"price_value"`
	PriceValueGross    *string `json:"price_valueGross"`
	PriceValueNet      *string `json:"price_valueNet"`
	ProductID          *int64  `json:"product_id"`
}

type ProductPriceCustomerCategoryRequest struct {
	CustomerCategoryID *int64  `json:"customerCategory_id"`
	PriceValue         *string `json:"price_value"`
}

func (as *ProductService) GetPriceCustomerCategories(ctx context.Context, productId *int) (*[]ProductPriceCustomerCategoryResponse, error) {
	responseData := make([]ProductPriceCustomerCategoryResponse, 0)

	u := fmt.Sprintf("products/%v/price/customerCategory", *productId)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductService) AddPriceCustomerCategory(ctx context.Context, productId *int, data *ProductPriceCustomerCategoryRequest) (*ProductPriceCustomerCategoryResponse, error) {
	responseData := ProductPriceCustomerCategoryResponse{}

	u := fmt.Sprintf("products/%v/price/customerCategory", *productId)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPut, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductService) GetPriceCustomerCategory(ctx context.Context, productId *int, customerCategoryId *int) (*ProductPriceCustomerCategoryResponse, error) {
	responseData := ProductPriceCustomerCategoryResponse{}

	u := fmt.Sprintf("products/%v/price/customerCategory/%v", *productId, *customerCategoryId)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductService) UpdatePriceCustomerCategory(ctx context.Context, productId *int, customerCategoryId *int, priceValue *string) (*ProductPriceCustomerCategoryResponse, error) {
	responseData := ProductPriceCustomerCategoryResponse{}

	requestBody := struct {
		PriceValue *string `json:"price_value"`
	}{
		PriceValue: priceValue,
	}

	u := fmt.Sprintf("products/%v/price/customerCategory/%v", *productId, *customerCategoryId)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, requestBody, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductService) DeletePriceCustomerCategory(ctx context.Context, productId *int, customerCategoryId *int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	u := fmt.Sprintf("products/%v/price/customerCategory/%v", *productId, *customerCategoryId)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
