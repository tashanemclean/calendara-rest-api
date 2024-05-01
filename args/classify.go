package args

type ClassifyText struct {
	Activity    []string       `body:"activity"`
  Categories  []string       `body:"categories"`
  City        string         `body:"city"`
  State       string         `body:"state"`
  Days       int         `body:"days"`
}