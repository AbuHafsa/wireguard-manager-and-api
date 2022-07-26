package middleware

import "net/http"

func SetCORSHeaders(res *http.ResponseWriter) {
	header := (*res).Header()

	header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	header.Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	header.Set("Access-Control-Allow-Origin", "*")
}
