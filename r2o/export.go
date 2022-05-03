package r2o

import (
	"context"
	"net/http"
)

// ExportSerivce handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type ExportService service

type FileResponse struct {
	Error    *bool   `json:"error"`
	File     *string `json:"file"`
	Filename *string `json:"filename"`
}

type JobResponse struct {
	JobStatusFinishedAt string `json:"jobStatus_finishedAt"`
	JobStatusID         int64  `json:"jobStatus_id"`
	JobStatusStartedAt  string `json:"jobStatus_startedAt"`
	JobStatusStatus     string `json:"jobStatus_status"`
}

type CashbookExportFormat string

const (
	CashbookExportFormatPdf   CashbookExportFormat = "pdf"
	CashbookExportFormatExcel CashbookExportFormat = "excel"
)

type CashbookExportRequest struct {
	DateFrom   *string               `json:"dateFrom"`
	DateTo     *string               `json:"dateTo"`
	FileFormat *CashbookExportFormat `json:"fileFormat"`
}

func (as *ExportService) ExportCashbook(ctx context.Context, data *CashbookExportRequest) (*JobResponse, error) {
	responseData := JobResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "cashbook/export", http.MethodGet, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

type AccountancyExportFormat string

const (
	AccountancyExportFormatBmd     AccountancyExportFormat = "bmd"
	AccountancyExportFormatBmdNtcs AccountancyExportFormat = "bmd-ntcs"
	AccountancyExportFormatDatev   AccountancyExportFormat = "datev"
	AccountancyExportFormatSap     AccountancyExportFormat = "sap"
	AccountancyExportFormatRzl     AccountancyExportFormat = "rzl"
)

type AccountancyExportRequest struct {
	DataSource     *string                  `json:"dataSource"`
	Day            *int                     `json:"day"`
	ExportCashbook *bool                    `json:"exportCashbook"`
	FinancialYear  *int                     `json:"financialYear"`
	Format         *AccountancyExportFormat `json:"format"`
	From           *string                  `json:"from"`
	Month          *int                     `json:"month"`
	To             *string                  `json:"to"`
	Year           *int                     `json:"year"`
}

func (as *ExportService) ExportAccountancy(ctx context.Context, data *AccountancyExportRequest) (*FileResponse, error) {
	responseData := FileResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "export/accountancy", http.MethodGet, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

type AccountancyCashbookExportRequest struct {
	DateFrom *string `json:"dateFrom"`
	DateTo   *string `json:"dateTo"`
	Day      *int    `json:"day"`
	Month    *int    `json:"month"`
	Year     *int    `json:"year"`
}

func (as *ExportService) ExportAccountancyCashbook(ctx context.Context, data *AccountancyCashbookExportRequest) (*FileResponse, error) {
	responseData := FileResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "export/accountancy/cashbook", http.MethodGet, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

type ExportSalesRequest struct {
	Month *int `json:"month"`
	Year  *int `json:"year"`
}

type ExportSalesResponse struct {
	MessageResponse
	JobStatusId *int `json:"jobStatus_id"`
}

func (as *ExportService) ExportSales(ctx context.Context, data *ExportSalesRequest) (*ExportSalesResponse, error) {
	responseData := ExportSalesResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "export/sales", http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
