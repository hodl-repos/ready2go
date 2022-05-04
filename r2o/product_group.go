package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// ProductGroupService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type ProductGroupService service

type ProductGroupResponse struct {
	ProductgroupActive      *bool   `json:"productgroup_active"`
	ProductgroupCreatedAt   *string `json:"productgroup_created_at"`
	ProductgroupDescription *string `json:"productgroup_description"`
	ProductgroupID          *int64  `json:"productgroup_id"`
	ProductgroupName        *string `json:"productgroup_name"`
	ProductgroupShortcut    *string `json:"productgroup_shortcut"`
	ProductgroupSortIndex   *int64  `json:"productgroup_sortIndex"`
	ProductgroupTypeID      *int64  `json:"productgroup_type_id"`
	ProductgroupUpdatedAt   *string `json:"productgroup_updated_at"`
}

type ProductGroupRequest struct {
	ProductgroupAccountingCode        *string `json:"productgroup_accountingCode"`
	ProductgroupAccountingCodeName    *string `json:"productgroup_accountingCodeName"`
	ProductgroupAccountingCodeVatRate *string `json:"productgroup_accountingCodeVatRate"`
	ProductgroupActive                *bool   `json:"productgroup_active"`
	ProductgroupDescription           *string `json:"productgroup_description"`
	ProductgroupName                  *string `json:"productgroup_name"`
	ProductgroupParent                *int64  `json:"productgroup_parent"`
	ProductgroupShortcut              *string `json:"productgroup_shortcut"`
	ProductgroupSortIndex             *int64  `json:"productgroup_sortIndex"`
	ProductgroupTypeID                *int64  `json:"productgroup_type_id"`
}

func (as *ProductGroupService) GetProductGroups(ctx context.Context, page *Pagination) (*[]ProductGroupResponse, error) {
	responseData := make([]ProductGroupResponse, 0)

	err := as.client.runHttpRequestWithParamsWithContext(ctx, "productgroups", http.MethodGet, page, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductGroupService) GetProductGroup(ctx context.Context, id *int) (*ProductGroupResponse, error) {
	responseData := ProductGroupResponse{}

	u := fmt.Sprintf("productgroups/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductGroupService) CreateProductGroup(ctx context.Context, data *ProductGroupRequest) (*ProductGroupResponse, error) {
	responseData := ProductGroupResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "productgroups", http.MethodPut, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductGroupService) UpdateProductGroup(ctx context.Context, id *int, data *ProductGroupRequest) (*ProductGroupResponse, error) {
	responseData := ProductGroupResponse{}

	u := fmt.Sprintf("productgroups/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductGroupService) DeleteProductGroup(ctx context.Context, id *int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	u := fmt.Sprintf("productgroups/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodDelete, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductGroupService) GetProductsByGroup(ctx context.Context, groupId *int) (*[]ProductResponse, error) {
	responseData := make([]ProductResponse, 0)

	u := fmt.Sprintf("productgroups/%v/products", *groupId)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodDelete, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
