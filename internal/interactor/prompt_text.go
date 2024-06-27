package interactor

import (
	"fmt"
	"strconv"

	"github.com/tashanemclean/calendara-rest-api-api/internal/prompt"
)

var activitiesIndex = []BaseArgs{
	{Id: "0", Name: "free"},
	{Id: "1", Name: "paid"},
}
var categoriesIndex = []BaseArgs{
	{Id: "0", Name: "hiking"},
	{Id: "1", Name: "fishing"},
	{Id: "2", Name: "basketball"},
	{Id: "3", Name: "swimming"},
	{Id: "4", Name: "walking"},
	{Id: "5", Name: "aquarium"},
	{Id: "6", Name: "other"},
}

type BaseArgs struct {
	Id string
	Name string
}

type ClassificationResult struct {
	Activities interface{} `json:"Activities"`
}
type PromptTextArgs struct {
	Activity    []string       `body:"activity"`
  Categories  []string       `body:"categories"`
  City        string         `body:"city"`
  State       string         `body:"state"`
  Days        int            `body:"days"`
}

type PromptTextResult struct {
	*BaseResult
	ClassificationResult interface{}
}

type promptText struct {
	textPrompt string
	rawDataArgs PromptTextArgs
	errors []error
	data interface{}
	// TODO define custom Manager for text classfication
	// classifyTextManager *classifyTextManager

}

func (ia *promptText) Execute() PromptTextResult {
	ia.prepareData()
	result, err := prompt.PromptText(ia.textPrompt)
	if err != nil {
		return ia.fail(err)
	}
	ia.data = result
	return ia.makeResult()
}

// Loop activities, if payload activity predicate,
// Then set string equal to activity of index
// Loop categories, if payload category predicate, then set string equal to categories of index
func (ia *promptText) prepareData() {
	activitiesData := ia.rawDataArgs.Activity
	var activities string
	for idx, elem := range activitiesData {
		if elem == activitiesIndex[idx].Id {
			txt := activitiesIndex[idx].Name
			activities = fmt.Sprintf("%s, %s", activities, txt)
		}
	}

	categoriesData := ia.rawDataArgs.Categories
	var categories string
	for idx, elem := range categoriesData {
		if elem == categoriesIndex[idx].Id {
			txt := categoriesIndex[idx].Name
			categories = fmt.Sprintf("%s, %s", categories, txt)
		}
	}
	
	// Find me 15 activities (days) free(activity type) hiking, walking, aquarium (categories), activities in Norwalk CT for {number of days} days)
	city := ia.rawDataArgs.City
	state := ia.rawDataArgs.State
	days := strconv.Itoa(ia.rawDataArgs.Days)
	
	ia.textPrompt = fmt.Sprintf("Find me%v %v, activities in %s %s for %s days", activities, categories, city, state, days)
}
func (ia *promptText) fail(err error) PromptTextResult {
	ia.errors = append(ia.errors, err)
	return ia.makeResult()
}

func (ia *promptText) makeResult() PromptTextResult {
	return PromptTextResult{
		ClassificationResult: ia.data,
		BaseResult: &BaseResult{
			Errors: ia.errors,
			Success: len(ia.errors) == 0,
		},
	}
}

func PromptText(args PromptTextArgs) Interactor[PromptTextResult] {
	return &promptText{
		rawDataArgs: args,
	}
}