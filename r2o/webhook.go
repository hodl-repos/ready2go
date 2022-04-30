package r2o

import (
	"context"
	"net/http"

	"github.com/hodl-repos/ready2go/helper"
)

// AccountingFinancialYearService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type WebhookService service

type WebhookUrlData struct {
	WebhookUrl *string `json:"webhookUrl" validate:"required"`
}

type WebhookEventResponse struct {
	ActiveEvents    *[]WebhookEvent `json:"activeEvents"`
	AvailableEvents *[]WebhookEvent `json:"availableEvents"`
}

type WebhookEventChangeResponse struct {
	Error   *bool   `json:"error"`
	Success *bool   `json:"success"`
	Message *string `json:"msg"`
}

func (as *WebhookService) GetUrl(ctx context.Context) (*WebhookUrlData, error) {
	responseData := WebhookUrlData{}

	err := as.client.runHttpRequestWithContext(ctx, "webhook", http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *WebhookService) UpdateUrl(ctx context.Context, data *WebhookUrlData) (*WebhookUrlData, error) {
	err := helper.ValidateStruct(data)

	if err != nil {
		return nil, err
	}

	responseData := WebhookUrlData{}

	err = as.client.runHttpRequestWithContext(ctx, "webhook", http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

type WebhookEvent string

const (
	WebhookEventProductCreated      WebhookEvent = "product.created"
	WebhookEventProductUpdated      WebhookEvent = "product.updated"
	WebhookEventProductDeleted      WebhookEvent = "product.deleted"
	WebhookEventProductGroupCreated WebhookEvent = "productGroup.created"
	WebhookEventProductGroupUpdated WebhookEvent = "productGroup.updated"
	WebhookEventProductGroupDeleted WebhookEvent = "productGroup.deleted"
	WebhookEventInvoiceCreated      WebhookEvent = "invoice.created"
	WebhookEventOrderItemCreated    WebhookEvent = "orderItem.created"
)

func (as *WebhookService) AddEvent(ctx context.Context, event *WebhookEvent) (*WebhookEventChangeResponse, error) {
	type request struct {
		AddEvent *string `json:"addEvent"`
	}

	eventString := (string)(*event)

	requestData := request{
		AddEvent: &eventString,
	}

	responseData := WebhookEventChangeResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "webhook/events", http.MethodPost, requestData, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *WebhookService) RemoveEvent(ctx context.Context, event *WebhookEvent) (*WebhookEventChangeResponse, error) {
	type request struct {
		RemoveEvent *string `json:"removeEvent"`
	}

	eventString := (string)(*event)

	requestData := request{
		RemoveEvent: &eventString,
	}

	responseData := WebhookEventChangeResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "webhook/events", http.MethodPost, requestData, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *WebhookService) FindEvents(ctx context.Context) (*WebhookEventResponse, error) {
	responseData := WebhookEventResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "webhook/events", http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
