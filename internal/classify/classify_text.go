package classify

import (
	"fmt"
	"reflect"

	"github.com/tashanemclean/calendara-rest-api-api/internal/request"
	"github.com/tashanemclean/calendara-rest-api-api/util/config"
)

type RequestParams map[string]interface{}
type ClassificationResult struct {
	Activities interface{}           `json:"activities"`
	Location   string          `json:"location"`
	Duration   string          `json:"duration"`
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

	values := reflect.ValueOf(*raw_result)
	keys := getMapKeys(values)
	result := toResultsStruct(values, keys)

	// r := &ClassificationResult{
	// 	Activities: []string{"some"},
	// 	Duration: "",
	// 	Location: "",
	// }

	return result, nil
}

func getMapKeys(values reflect.Value)[]reflect.Value {
	var mapKeys []reflect.Value
	if values.Kind().String() == "map" {
		mapKeys = append(mapKeys, values.MapKeys()...)
	}
	return mapKeys
}

func toResultsStruct(values reflect.Value,keys []reflect.Value) ClassificationResult {
	var result ClassificationResult
	for _, key := range keys {
		if key.String() == "activities" {
			value := values.MapIndex(key)
			result = ClassificationResult{
				Activities:  value.Interface(),
			}
		}
		if key.String() == "duration" {
			value := values.MapIndex(key)
			result = ClassificationResult{
				Duration: fmt.Sprintf("%v",value.Interface()),
			}
		}
		if key.String() == "location" {
			value := values.MapIndex(key)
			result = ClassificationResult{
				Location: fmt.Sprintf("%v",value.Interface()),
			}
		}
	}
	return result
}