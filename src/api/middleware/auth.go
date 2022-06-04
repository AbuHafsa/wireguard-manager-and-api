package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
)

type FailedAuthentication struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func VerifyAuthorizationHeader(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	auth := viper.GetString("SERVER.AUTH")
	if auth != "" { //check AUTH
		authHeader := req.Header.Get("Authorization")
		if auth != authHeader {
			res.WriteHeader(http.StatusBadRequest)
			payload, _ := json.Marshal(FailedAuthentication{
				StatusCode: 400,
				Message:    "Invalid authentication key",
			})
			res.Write(payload)
			return
		}
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		VerifyAuthorizationHeader(res, req)
		next.ServeHTTP(res, req)
	})
}
