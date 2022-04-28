package r2o

import (
	"testing"
)

const (
	_TEST_DEV_TOKEN = ""
)

func TestGetAccountToken(t *testing.T) {
	GetAccountAccessToken(_TEST_DEV_TOKEN)
}
