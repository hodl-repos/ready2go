package helper

import (
	"encoding/json"
	"errors"
	"io"
)

type RateLimitExceededError struct {
	InnerError RateLimitExceeded

	Err error
}

func (r *RateLimitExceededError) Error() string {
	return r.Err.Error()
}

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

func ParseRateLimitExceededError(responseBody *io.ReadCloser) error {
	data := RateLimitExceeded{}

	d := json.NewDecoder(*responseBody)
	d.DisallowUnknownFields()

	if err := d.Decode(&data); err != nil {
		return nil
	}

	if ValidateStruct(&data) != nil {
		return nil
	}

	return &RateLimitExceededError{
		InnerError: data,
		Err:        errors.New("RateLimit for AccountToken exceeded"),
	}
}
