package r2o

import (
	"context"
	"net/http"
)

// DailyReportService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type DailyReportService service

type DailyReportResponse struct {
	MessageResponse
	Status *string `json:"status"`
}

func (as *DailyReportService) OpenDay(ctx context.Context) (*MessageResponse, error) {
	responseData := MessageResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "dailyReport/open", http.MethodPut, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *DailyReportService) CloseDay(ctx context.Context) (*MessageResponse, error) {
	responseData := MessageResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "dailyReport/close", http.MethodPut, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *DailyReportService) GetStatus(ctx context.Context) (*DailyReportResponse, error) {
	responseData := DailyReportResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "dailyReport/status", http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
