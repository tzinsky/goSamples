package handlers

import (
	"net/http"
)

func getKey() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Missing key name in query string", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
