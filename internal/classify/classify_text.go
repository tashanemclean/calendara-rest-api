package classify

import (
	"fmt"
	"reflect"

	"github.com/tashanemclean/calendara-rest-api-api/internal/request"
	"github.com/tashanemclean/calendara-rest-api-api/util/config"
)

type RequestParams map[string]interface{}
type ClassificationResult struct {
	Activities any      `json:"activities"`
	Location   string           `json:"location"`
	Duration   string           `json:"duration"`
}

func ClassifyText(classifyText string) (ClassificationResult,error) {
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
		return ClassificationResult{}, err
	}

	var result ClassificationResult
	values := reflect.ValueOf(*raw_result)
	if values.Kind().String() == "map" {
		for _, key := range values.MapKeys() {
			value := values.MapIndex(key)
			if key.String() == "activities" {
				result = ClassificationResult{
					Activities:  value.Interface(),
				}
			}
			if key.String() == "duration" {
				result = ClassificationResult{
					Duration: value.String(),
				}
			}
			if key.String() == "location" {
				result = ClassificationResult{
					Location: value.String(),
				}
			}
	
		}
	}

	// r := &ClassificationResult{
	// 	Activities: []string{"some"},
	// 	Duration: "",
	// 	Location: "",
	// }

	return result, nil
}