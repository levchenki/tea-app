package admin

import "net/http"

func AdminOnly(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// todo validate admin
		isAdmin := true
		if !isAdmin {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
