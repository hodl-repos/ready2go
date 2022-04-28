package helper

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlBuilding_Fail_No_Trailing_Slash(t *testing.T) {
	t.Parallel()

	baseURL, _ := url.Parse("https://www.google.com")

	testPath := "hi"

	resultUrl, err := BuildApiUrl(baseURL, &testPath)

	assert.NotNil(t, err)
	assert.Nil(t, resultUrl)
}

func TestUrlBuilding_Pass_1(t *testing.T) {
	t.Parallel()

	baseURL, _ := url.Parse("https://www.google.com/")

	testPath := "hi"

	resultUrl, err := BuildApiUrl(baseURL, &testPath)

	assert.Nil(t, err)
	assert.NotNil(t, resultUrl)
	assert.Equal(t, "https://www.google.com/hi", *resultUrl)
}

func TestUrlBuilding_Pass_2(t *testing.T) {
	t.Parallel()

	baseURL, _ := url.Parse("https://www.google.com/")

	testPath := "/hi"

	resultUrl, err := BuildApiUrl(baseURL, &testPath)

	assert.Nil(t, err)
	assert.NotNil(t, resultUrl)
	assert.Equal(t, "https://www.google.com/hi", *resultUrl)
}
