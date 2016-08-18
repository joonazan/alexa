package alexa

import (
	"encoding/json"
	"net/http"
)

func HandleFunc(path string, f func(Request) Response) {
	http.HandleFunc(path, Handler(f))
}

func Handler(f func(Request) Response) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		var request Request
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(struct {
			Version  string   `json:"version"`
			Response Response `json:"response"`
		}{
			Version:  "1.0",
			Response: f(request),
		})

	}
}
