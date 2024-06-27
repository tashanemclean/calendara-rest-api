package args

type PromptText struct {
	Activity    []string       `body:"activity"`
  Categories  []string       `body:"categories"`
  City        string         `body:"city"`
  State       string         `body:"state"`
  Days       int         `body:"days"`
}