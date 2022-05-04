package r2o

import (
	"context"
	"fmt"
	"net/http"
)

type ProductStockResponse struct {
	ProductID           *int64  `json:"product_id"`
	ProductReorderLevel *string `json:"product_reorderLevel"`
	ProductSafetyStock  *string `json:"product_safetyStock"`
	ProductStock        *string `json:"product_stock"`
	ProductUnit         *string `json:"product_unit"`
}

type ProductStockRequest struct {
	ProductReorderLevel *string `json:"product_reorderLevel"`
	ProductSafetyStock  *string `json:"product_safetyStock"`
	ProductStock        *string `json:"product_stock"`
	ProductStockDelta   *string `json:"product_stockDelta"`
}

func (as *ProductService) GetStock(ctx context.Context, id *int) (*ProductStockResponse, error) {
	responseData := ProductStockResponse{}

	u := fmt.Sprintf("products/%v/stock", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductService) GetStockByItemNumber(ctx context.Context, itemNumber *int) (*ProductStockResponse, error) {
	responseData := ProductStockResponse{}

	u := fmt.Sprintf("products/itemNumber/%v/stock", *itemNumber)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductService) UpdateStock(ctx context.Context, id *int, data *ProductStockRequest) (*ProductStockResponse, error) {
	responseData := ProductStockResponse{}

	u := fmt.Sprintf("products/%v/stock", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
