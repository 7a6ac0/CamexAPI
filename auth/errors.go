package auth

import (
	u "cashgone/utils"
	"net/http"
)

var NotFoundHandler = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u.Respond(w, http.StatusNotFound, u.Message(false, "This resources was not found on our server"))
		next.ServeHTTP(w, r)
	})
}
