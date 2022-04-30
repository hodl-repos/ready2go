package r2o

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hodl-repos/ready2go/helper"
)

const (
	_BASE_API_URL_V1 = "https://api.ready2order.com/v1/"
)

type service struct {
	client *Client
}

// A Client manages communication with the ready2order API.
type Client struct {
	client  *http.Client // HTTP client used to communicate with the API.
	baseURL *url.URL

	accountApiToken *string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Ready2Order API.
	Account                 *AccountService
	AccountingFinancialYear *AccountingFinancialYearService
	Webhook                 *WebhookService
}

// NewClient returns a new Ready2Order API client. If a nil httpClient is
// provided, a new http.Client will be used.
func NewClient(accountApiToken *string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(_BASE_API_URL_V1)

	c := &Client{client: httpClient, accountApiToken: accountApiToken, baseURL: baseURL}
	c.common.client = c

	c.Account = (*AccountService)(&c.common)
	c.AccountingFinancialYear = (*AccountingFinancialYearService)(&c.common)
	c.Webhook = (*WebhookService)(&c.common)

	return c
}

func (c *Client) runHttpRequest(path, method string, requestData interface{}, responseData interface{}) error {
	return c.runHttpRequestWithContext(context.Background(), path, method, requestData, responseData)
}

func (c *Client) runHttpRequestWithContext(ctx context.Context, path, method string, requestData interface{}, responseData interface{}) error {
	//URL BUILDING
	apiUrl, err := helper.BuildApiUrl(c.baseURL, &path)
	if err != nil {
		return err
	}

	//REQUEST BODY
	requestBody, err := helper.JsonToIoReader(requestData)
	if err != nil {
		return err
	}

	//CREATE A NEW HTTP REQUEST
	req, err := http.NewRequestWithContext(ctx, method, *apiUrl, requestBody)

	if err != nil {
		return err
	}

	//ADD AUTHORIZATION HEADER
	req.Header.Add("Authorization", "Bearer "+*c.accountApiToken)

	//ADD CONTENT-TYPE IF CONTENT IS PRESENT
	if requestBody != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	//SEND REQUEST TO API
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		//try parsing the error -> is not nil if the error can be parsed
		if rateLimitError := helper.ParseRateLimitExceededError(&resp.Body); rateLimitError != nil {
			return rateLimitError
		}

		// TODO: https://developers.google.com/drive/api/guides/handle-errors
		return fmt.Errorf("HTTP Request got an unexpected statuscode %v", resp.StatusCode)
	}

	//DECODE RESPONSE IF THE RESPONSE IS WANTED
	helper.DecodeHttpResponse(resp, responseData)

	return nil
}
