package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/yasushiasahi/karabiner-member/server/data"
)

// HandleMember returns HandleFunc for Member
func HandleMember() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		action := strings.Split(r.URL.Path, "/")[3]
		fmt.Println(action)

		switch action {
		case "get":
			getMembers(w, r)
		default:
			http.NotFound(w, r)
		}
	}
}

func getMembers(w http.ResponseWriter, r *http.Request) {
	ms, err := data.Scrape()
	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(ms)
	if err != nil {
		panic(err)
	}
}
