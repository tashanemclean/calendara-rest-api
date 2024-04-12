package interactor

// Definte genric interactor type to improve resuability
type Interactor[T any] interface {
	Execute() T
}

type BaseResult struct {
	Success bool
	Errors  []error
}

// Base generic error handle accross all interactor instances
func (b *BaseResult) IsError() bool {
	return len(b.Errors) > 0
}

func (b *BaseResult) AsError() error {
	return b.Errors[0]
}