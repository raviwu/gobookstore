package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raviwu/gobookstore/models"
)

var (
	db = models.SetupModels()
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestFindBooks(t *testing.T) {
	body := `{"data":[]}`

	router := SetupRouter()

	w := performRequest(router, "GET", "/books")

	if w.Code != http.StatusOK {
		t.Error("failing request")
	}

	if w.Body.String() != body {
		t.Errorf("expected: %s\nactual: %s", body, w.Body)
	}
}
