package r2o

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

const (
	_TEST_ACCOUNT_TOKEN = ""
)

func TestSomething(t *testing.T) {
	// Create a new request using http
	req, err := http.NewRequest("GET", _BASE_API_URL_V1+"/company", nil)

	if err != nil {
		fmt.Errorf("GetAccountAccessToken: cannot create http request")
	}

	// add authorization header to the req
	req.Header.Add("Authorization", "Bearer "+_TEST_ACCOUNT_TOKEN)
	req.Header.Add("Content-Type", "application/json")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Errorf("GetAccountAccessToken: cannot send http request")
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("GetAccountAccessToken: http status wrong")
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}

	data := string(b)

	_ = data
}
