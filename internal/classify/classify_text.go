package classify

import (
	"fmt"

	"github.com/tashanemclean/calendara-rest-api-api/internal/request"
	"github.com/tashanemclean/calendara-rest-api-api/util/config"
)

type RequestParams map[string]interface{}
type ClassificationResult struct {
	Activities any
	Location   string   `json:"location"`
	Duration   string   `json:"duration"`
}

func ClassifyText(classifyText string) (*ClassificationResult,error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	// TODO ensure type is string

	params := RequestParams{
		"text": classifyText,
	}

	url := fmt.Sprintf("%s/api/text", config.Config.ApiBaseUrl)
	result, err := request.Post[ClassificationResult](url, params, headers)
	if err != nil || result == nil {
		return nil, err
	}

	return result, nil
}