package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// LanguageService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type LanguageService service

type Language struct {
	LanguageCode   *string `json:"language_code"`
	LanguageID     *int    `json:"language_id"`
	LanguageLocale *string `json:"language_locale"`
	LanguageName   *string `json:"language_name"`
}

func (as *LanguageService) GetLanguages(ctx context.Context) (*[]Language, error) {
	responseData := make([]Language, 0)

	err := as.client.runHttpRequestWithContext(ctx, "languages", http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *LanguageService) GetLanguage(ctx context.Context, id *int) (*Language, error) {
	responseData := Language{}

	u := fmt.Sprintf("languages/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
