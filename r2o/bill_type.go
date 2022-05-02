package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// BillService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type BillTypeService service

type BillType struct {
	BillTypeID     *int    `json:"billType_id"`
	BillTypeName   *string `json:"billType_name"`
	BillTypeSymbol *string `json:"billType_symbol"`
}

func (as *BillTypeService) GetBillTypes(ctx context.Context) (*[]BillType, error) {
	responseData := make([]BillType, 0)

	err := as.client.runHttpRequestWithContext(ctx, "billTypes", http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *BillTypeService) GetBillType(ctx context.Context, id *int) (*BillType, error) {
	responseData := BillType{}

	u := fmt.Sprintf("billTypes/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
