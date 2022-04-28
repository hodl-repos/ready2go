package r2o

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func GetAccountAccessToken(developerToken string) error {
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
		AuthorizationCallbackUri: "https://example.com?auth=succeess",
	}

	json_data, err := json.Marshal(&request)

	if err != nil {
		fmt.Errorf("GetAccountAccessToken: cannot serialize request-dto")
		return err
	}

	// Create a new request using http
	req, err := http.NewRequest("POST", _BASE_API_URL_V1+"/developerToken/grantAccessToken", bytes.NewBuffer(json_data))

	if err != nil {
		fmt.Errorf("GetAccountAccessToken: cannot create http request")
		return err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", "Bearer "+developerToken)
	req.Header.Add("Content-Type", "application/json")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Errorf("GetAccountAccessToken: cannot send http request")
		return err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("GetAccountAccessToken: http status wrong")
		return errors.New("GetAccountAccessToken: wront status @ get request: " + resp.Status)
	}

	defer resp.Body.Close()

	response := responseDto{}

	d := json.NewDecoder(resp.Body)
	d.DisallowUnknownFields()

	if err := d.Decode(&response); err != nil {
		return err
	}

	return nil
}
