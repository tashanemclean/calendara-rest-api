package classify

import (
	"fmt"

	"github.com/tashanemclean/calendara-rest-api-api/internal/request"
	"github.com/tashanemclean/calendara-rest-api-api/util/config"
)

type RequestParams map[string]interface{}
type ClassificationResult struct {
	Activities interface{}      `json:"activities"`
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
	raw_result, err := request.Post[interface{}](url, params, headers)
	if err != nil {
		return nil, err
	}

	fmt.Println(*&raw_result, "** raw result")

	result := &ClassificationResult{
		Activities: []string{"some"},
		Duration: "",
		Location: "",
	}

	return result, nil
}