package data

import "errors"

var (
	ErrOldTimestamp = errors.New("timestamp is old")
	ErrListNotFound = errors.New("list not found")
	ErrNoApiKey     = errors.New("no API key")
)
