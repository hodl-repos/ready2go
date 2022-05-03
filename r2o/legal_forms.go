package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// LegalFormService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type LegalFormService service

type LegalForm struct {
	LegalFormCountryID  *int    `json:"legalForm_country_id"`
	LegalFormID         *int    `json:"legalForm_id"`
	LegalFormIdentifier *string `json:"legalForm_identifier"`
	LegalFormName       *string `json:"legalForm_name"`
}

func (as *LegalFormService) GetLegalForms(ctx context.Context) (*[]LegalForm, error) {
	responseData := make([]LegalForm, 0)

	err := as.client.runHttpRequestWithContext(ctx, "legalForms", http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *LegalFormService) GetLegalForm(ctx context.Context, id *int) (*LegalForm, error) {
	responseData := LegalForm{}

	u := fmt.Sprintf("legalForms/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
