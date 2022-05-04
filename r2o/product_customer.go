package r2o

import (
	"context"
	"fmt"
	"net/http"
)

type ProductPriceCustomerResponse struct {
	CustomerID      *int64  `json:"customer_id"`
	PriceID         *int64  `json:"price_id"`
	PriceTimestamp  *string `json:"price_timestamp"`
	PriceValue      *string `json:"price_value"`
	PriceValueGross *string `json:"price_valueGross"`
	PriceValueNet   *string `json:"price_valueNet"`
	ProductID       *int64  `json:"product_id"`
}

type ProductPriceCustomerRequest struct {
	CustomerID *int64  `json:"customer_id"`
	PriceValue *string `json:"price_value"`
}

func (as *ProductService) GetPriceCustomers(ctx context.Context, productId *int) (*[]ProductPriceCustomerResponse, error) {
	responseData := make([]ProductPriceCustomerResponse, 0)

	u := fmt.Sprintf("products/%v/price/customer", *productId)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductService) AddPriceCustomer(ctx context.Context, productId *int, data *ProductPriceCustomerRequest) (*ProductPriceCustomerResponse, error) {
	responseData := ProductPriceCustomerResponse{}

	u := fmt.Sprintf("products/%v/price/customer", *productId)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPut, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductService) GetPriceCustomer(ctx context.Context, productId *int, customerCategoryId *int) (*ProductPriceCustomerResponse, error) {
	responseData := ProductPriceCustomerResponse{}

	u := fmt.Sprintf("products/%v/price/customer/%v", *productId, *customerCategoryId)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductService) UpdatePriceCustomer(ctx context.Context, productId *int, customerCategoryId *int, priceValue *string) (*ProductPriceCustomerResponse, error) {
	responseData := ProductPriceCustomerResponse{}

	requestBody := struct {
		PriceValue *string `json:"price_value"`
	}{
		PriceValue: priceValue,
	}

	u := fmt.Sprintf("products/%v/price/customer/%v", *productId, *customerCategoryId)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, requestBody, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductService) DeletePriceCustomer(ctx context.Context, productId *int, customerCategoryId *int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	u := fmt.Sprintf("products/%v/price/customer/%v", *productId, *customerCategoryId)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
