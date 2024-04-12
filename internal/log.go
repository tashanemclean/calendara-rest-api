package internal

// TODO setup internal stackframe logger
type stackFrame struct {
	Func   string `json:"func"`
	Source string `json:"source"`
	Line   int    `json:"line"`
}