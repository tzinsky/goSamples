package handlers

import (
	"io/ioutil"
	"net/http"
)

func putKey() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Missing key name in query string", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		val, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading PUT body", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)

	})
}
