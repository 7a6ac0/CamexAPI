package auth

import (
	u "CamexAPI/utils"
	"net/http"
)

// Error replies to the request with the specified error message and HTTP code.
// It does not otherwise end the request; the caller should ensure no further
// writes are done to w.
// The error message should be plain text.
func Error(w http.ResponseWriter, error string, code int) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	u.Respond(w, code, u.Message(false, error))
}

// NotFound replies to the request with an HTTP 404 not found error.
func NotFound(w http.ResponseWriter, r *http.Request) { Error(w, "404 page not found", http.StatusNotFound) }

// NotFoundHandler returns a simple request handler
// that replies to each request with a ``404 page not found'' reply.
func NotFoundHandler() http.Handler { return http.HandlerFunc(NotFound) }