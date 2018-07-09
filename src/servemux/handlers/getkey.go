package handlers

import (
	"fmt"
	"goSamples/servemux/storage"
	"net/http"
)

func GetKey(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Missing key name in query string", http.StatusBadRequest)
			return
		}
		val, err := db.Get(key)
		if err == storage.ErrNotFound {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, fmt.Sprintf("Error getting value from database: %s", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(val)
	})
}