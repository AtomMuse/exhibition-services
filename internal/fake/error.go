package fake

import "errors"

var (
	ErrSomethingWentWrong error = errors.New("something went wrong")
)
