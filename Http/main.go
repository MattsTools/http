package Http

import (
	"bytes"
	"encoding/json"
	"github.com/mattstools/weberrors/WebErrors"
	"net/http"
)

func GetRequest(url string, authHeader string, authToken string, marshalTo interface{}) (interface{}, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(authHeader, authToken)

	resp, respErr := client.Do(req)

	if respErr != nil {
		return nil, respErr
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return nil, WebErrors.WebError{
			StatusCode:      resp.StatusCode,
			ErrorMessage:    resp.Status,
			ErrorIdentifier: "EXTCALL",
		}
	}

	jsonDecodeErr := json.NewDecoder(resp.Body).Decode(marshalTo)

	return marshalTo, jsonDecodeErr
}

func PostRequest(url string, authHeader string, authToken string, body interface{}, marshalTo interface{}) (interface{}, error) {
	marshalled, marshalErr := json.Marshal(body)
	if marshalErr != nil {
		return nil, marshalErr
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshalled))
	if err != nil {
		return nil, err
	}

	if authHeader != "" && authToken != "" {
		req.Header.Add(authHeader, authToken)
	}
	resp, respErr := client.Do(req)

	if respErr != nil {
		return nil, respErr
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return nil, WebErrors.WebError{
			StatusCode:      resp.StatusCode,
			ErrorMessage:    resp.Status,
			ErrorIdentifier: "EXTCALL",
		}
	}

	jsonDecodeErr := json.NewDecoder(resp.Body).Decode(marshalTo)

	return marshalTo, jsonDecodeErr
}
