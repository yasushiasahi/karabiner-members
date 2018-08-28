package main

import (
	"fmt"
	"net/http"

	"github.com/yasushiasahi/karabiner-member/server/api"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:3000",
	}

	http.HandleFunc("/api/member/", api.HandleMember())

	fmt.Println("server lisning at localhost:3000")
	server.ListenAndServe()
}
