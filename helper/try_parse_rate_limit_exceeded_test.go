package helper

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	_RATE_LIMIT_EXCEEDED_RESPONSE_1 = `{"error":true,"requestId":"7039c569fd1577f5-VIE","msg":"You reached the rate limit of 60 requests per minute. Your IP (2001:871:25f:20ba:f9d9:ceaf:4d72:2163) blocked for at least 1 minutes! Try again later!","code":"rateLimitExceeded","details":{"rateLimitRequestPerMinute":60,"rateLimitRequest":60,"rateLimitMinutes":1}}`
	_RATE_LIMIT_EXCEEDED_RESPONSE_2 = `{"error":true,"requestId":"7039c569fd1577f5-VIE"}`
)

func TestRateLimitExceeded(t *testing.T) {
	var tests = []struct {
		index int
		json  string
		valid bool
	}{
		{0, _RATE_LIMIT_EXCEEDED_RESPONSE_1, true},
		{1, "_RATE_LIMIT_EXCEEDED_RESPONSE_1", false},
		{2, _RATE_LIMIT_EXCEEDED_RESPONSE_2, false},
	}

	for _, tt := range tests {
		testcase := tt //otherwise tt will be replaced in parallel testing before finishing

		t.Run(fmt.Sprintf("TestRateLimitExceeded_%v", testcase.index), func(t *testing.T) {
			t.Parallel()

			r := io.NopCloser(strings.NewReader(testcase.json)) // r type is io.ReadCloser

			errorStruct := TryParseRateLimitExceeded(&r)

			if testcase.valid {
				assert.NotNil(t, errorStruct)
			} else {
				assert.Nil(t, errorStruct)
			}
		})
	}
}
