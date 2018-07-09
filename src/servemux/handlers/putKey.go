package handlers

import (
	"io/ioutil"
	"net/http"

	"goSamples/servemux/storage"
)

func PutKey(db storage.DB) http.Handler {
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

		err = db.Set(key, val)
		if err != nil {
			http.Error(w, "Error setting value in DB", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	})
}
