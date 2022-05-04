package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// PrinterService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type PrinterService service

type PrinterResponse struct {
	PrinterCharacters   *int    `json:"printer_characters"`
	PrinterCreatedAt    *string `json:"printer_createdAt"`
	PrinterDescription  *string `json:"printer_description"`
	PrinterDeviceName   *string `json:"printer_deviceName"`
	PrinterDoubleHeight *bool   `json:"printer_doubleHeight"`
	PrinterID           *int    `json:"printer_id"`
	PrinterIPAddress    *string `json:"printer_ipAddress"`
	PrinterManufacturer *string `json:"printer_manufacturer"`
	PrinterModel        *string `json:"printer_model"`
	PrinterName         *string `json:"printer_name"`
}

func (as *PrinterService) GetPrinters(ctx context.Context) (*[]PrinterResponse, error) {
	responseData := make([]PrinterResponse, 0)

	err := as.client.runHttpRequestWithContext(ctx, "printer", http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *PrinterService) GetPrinter(ctx context.Context, id *int) (*PrinterResponse, error) {
	responseData := PrinterResponse{}
	u := fmt.Sprintf("printer/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *PrinterService) TestPrint(ctx context.Context) (*MessageResponse, error) {
	responseData := MessageResponse{}
	err := as.client.runHttpRequestWithContext(ctx, "printer/testPrint", http.MethodPut, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
