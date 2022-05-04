package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// ProductService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type ProductService service

type ProductRequest struct {
	ProductActive  *bool   `json:"product_active"`
	ProductBarcode *string `json:"product_barcode"`
	ProductBase    *struct {
		ProductID *int64 `json:"product_id"`
	} `json:"product_base"`
	ProductDescription       *string `json:"product_description"`
	ProductDiscountable      *bool   `json:"product_discountable"`
	ProductExternalReference *string `json:"product_externalReference"`
	ProductItemnumber        *string `json:"product_itemnumber"`
	ProductName              *string `json:"product_name"`
	ProductPrice             *string `json:"product_price"`
	ProductPriceIncludesVat  *bool   `json:"product_priceIncludesVat"`
	ProductSortIndex         *int64  `json:"product_sortIndex"`
	ProductStockEnabled      *bool   `json:"product_stock_enabled"`
	ProductStockReorderLevel *string `json:"product_stock_reorderLevel"`
	ProductStockSafetyStock  *string `json:"product_stock_safetyStock"`
	ProductStockUnit         *string `json:"product_stock_unit"`
	ProductStockValue        *string `json:"product_stock_value"`
	ProductType              *string `json:"product_type"`
	ProductVat               *string `json:"product_vat"`
	ProductVatID             *int64  `json:"product_vat_id"`
	Productgroup             *struct {
		ProductgroupID *int64 `json:"productgroup_id"`
	} `json:"productgroup"`
	ProductgroupID int64 `json:"productgroup_id"`
}

type ProductIngredient struct {
	IngredientID       int64  `json:"ingredient_id"`
	IngredientName     string `json:"ingredient_name"`
	IngredientQuantity string `json:"ingredient_quantity"`
	IngredientUnit     string `json:"ingredient_unit"`
}

type ProductVariation struct {
	VariationID    int64  `json:"variation_id"`
	VariationName  string `json:"variation_name"`
	VariationPrice string `json:"variation_price"`
	VariationStock string `json:"variation_stock"`
}

type ProductResponse struct {
	ProductAccountingCode            *string               `json:"product_accountingCode"`
	ProductAlternativeNameInPos      *string               `json:"product_alternativeNameInPos"`
	ProductAlternativeNameOnReceipts *string               `json:"product_alternativeNameOnReceipts"`
	ProductBarcode                   *string               `json:"product_barcode"`
	ProductColorClass                *string               `json:"product_colorClass"`
	ProductCreatedAt                 *string               `json:"product_created_at"`
	ProductCustomPrice               *bool                 `json:"product_customPrice"`
	ProductCustomQuantity            *bool                 `json:"product_customQuantity"`
	ProductDescription               *string               `json:"product_description"`
	ProductDiscountable              *bool                 `json:"product_discountable"`
	ProductExpressMode               *bool                 `json:"product_expressMode"`
	ProductExternalReference         *string               `json:"product_externalReference"`
	ProductFav                       *bool                 `json:"product_fav"`
	ProductHighlight                 *bool                 `json:"product_highlight"`
	ProductID                        *int64                `json:"product_id"`
	ProductIngredientsEnabled        *bool                 `json:"product_ingredients_enabled"`
	ProductItemnumber                *string               `json:"product_itemnumber"`
	ProductName                      *string               `json:"product_name"`
	ProductPrice                     *string               `json:"product_price"`
	ProductPriceIncludesVat          *bool                 `json:"product_priceIncludesVat"`
	ProductSideDishOrder             *bool                 `json:"product_sideDishOrder"`
	ProductSoldOut                   *bool                 `json:"product_soldOut"`
	ProductSortIndex                 *int64                `json:"product_sortIndex"`
	ProductStockEnabled              *bool                 `json:"product_stock_enabled"`
	ProductStockReorderLevel         *string               `json:"product_stock_reorderLevel"`
	ProductStockSafetyStock          *string               `json:"product_stock_safetyStock"`
	ProductStockUnit                 *string               `json:"product_stock_unit"`
	ProductStockValue                *string               `json:"product_stock_value"`
	ProductType                      *string               `json:"product_type"`
	ProductTypeID                    *int64                `json:"product_type_id"`
	ProductUpdatedAt                 *string               `json:"product_updated_at"`
	ProductVariationsEnabled         *bool                 `json:"product_variations_enabled"`
	ProductVat                       *string               `json:"product_vat"`
	Productgroup                     *ProductGroupResponse `json:"productgroup"`
	ProductgroupID                   *int64                `json:"productgroup_id"`
	Productingredient                *[]ProductIngredient  `json:"productingredient"`
	Productvariation                 *[]ProductVariation   `json:"productvariation"`
}

type ProductIncludeQuery struct {
	IncludeProductGroup       *bool `url:"includeProductGroup"`
	IncludeProductVariations  *bool `url:"includeProductVariations"`
	IncludeProductIngredients *bool `url:"includeProductIngredients"`
}

type ProductGetQuery struct {
	Pagination
	ProductIncludeQuery
	ProductGroupId *int    `url:"productgroup_id"`
	ItemNumber     *string `url:"itemNumber"`
	Name           *string `url:"name"`
	Barcode        *string `url:"barcode"`
	Keywords       *string `url:"q"`
}

func (as *ProductService) GetProducts(ctx context.Context, page *ProductGetQuery) (*[]ProductResponse, error) {
	responseData := make([]ProductResponse, 0)

	err := as.client.runHttpRequestWithParamsWithContext(ctx, "products", http.MethodGet, page, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductGroupService) GetProduct(ctx context.Context, id *int, query *ProductIncludeQuery) (*ProductResponse, error) {
	responseData := ProductResponse{}

	u := fmt.Sprintf("products/%v", *id)
	err := as.client.runHttpRequestWithParamsWithContext(ctx, u, http.MethodGet, query, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductService) CreateProduct(ctx context.Context, data *ProductRequest) (*ProductResponse, error) {
	responseData := ProductResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "products", http.MethodPut, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductService) UpdateProduct(ctx context.Context, id *int, data *ProductRequest) (*ProductResponse, error) {
	responseData := ProductResponse{}

	u := fmt.Sprintf("products/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *ProductService) DeleteProduct(ctx context.Context, id *int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	u := fmt.Sprintf("products/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodDelete, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

type ProductStockBatchRequest struct {
	ProductID    string `json:"product_id"`
	ProductStock string `json:"product_stock"`
}

type ProductStockBatchResponse struct {
	UpdatedIds struct {
		ID int64 `json:"id"`
	} `json:"updatedIds"`
	UpdatedItemNumbers struct {
		Number int64 `json:"number"`
	} `json:"updatedItemNumbers"`
}

func (as *ProductService) BatchProductStock(ctx context.Context, data *ProductStockBatchRequest) (*ProductStockBatchResponse, error) {
	responseData := ProductStockBatchResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "products/batch/stock", http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
