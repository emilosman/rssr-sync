package data

import "errors"

var (
	ErrOldTimestamp              = errors.New("timestamp is old")
	ErrdClientOldTimestampUpdate = errors.New("client timestamp is old, providing list")
	ErrListNotFound              = errors.New("list not found")
	ErrNoApiKey                  = errors.New("apiKey required")
)
