package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/labstack/gommon/log"
)

func do(baseUrl string, method string, params RequestParams, body json.RawMessage, headers RequestHeaders) (json.RawMessage, error) {
	if params == nil {
		params = make(map[string]string)
	}

	bodyReader := bytes.NewReader(body)
	queryParams := url.Values{}

	for k, v := range params {
		queryParams.Add(k, v)
	}

	var requestURL = baseUrl

	if len(queryParams) > 0 {
		requestURL = fmt.Sprintf("%s?%s", requestURL, queryParams.Encode())
	}

	log.Info("http request", "http", nil, "url", requestURL, "body", string(body), "method", method)

	client := &http.Client{}
	req, err := http.NewRequest(method, requestURL, bodyReader)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	// custom headers 
	for k,v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Error("request error", "http", err, "url", requestURL, "body", string(body), "method", method, "error", err)
		return nil, err
	}

	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)

	// Non successful responses are parsed as ErrorResponse
	if resp.StatusCode >= http.StatusBadRequest {
		log.Error("request error", "http", err, "response", string(b), "url", requestURL, "body", string(body), "method", method, "status_code", resp.StatusCode)
		return nil, newRequestError(resp.StatusCode, b)
	}

	return b, err
}