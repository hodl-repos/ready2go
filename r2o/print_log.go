package r2o

import (
	"context"
	"net/http"
)

// PrintLogService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type PrintLogService service

type PrintLogResponse struct {
	PrintJobContent   string `json:"printJob_content"`
	PrintJobImageURL  string `json:"printJob_imageURL"`
	PrintJobPrintLogo bool   `json:"printJob_printLogo"`
	PrinterID         int64  `json:"printer_id"`
}

func (as *PrintLogService) GetPrintLogs(ctx context.Context, page *Pagination) (*[]PrintLogResponse, error) {
	responseData := make([]PrintLogResponse, 0)

	err := as.client.runHttpRequestWithParamsWithContext(ctx, "printLog", http.MethodGet, page, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
