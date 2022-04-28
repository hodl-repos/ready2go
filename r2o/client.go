package r2o

const (
	_BASE_API_URL_V1 = "https://api.ready2order.com/v1"
)

type service struct {
	client *Client
}

// A Client manages communication with the ready2order API.
type Client struct {
	client   *http.Client // HTTP client used to communicate with the API.

	accountApiToken *string

	// Base URL for API requests. Defaults to the public GitHub API, but can be
	// set to a domain endpoint to use with GitHub Enterprise. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the GitHub API.
	Account        *ActionsService
}

type Ready2GoClient struct {
	accountApiToken *string
}

func NewService(accountApiToken string) Ready2GoClient {
	return Ready2GoClient{
		accountApiToken: &accountApiToken,
	}
}

// NewClient returns a new GitHub API client. If a nil httpClient is
// provided, a new http.Client will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(_BASE_API_URL_V1)

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.common.client = c
	c.Account = (*ActionsService)(&c.common)
	return c
}