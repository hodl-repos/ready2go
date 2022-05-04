package r2o

import (
	"context"
	"net/http"
)

// OrderService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type OrderService service

type GroupByType string

const (
	GroupByTypeTable   GroupByType = "table_id"
	GroupByTypeProduct GroupByType = "product_id"
)

type OrderRequest struct {
	Pagination

	//table_id or product_id or both
	GroupBy       *[]GroupByType `url:"groupBy"`
	TrainingsMode *bool          `url:"trainingsMode"`
	TableId       *int           `url:"table_id"`
	ProductId     *int           `url:"product_id"`
}

type OrderResponse struct {
	OrderComment                  *string `json:"order_comment"`
	OrderCourse                   *int    `json:"order_course"`
	OrderCreatedAt                *string `json:"order_created_at"`
	OrderDiscountName             *string `json:"order_discount_name"`
	OrderDiscountValueGross       *string `json:"order_discount_valueGross"`
	OrderDiscountValueNet         *string `json:"order_discount_valueNet"`
	OrderDiscountable             *bool   `json:"order_discountable"`
	OrderGroup                    *int    `json:"order_group"`
	OrderID                       *int    `json:"order_id"`
	OrderMergeable                *bool   `json:"order_mergeable"`
	OrderNumber                   *int    `json:"order_number"`
	OrderPriceGross               *string `json:"order_priceGross"`
	OrderPriceNet                 *string `json:"order_priceNet"`
	OrderProductItemnumber        *string `json:"order_product_itemnumber"`
	OrderProductName              *string `json:"order_product_name"`
	OrderProductPriceGrossPerUnit *string `json:"order_product_priceGrossPerUnit"`
	OrderProductPriceNetPerUnit   *string `json:"order_product_priceNetPerUnit"`
	OrderProductSerialnumber      *string `json:"order_product_serialnumber"`
	OrderQuantity                 *string `json:"order_quantity"`
	OrderRetour                   *bool   `json:"order_retour"`
	OrderStatus                   *int    `json:"order_status"`
	OrderVat                      *string `json:"order_vat"`
	OrderVatRate                  *string `json:"order_vatRate"`
	ProductID                     *int    `json:"product_id"`
	ProductgroupTypeID            *int    `json:"productgroup_type_id"`
	TableID                       *int    `json:"table_id"`
	UserID                        *int    `json:"user_id"`
}

func (as *OrderService) GetOrders(ctx context.Context, data *OrderRequest) (*[]OrderResponse, error) {
	responseData := make([]OrderResponse, 0)

	err := as.client.runHttpRequestWithParamsWithContext(ctx, "orders", http.MethodGet, data, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
