package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yasushiasahi/karabiner-member/server/api"
	"github.com/yasushiasahi/karabiner-member/server/data"
)

func TestGetMenbers(t *testing.T) {
	boss := data.Member{
		Img:   "https://www.karabiner.tech/asset/images/common/member_uz.jpg",
		Title: "代表取締役",
		Name:  "ゆーじ",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/member/", api.HandleMember())

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/member/get", nil)
	mux.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}

	var ms []data.Member
	json.Unmarshal(w.Body.Bytes(), &ms)

	if ms[0].Img != boss.Img {
		t.Errorf("boss Img is not correct. got %v", ms[0].Img)
	}

	if ms[0].Title != boss.Title {
		t.Errorf("boss Title is not correct. got %v", ms[0].Title)
	}

	if ms[0].Name != boss.Name {
		t.Errorf("boss Name is not correct. got %v", ms[0].Name)
	}

}
