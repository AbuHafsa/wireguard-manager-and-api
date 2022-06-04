package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gitlab.com/raspberry.tech/wireguard-manager-and-api/src/api/middleware"
)

func TestAuthMiddleware(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/manager/key", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middleware.Respond(w)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := middleware.FailedAuthentication{
		StatusCode: 400,
		Message:    "Invalid authentication key",
	}

	var got middleware.FailedAuthentication
	err = json.Unmarshal(rr.Body.Bytes(), &got)
	if err != nil {
		t.Errorf("unable to parse body")
	}
	if got.StatusCode != expected.StatusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", got.StatusCode, expected.StatusCode)
	}
}
