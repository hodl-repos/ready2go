package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// PaymentMethodService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type PaymentMethodService service

type PaymentMethodResponse struct {
	PaymentTypeID         *int    `json:"paymentType_id"`
	PaymentAccountingCode *string `json:"payment_accountingCode"`
	PaymentDescription    *string `json:"payment_description"`
	PaymentID             *int    `json:"payment_id"`
	PaymentMarkAsPaid     *bool   `json:"payment_markAsPaid"`
	PaymentName           *string `json:"payment_name"`
}

type PaymentMethodRequest struct {
	PaymentTypeID      *int    `json:"paymentType_id"`
	PaymentDescription *string `json:"payment_description"`
	PaymentMarkAsPaid  *bool   `json:"payment_markAsPaid"`
	PaymentName        *string `json:"payment_name"`
}

func (as *PaymentMethodService) GetPaymentMethods(ctx context.Context) (*[]PaymentMethodResponse, error) {
	responseData := make([]PaymentMethodResponse, 0)

	err := as.client.runHttpRequestWithContext(ctx, "paymentMethods", http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *PaymentMethodService) GetPaymentMethod(ctx context.Context, id *int) (*PaymentMethodResponse, error) {
	responseData := PaymentMethodResponse{}

	u := fmt.Sprintf("paymentMethods/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *PaymentMethodService) CreatePaymentMethod(ctx context.Context, data *PaymentMethodRequest) (*PaymentMethodResponse, error) {
	responseData := PaymentMethodResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "paymentMethods", http.MethodPut, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *PaymentMethodService) UpdatePaymentMethod(ctx context.Context, id *int, data *PaymentMethodRequest) (*PaymentMethodResponse, error) {
	responseData := PaymentMethodResponse{}

	u := fmt.Sprintf("paymentMethods/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *PaymentMethodService) DeletePaymentMethods(ctx context.Context, id *int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	u := fmt.Sprintf("paymentMethods/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodDelete, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
