package web

import "errors"

var (
	ErrMethodNotAllowed = errors.New("method not allowed")
	ErrInvalidJsonBody  = errors.New("invalid JSON body")
	ErrNoApiKey         = errors.New("apiKey required")
)
