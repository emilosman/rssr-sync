package web

import (
	"testing"
)

type mockLists struct {
	GetListFunc func(apiKey string, ts int64) ([]byte, error)
}

func TestServer(t *testing.T) {
	t.Run("requires apiKey", func(t *testing.T) {
	})
}
