package middleware

import (
	"net/http"
)

func EnableCORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		res.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		res.Header().Set("Access-Control-Allow-Origin", "*")

		if req.Method == "OPTIONS" {
			res.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(res, req)
	})
}
