package request

import (
	"encoding/json"
	"net/http"
)

func Get[T any](requestUrl string, params RequestParams, headers RequestHeaders) (*T, error) {
	resp, err := do(requestUrl, http.MethodGet, params, nil, headers)

	if err != nil {
		return nil, err
	}

	var t *T

	err = json.Unmarshal(resp, &t)
	return t, err
}