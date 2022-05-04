package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// PrintJobService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type PrintJobService service

type PrintJobResponse struct {
	PrintJobContent   *string `json:"printJob_content"`
	PrintJobCreatedAt *string `json:"printJob_createdAt"`
	PrintJobID        *int    `json:"printJob_id"`
	PrintJobImageURL  *string `json:"printJob_imageURL"`
	PrintJobPrintLogo *bool   `json:"printJob_printLogo"`
	PrintJobPrintedAt *string `json:"printJob_printedAt"`
	PrinterID         *int    `json:"printer_id"`
}

type PrintJobRequest struct {
	PrintJobContent   *string `json:"printJob_content"`
	PrintJobImageURL  *string `json:"printJob_imageURL"`
	PrintJobPrintLogo *bool   `json:"printJob_printLogo"`
	PrinterID         *int    `json:"printer_id"`
}

func (as *PrintJobService) GetPrintJobs(ctx context.Context) (*[]PrintJobResponse, error) {
	responseData := make([]PrintJobResponse, 0)

	err := as.client.runHttpRequestWithContext(ctx, "print", http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *PrintJobService) CreatePrintJobs(ctx context.Context, data *PrintJobRequest) (*PrintJobResponse, error) {
	responseData := PrintJobResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "print", http.MethodPut, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *PrintJobService) DeletePrintJob(ctx context.Context, id *int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	u := fmt.Sprintf("print/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodDelete, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
