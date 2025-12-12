package main

import (
	"fmt"
	"net/http"

	"github.com/emilosman/rssr-sync/internal/data"
	"github.com/emilosman/rssr-sync/internal/web"
)

func main() {
	ls := data.Load()
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", web.Server(ls))
}
