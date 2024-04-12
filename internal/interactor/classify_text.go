package interactor

import (
	"log"

	"github.com/tashanemclean/genai-rest-api/internal/classify"
)

type ClassifyTextArgs struct {
	Text string
}

type ClassifyTextResult struct {
	*BaseResult
	ClassifyResult any
}

type classifyText struct {
	text string
	errors []error
	data any
	// TODO define custom Manager for text classfication
	// classifyTextManager *classifyTextManager

}

func (ia *classifyText) Execute() ClassifyTextResult {
	result, err := classify.ClassifyText(ia.text)
	if err != nil {
		log.Fatal("Error classfying text", err)
		return ia.fail(err)
	}
	ia.data = result
	return ia.makeResult()
}

func (ia *classifyText) fail(err error) ClassifyTextResult {
	ia.errors = append(ia.errors, err)
	return ia.makeResult()
}

func (ia *classifyText) makeResult() ClassifyTextResult {
	return ClassifyTextResult{
		ClassifyResult: ia.data,
		BaseResult: &BaseResult{
			Errors: ia.errors,
			Success: len(ia.errors) == 0,
		},
	}
}

func ClassifyText(args ClassifyTextArgs) Interactor[ClassifyTextResult] {
	return &classifyText{
		text: args.Text,
	}
}