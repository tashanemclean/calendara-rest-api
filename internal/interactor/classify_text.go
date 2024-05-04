package interactor

import (
	"fmt"
	"log"
	"strconv"

	"github.com/tashanemclean/calendara-rest-api-api/internal/classify"
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
	Activities any
	Location string
	Duration string
}
type ClassifyTextArgs struct {
	Activity    []string       `body:"activity"`
  Categories  []string       `body:"categories"`
  City        string         `body:"city"`
  State       string         `body:"state"`
  Days        int            `body:"days"`
}

type ClassifyTextResult struct {
	*BaseResult
	*ClassificationResult
}

type classifyText struct {
	textPrompt string
	rawDataArgs ClassifyTextArgs
	errors []error
	data ClassificationResult
	// TODO define custom Manager for text classfication
	// classifyTextManager *classifyTextManager

}

func (ia *classifyText) Execute() ClassifyTextResult {
	ia.prepareData()
	result, err := classify.ClassifyText(ia.textPrompt)
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
// 	Loop activities, if payload activity predicate,
// Then set string equal to activity of index
// Loop categories, if payload category predicate, then set string equal to categories of index
func (ia *classifyText) prepareData() {
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
		rawDataArgs: args,
	}
}