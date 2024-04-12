package request

import "encoding/json"

type RequestError struct {
	Message 		string `json:"message"`
	Code 			int `json:"code"`
	OriginalBytes 	[]byte `json:"original_bytes"`
}

func newRequestError(code int, b []byte) *RequestError {
	var err *RequestError
	json.Unmarshal(b, &err)
	err.OriginalBytes = b
	err.Code = code
	return err
}

func (err *RequestError) Error() string {
	if err.Message != "" {
		return err.Message
	}

	return string(err.OriginalBytes)
}