package request

import (
	"encoding/json"
	"net/http"
)

func Post[T any](url string, obj any, headers RequestHeaders) (*T, error) {
	b, err := json.Marshal(obj)

	if err != nil {
		return nil, err
	}

	reqParams := map[string]string{}
	resp, err := do(url, http.MethodPost, reqParams, b, headers)

	if resp == nil || err != nil {
		return nil, err
	}

	var results map[string]interface{}

	var t *T
	err = json.Unmarshal(resp, &results)

	if err != nil {
		return nil, err
	}

	return t, nil
}