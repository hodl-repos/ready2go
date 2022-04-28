package r2o

const (
	_BASE_API_URL_V1 = "https://api.ready2order.com/v1"
)

type Ready2GoClient struct {
	accountApiToken *string
}

func NewService(accountApiToken string) Ready2GoClient {
	return Ready2GoClient{
		accountApiToken: &accountApiToken,
	}
}
