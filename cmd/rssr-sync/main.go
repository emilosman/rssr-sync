package main

import (
	"net/http"

	"github.com/emilosman/rssr-sync/internal/data"
	"github.com/emilosman/rssr-sync/internal/web"
)

func main() {
	ls := data.LoadLists()
	http.ListenAndServe(":8080", web.Server(ls))
}
