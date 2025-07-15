package middleware

import "net/http"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if !ok {
			w.Write([]byte("Username or password cannot be empty"))
			return
		}

		if username == "admin" && password == "admin" {
			next.ServeHTTP(w, r)
			return
		} else {
			w.Write([]byte("Invalid username or password"))
			return
		}
	})
}
