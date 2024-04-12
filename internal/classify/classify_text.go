package classify

import (
	"fmt"
	"strconv"

	"github.com/tashanemclean/genai-rest-api/internal/request"
	"github.com/tashanemclean/genai-rest-api/util/config"
)

type RequestParams map[string]interface{}

type ClassificationResult struct {
	
}

func ClassifyText(classifyText string) (*ClassificationResult,error) {
	headers := map[string]string{
		"Content-Type": "appliccation/json",
	}

	// ensure arg passes is text type
	text, err := strconv.Atoi(classifyText)
	if err != nil {
		return nil,err
	}

	params := RequestParams{
		"text": text,
	}

	url := fmt.Sprintf("%s/api/text", config.Config.ApiBaseUrl)
	result, err := request.Post[ClassificationResult](url, params, headers)

	if err != nil || result == nil {
		return nil, err
	}

	return result, nil
}