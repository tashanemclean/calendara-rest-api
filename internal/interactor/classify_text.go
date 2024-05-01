package interactor

import (
	"log"

	"github.com/tashanemclean/calendara-rest-api-api/internal/classify"
)

type ClassificationResult struct {
	Activities any
	Location string
	Duration string
}
type ClassifyTextArgs struct {
	Text string
}

type ClassifyTextResult struct {
	*BaseResult
	*ClassificationResult
}

type classifyText struct {
	text string
	errors []error
	data ClassificationResult
	// TODO define custom Manager for text classfication
	// classifyTextManager *classifyTextManager

}

func (ia *classifyText) Execute() ClassifyTextResult {
	result, err := classify.ClassifyText(ia.text)
	if err != nil {
		log.Fatal("Error classfying text ", err)
		return ia.fail(err)
	}
	ia.data = ClassificationResult{
		Activities: result.Activities,
		Duration: result.Duration,
		Location: result.Location,
	}
	return ia.makeResult()
}

func (ia *classifyText) fail(err error) ClassifyTextResult {
	ia.errors = append(ia.errors, err)
	return ia.makeResult()
}

func (ia *classifyText) makeResult() ClassifyTextResult {
	return ClassifyTextResult{
		ClassificationResult: &ClassificationResult{
			Activities: ia.data.Activities,
			Duration: ia.data.Duration , 
			Location: ia.data.Location,
		},
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