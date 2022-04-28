package helper

import (
	"encoding/json"
	"net/http"
)

func DecodeHttpResponse(response *http.Response, data interface{}) error {
	if data == nil {
		return nil
	}

	d := json.NewDecoder(response.Body)
	d.DisallowUnknownFields()

	if err := d.Decode(&data); err != nil {
		return err
	}

	return nil
}
