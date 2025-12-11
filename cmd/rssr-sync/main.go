package main

import (
	"net/http"

	"github.com/emilosman/rssr-sync/internal/data"
	"github.com/emilosman/rssr-sync/internal/web"
)

func main() {
	ls := &data.Lists{ListIndex: make(map[string]*data.List)}
	http.ListenAndServe(":8080", web.Server(ls))
}
