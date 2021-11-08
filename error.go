package gomath

import "errors"

var (
	// ErrInvalidInput is returned when the input is invalid
	ErrValueType = errors.New("value type error")
	// ErrOutOfIndex is returned when the index is out of range
	ErrOutOfIndex = errors.New("index out of range")
)
