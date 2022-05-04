package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// PaymentMethodTypeService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type PaymentMethodTypeService service

type PaymentMethodTypeResponse struct {
	PaymentTypeID         *int    `json:"paymentType_id"`
	PaymentTypeIdentifier *string `json:"paymentType_identifier"`
	PaymentTypeName       *string `json:"paymentType_name"`
}

func (as *PaymentMethodTypeService) GetPaymentMethodTypes(ctx context.Context) (*[]PaymentMethodTypeResponse, error) {
	responseData := make([]PaymentMethodTypeResponse, 0)

	err := as.client.runHttpRequestWithContext(ctx, "paymentMethodsType", http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *PaymentMethodTypeService) GetPaymentMethodType(ctx context.Context, id *int) (*PaymentMethodTypeResponse, error) {
	responseData := PaymentMethodTypeResponse{}

	u := fmt.Sprintf("paymentMethodsType/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
