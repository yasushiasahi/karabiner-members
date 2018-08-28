package main

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/yasushiasahi/karabiner-member/server/api"
// 	"github.com/yasushiasahi/karabiner-member/server/data"
// )

// func TestCreateUser(t *testing.T) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/api/member/", api.HandleMember())

// 	jsonBody := strings.NewReader(`{"ID":"","Name":"テスト太郎","PW":"password","Token": ""}`
// 	)

// 	w := httptest.NewRecorder()
// 	r, _ := http.NewRequest("POST", "/api/user/create/", jsonBody)
// 	mux.ServeHTTP(w, r)

// 	if w.Code != 200 {
// 		t.Errorf("Response code is %v", w.Code)
// 	}

// 	var u data.User
// 	json.Unmarshal(w.Body.Bytes(), u)

// 	if u.Name != "テスト太郎" {
// 		t.Errorf("u.Name shuld be 'テスト太郎'. got %v", u.Name)
// 	}

// 	if u.PW != data.MakeHash("password") {
// 		t.Errorf("u.Token is not corrre. got %v", u.PW)
// 	}

// }
