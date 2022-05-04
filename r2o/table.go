package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// TableService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type TableService service

type TableResponse struct {
	TableAreaID            int64  `json:"tableArea_id"`
	TableCheckoutMode      bool   `json:"table_checkoutMode"`
	TableCreatedAt         string `json:"table_created_at"`
	TableDescription       string `json:"table_description"`
	TableID                int64  `json:"table_id"`
	TableIsTemporay        bool   `json:"table_isTemporay"`
	TableLastInteractionAt string `json:"table_lastInteractionAt"`
	TableName              string `json:"table_name"`
	TableOrder             int64  `json:"table_order"`
	TableUpdatedAt         string `json:"table_updated_at"`
}

type TableRequest struct {
	TableAreaID      int64  `json:"tableArea_id"`
	TableDescription string `json:"table_description"`
	TableName        string `json:"table_name"`
}

func (as *TableService) GetTables(ctx context.Context, page *Pagination) (*[]TableResponse, error) {
	responseData := make([]TableResponse, 0)

	err := as.client.runHttpRequestWithParamsWithContext(ctx, "tables", http.MethodGet, page, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *TableService) GetTable(ctx context.Context, id *int) (*TableResponse, error) {
	responseData := TableResponse{}

	u := fmt.Sprintf("tables/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *TableService) CreateTable(ctx context.Context, data *TableRequest) (*TableResponse, error) {
	responseData := TableResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "tables", http.MethodPut, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *TableService) UpdateTable(ctx context.Context, id *int, data *TableRequest) (*TableResponse, error) {
	responseData := TableResponse{}

	u := fmt.Sprintf("tables/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *TableService) DeleteTable(ctx context.Context, id *int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	u := fmt.Sprintf("tables/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodDelete, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
