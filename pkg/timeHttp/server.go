package timehttp

import (
	"fmt"
	"net/http"
	"time"
)

// getTimeJsonHandler returns an HTTP handler function that responds with the current time in json format.
func getTimeJsonHandler() http.HandlerFunc {
	currentTime := time.Now().Format(time.RFC822)

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"datetime": "%s"}`, currentTime)
	}
}

// getTimePlainHandler returns an HTTP handler function that responds with the current time in plain text.
func getTimePlainHandler() http.HandlerFunc {
	currentTime := time.Now().Format(time.RFC822)

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		fmt.Fprintf(w, "%s", currentTime)
	}
}

func StartServer(port string) error {

	http.HandleFunc("/datetime/plain", getTimePlainHandler())
	http.HandleFunc("/datetime/json", getTimeJsonHandler())

	return http.ListenAndServe(":"+port, nil)
}
