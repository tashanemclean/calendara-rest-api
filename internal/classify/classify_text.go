package classify

import (
	"fmt"

	"github.com/tashanemclean/calendara-rest-api-api/internal/request"
	"github.com/tashanemclean/calendara-rest-api-api/util/config"
)

type RequestParams map[string]interface{}
type ClassificationResult struct {
	Activities interface{}      `json:"activities"`
	Location   string           `json:"location"`
	Duration   string           `json:"duration"`
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
	raw_result, err := request.Post[map[string]interface{}](url, params, headers)
	if err != nil {
		return nil, err
	}

	var result *ClassificationResult
	for key, value := range *raw_result {
		if key == "activities" {
			result = &ClassificationResult{
				Activities:  value,
			}
		}
		if key == "duration" {
			result = &ClassificationResult{
				Duration: fmt.Sprintf("%s",value),
			}
		}
		if key == "location" {
			result = &ClassificationResult{
				Location: fmt.Sprintf("%s",value),
			}
		}
	}

	return result, nil
}