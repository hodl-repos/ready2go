package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// AccountingFinancialYearService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type AccountingFinancialYearService service

type AccountingFinancialYearResponse struct {
	AccountingFinancialYearActive    *bool  `json:"accountingFinancialYear_active"`
	AccountingFinancialYearID        *int64 `json:"accountingFinancialYear_id"`
	AccountingFinancialYearNumber    *int64 `json:"accountingFinancialYear_number"`
	AccountingFinancialYearTrunkYear *bool  `json:"accountingFinancialYear_trunkYear"`
}

func (as *AccountingFinancialYearService) GetAll(ctx context.Context) (*[]AccountingFinancialYearResponse, error) {
	var responseData []AccountingFinancialYearResponse

	err := as.client.runHttpRequestWithContext(ctx, "accounting/financialYears", http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *AccountingFinancialYearService) GetById(ctx context.Context, id *int64) (*AccountingFinancialYearResponse, error) {
	responseData := AccountingFinancialYearResponse{}

	u := fmt.Sprintf("accounting/financialYears/%v", id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
