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

	var result ClassificationResult
	values := reflect.ValueOf(*raw_result)
	if values.Kind().String() == "map" {
		iter := getIter(*raw_result)
		result = toResultsStruct(iter)
	}


	return result, nil
}

func getIter(data interface{}) *reflect.MapIter {
	iter := reflect.ValueOf(data).MapRange()
	return iter
}

func toResultsStruct(iter *reflect.MapIter ) ClassificationResult {
	var result ClassificationResult
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		if k.String() == "activities" {
			result.Activities = v.Interface()
		}
		if k.String() == "duration" {
			result.Duration = fmt.Sprintf("%v", v.Interface())
		}
		if k.String() == "location" {
			result.Location = fmt.Sprintf("%v", v.Interface())
		}
	}

	return result
}