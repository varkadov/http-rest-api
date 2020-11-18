package apiserver

import "net/http"

type responseWrite struct {
	// NOTE !!!
	http.ResponseWriter
	code int
}

// NOTE !!!
func (w *responseWrite) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
