package god

import "net/http"

func writeContentType(contentType string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", contentType)
}

// func writeStatusCode(code int, w http.ResponseWriter) {
// 	w.WriteHeader(code)
// }
