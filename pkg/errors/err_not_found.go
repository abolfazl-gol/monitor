package errors

import "fmt"

type ErrNotFound struct {
	Resource string
	Column   string
	Value    interface{}
}

func (err ErrNotFound) Error() string {
	return fmt.Sprintf("%s{%s: %v} not_found", err.Resource, err.Column, err.Value)
}

func IsNotFound(err error) bool {
	_, ok := err.(ErrNotFound)
	return ok
}
