package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/viper"
	"gitlab.com/raspberry.tech/wireguard-manager-and-api/src/api/middleware"
)

func setUpConfigurationFile() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/opt/wgManagerAPI/")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic("Unable to read configuration file")
	}
}

func Test_Returns_BadRequest_When_Invalid_Authentication_Key(t *testing.T) {
	setUpConfigurationFile()
	viper.Set("SERVER.AUTH", "eyJhbGciOiJIUzI1NiJ9")

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middleware.VerifyAuthorizationHeader(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := middleware.FailedAuthentication{
		StatusCode: 400,
		Message:    "Invalid authentication key",
	}

	var got middleware.FailedAuthentication
	err = json.Unmarshal(rr.Body.Bytes(), &got)
	if err != nil {
		t.Errorf("unable to parse body %v", err)
	}
	if got.StatusCode != expected.StatusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", got.StatusCode, expected.StatusCode)
	}
}

func Test_Returns_StatusOK_When_Valid_Authentication_Key(t *testing.T) {
	setUpConfigurationFile()
	viper.Set("SERVER.AUTH", "eyJhbGciOiJIUzI1NiJ9")

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middleware.VerifyAuthorizationHeader(w, r)
	})

	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiJ9")
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
