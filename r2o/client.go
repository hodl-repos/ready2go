package r2o

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	_BASE_API_URL_V1 = "https://api.ready2order.com/v1"
)

type service struct {
	client *Client
}

// A Client manages communication with the ready2order API.
type Client struct {
	client *http.Client // HTTP client used to communicate with the API.

	accountApiToken *string

	// Base URL for API requests. Defaults to the public GitHub API, but can be
	// set to a domain endpoint to use with GitHub Enterprise. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the GitHub API.
	Account *AccountService
}

// NewClient returns a new GitHub API client. If a nil httpClient is
// provided, a new http.Client will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client, accountApiToken *string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(_BASE_API_URL_V1)

	c := &Client{client: httpClient, accountApiToken: accountApiToken, BaseURL: baseURL}
	c.common.client = c

	c.Account = (*AccountService)(&c.common)

	return c
}

//needs the context WITH authorization
func (c *Client) runHttpRequest(ctx context.Context, path, method string, requestData interface{}, responseData interface{}) error {
	//URL BUILDING
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(path)
	if err != nil {
		return err
	}

	//REQUEST BODY
	var requestBody io.Reader = nil

	if requestData != nil {
		json_data, err := json.Marshal(requestData)

		if err != nil {
			return err
		}

		requestBody = bytes.NewBuffer(json_data)
	}

	// Create a new request using http
	req, err := http.NewRequestWithContext(ctx, method, u.String(), requestBody)

	if err != nil {
		return err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", "Bearer "+*c.accountApiToken)

	if requestBody != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	// Send req using http Client
	resp, err := c.client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		// TODO: https://developers.google.com/drive/api/guides/handle-errors
		return fmt.Errorf("HTTP Request got an unexpected statuscode %v", resp.StatusCode)
	}

	if responseData != nil {
		d := json.NewDecoder(resp.Body)
		d.DisallowUnknownFields()

		if err := d.Decode(&responseData); err != nil {
			return err
		}
	}

	return nil
}
