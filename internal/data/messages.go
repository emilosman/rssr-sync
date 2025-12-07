package data

import "errors"

var (
	ErrOldTimestamp = errors.New("Timestamp is old")
	ErrListNotFound = errors.New("List not found")
	ErrNoApiKey     = errors.New("No API key")
)
