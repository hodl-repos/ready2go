package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// CountryService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type CouponService service

type CouponPurpose string

const (
	//every item applicable to this coupon has the same vat
	CouponPurposeSingle CouponPurpose = "single"

	//applicable items have multiple possible vat values
	CouponPurposeMultiple CouponPurpose = "multiple"
)

type CouponResponse struct {
	CouponCategoryID  *int           `json:"couponCategory_id"`
	CouponContainsVat *string        `json:"coupon_containsVat"`
	CouponCreatedAt   *string        `json:"coupon_created_at"`
	CouponID          *int           `json:"coupon_id"`
	CouponIdentifier  *string        `json:"coupon_identifier"`
	CouponIssuedAt    *string        `json:"coupon_issuedAt"`
	CouponName        *string        `json:"coupon_name"`
	CouponPurpose     *CouponPurpose `json:"coupon_purpose"`
	CouponTestMode    *bool          `json:"coupon_testMode"`
	CouponUpdatedAt   *string        `json:"coupon_updated_at"`
	CouponValidUntil  *string        `json:"coupon_validUntil"`
	CustomerID        *int           `json:"customer_id"`
}

type CouponRequest struct {
	CouponCategoryID  *string        `json:"couponCategory_id"`
	CouponContainsVat *string        `json:"coupon_containsVat"`
	CouponIssuedAt    *string        `json:"coupon_issuedAt"`
	CouponName        *string        `json:"coupon_name"`
	CouponPurpose     *CouponPurpose `json:"coupon_purpose"`
	CouponTestMode    *bool          `json:"coupon_testMode"`
	CouponType        *string        `json:"coupon_type"`
	CouponValidUntil  *string        `json:"coupon_validUntil"`
	CouponValue       *string        `json:"coupon_value"`
	CustomerID        *int           `json:"customer_id"`
}

type CouponUpdate struct {
	Value *string `json:"value"`
	Vat   *int    `json:"vat"`
}

func (as *CouponService) GetCoupons(ctx context.Context, page *Pagination) (*[]CouponResponse, error) {
	responseData := make([]CouponResponse, 0)

	err := as.client.runHttpRequestWithParamsWithContext(ctx, "coupons", http.MethodGet, page, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CouponService) GetCoupon(ctx context.Context, id *int) (*CouponResponse, error) {
	responseData := CouponResponse{}

	u := fmt.Sprintf("coupons/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CouponService) CreateCoupon(ctx context.Context, data *CouponRequest) (*CouponResponse, error) {
	responseData := CouponResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "coupons", http.MethodPut, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CouponService) UpdateCoupon(ctx context.Context, id *int, data *CouponRequest) (*CouponResponse, error) {
	responseData := CouponResponse{}

	u := fmt.Sprintf("coupons/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CouponService) DeleteCoupon(ctx context.Context, id *int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	u := fmt.Sprintf("coupons/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodDelete, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CouponService) Charge(ctx context.Context, id *int, data *CouponUpdate) (*MessageResponse, error) {
	responseData := MessageResponse{}

	u := fmt.Sprintf("coupons/%v/charge", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CouponService) Redeem(ctx context.Context, id *int, data *CouponUpdate) (*MessageResponse, error) {
	responseData := MessageResponse{}

	u := fmt.Sprintf("coupons/%v/redeem", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
