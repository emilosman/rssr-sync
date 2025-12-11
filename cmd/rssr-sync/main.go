package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emilosman/rssr-sync/internal/data"
	"github.com/emilosman/rssr-sync/internal/web"
)

func main() {
	lists := &data.Lists{ListIndex: make(map[string]*data.List)}
	lists.ListIndex["abc"] = &data.List{
		Ts:     10,
		ApiKey: "abc",
		Data:   []byte(`"message": "hello`),
	}

	server := web.Server(lists)

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", server))

}
