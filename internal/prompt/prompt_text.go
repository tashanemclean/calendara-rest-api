package prompt

import (
	"fmt"

	"github.com/tashanemclean/calendara-rest-api-api/internal/request"
	"github.com/tashanemclean/calendara-rest-api-api/util/config"
)

type RequestParams map[string]interface{}
type ClassificationResult struct {
	Activities interface{} `json:"activities"`
}

func PromptText(input string) (*ClassificationResult,error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	params := RequestParams{
		"text": input,
	}

	url := fmt.Sprintf("%s/api/text", config.Config.ApiBaseUrl)
	result, err := request.Post[ClassificationResult](url, params, headers)
	if err != nil {
		return nil, err
	}


	return result, nil
}