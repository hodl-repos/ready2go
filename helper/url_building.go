package helper

import (
	"fmt"
	"net/url"
	"strings"
)

//tries to combine a baseUrl and a path to a final url
//baseUrl must have a trailing slash
func BuildApiUrl(baseUrl *url.URL, path *string) (*string, error) {
	if !strings.HasSuffix(baseUrl.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", baseUrl)
	}

	u, err := baseUrl.Parse(*path)
	if err != nil {
		return nil, err
	}

	urlString := u.String()

	return &urlString, nil
}
