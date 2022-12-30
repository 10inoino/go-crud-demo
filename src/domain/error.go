package domain

import "errors"

type InvalidInputError struct {
	error
}

func NewInvalidInputError(msg string) error {
	return &InvalidInputError{
		errors.New(msg),
	}
}

type NotFoundError struct {
	error
}

func IsNotFoundError(err error) bool {
	_, ok := err.(*NotFoundError)
	return ok
}

func NewNotFoundError(msg string) error {
	return &NotFoundError{
		errors.New(msg),
	}
}

type DuplicateError struct {
	error
}

func IsDuplicateError(err error) bool {
	_, ok := err.(*DuplicateError)
	return ok
}

func NewDuplicateError(msg string) error {
	return &DuplicateError{
		errors.New(msg),
	}
}
