package authorization

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hodl-repos/ready2go/helper"
)

const (
	_BASE_API_URL_V1 = "https://api.ready2order.com/v1"
)

//returns the uri for the creation of a account-token
func GetAccountAccessToken(developerToken string) (*string, error) {
	type requestDto struct {
		AuthorizationCallbackUri string `json:"authorizationCallbackUri"`
	}

	type responseDto struct {
		Error            *bool   `json:"error"`
		Success          *bool   `json:"success"`
		ExpiresAt        *int64  `json:"expiresAt"`
		GrantAccessUri   *string `json:"grantAccessUri"`
		GrantAccessToken *string `json:"grantAccessToken"`
		Status           *string `json:"status"`
	}

	request := requestDto{
		AuthorizationCallbackUri: "https://mauracher.cc/r2o",
	}

	requestBody, err := helper.JsonToIoReader(&request)

	if err != nil {
		fmt.Errorf("GetAccountAccessToken: cannot serialize request-dto")
		return nil, err
	}

	// Create a new request using http
	req, err := http.NewRequest("POST", _BASE_API_URL_V1+"/developerToken/grantAccessToken", requestBody)

	if err != nil {
		fmt.Errorf("GetAccountAccessToken: cannot create http request")
		return nil, err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", "Bearer "+developerToken)
	req.Header.Add("Content-Type", "application/json")

	// Send req using http Client
	client := &http.Client{}
	httpResp, err := client.Do(req)

	if err != nil {
		fmt.Errorf("GetAccountAccessToken: cannot send http request")
		return nil, err
	}

	if httpResp.StatusCode != http.StatusOK {
		fmt.Errorf("GetAccountAccessToken: http status wrong")
		return nil, errors.New("GetAccountAccessToken: wront status @ get request: " + httpResp.Status)
	}

	defer httpResp.Body.Close()

	response := responseDto{}

	helper.DecodeHttpResponse(httpResp, &response)

	return response.GrantAccessUri, nil
}
