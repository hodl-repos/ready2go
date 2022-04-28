package helper

import (
	"bytes"
	"encoding/json"
	"io"
)

func JsonToIoReader(data interface{}) (io.Reader, error) {
	if data == nil {
		return nil, nil
	}

	json_data, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(json_data)

	return buffer, nil
}
