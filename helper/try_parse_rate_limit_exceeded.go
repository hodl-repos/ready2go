package helper

import (
	"encoding/json"
	"io"
)

type RateLimitExceeded struct {
	Code    *string `json:"code" validate:"required"`
	Details struct {
		RateLimitMinutes          *int64 `json:"rateLimitMinutes" validate:"required"`
		RateLimitRequest          *int64 `json:"rateLimitRequest" validate:"required"`
		RateLimitRequestPerMinute *int64 `json:"rateLimitRequestPerMinute" validate:"required"`
	} `json:"details" validate:"required"`
	Error     *bool   `json:"error" validate:"required"`
	Msg       *string `json:"msg" validate:"required"`
	RequestID *string `json:"requestId" validate:"required"`
}

func TryParseRateLimitExceeded(responseBody *io.ReadCloser) *RateLimitExceeded {
	data := RateLimitExceeded{}

	d := json.NewDecoder(*responseBody)
	d.DisallowUnknownFields()

	if err := d.Decode(&data); err != nil {
		return nil
	}

	if ValidateStruct(&data) != nil {
		return nil
	}

	return &data
}
