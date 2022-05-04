package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// TableAreaService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type TableAreaService service

type TableAreaResponse struct {
	TableAreaActive                *bool   `json:"tableArea_active"`
	TableAreaAllowTemporaryTables  *bool   `json:"tableArea_allowTemporaryTables"`
	TableAreaID                    *int    `json:"tableArea_id"`
	TableAreaIntelligentFavourites *int    `json:"tableArea_intelligentFavourites"`
	TableAreaName                  *string `json:"tableArea_name"`
	TableAreaOrder                 *int    `json:"tableArea_order"`
	TableAreaShortName             *string `json:"tableArea_shortName"`
}

type TableAreaRequest struct {
	TableAreaActive               *bool   `json:"tableArea_active"`
	TableAreaAllowTemporaryTables *bool   `json:"tableArea_allowTemporaryTables"`
	TableAreaName                 *string `json:"tableArea_name"`
	TableAreaOrder                *string `json:"tableArea_order"`
	TableAreaShortName            *string `json:"tableArea_shortName"`
}

func (as *TableAreaService) GetTableAreas(ctx context.Context, page *Pagination) (*[]TableAreaResponse, error) {
	responseData := make([]TableAreaResponse, 0)

	err := as.client.runHttpRequestWithParamsWithContext(ctx, "tableAreas", http.MethodGet, page, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *TableAreaService) GetTableArea(ctx context.Context, id *int) (*TableAreaResponse, error) {
	responseData := TableAreaResponse{}

	u := fmt.Sprintf("tableAreas/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *TableAreaService) CreateTableArea(ctx context.Context, data *TableAreaRequest) (*TableAreaResponse, error) {
	responseData := TableAreaResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "tableAreas", http.MethodPut, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *TableAreaService) UpdateTableArea(ctx context.Context, id *int, data *TableAreaRequest) (*TableAreaResponse, error) {
	responseData := TableAreaResponse{}

	u := fmt.Sprintf("tableAreas/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *TableAreaService) DeleteTableArea(ctx context.Context, id *int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	u := fmt.Sprintf("tableAreas/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodDelete, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
