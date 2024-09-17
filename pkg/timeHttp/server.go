package timehttp

import (
	"fmt"
	"net/http"
	"time"
)

func GetTimeHandler() http.HandlerFunc {
	currentTime := time.Now().Format(time.RFC822)

	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/datetime" {
			http.NotFound(w, r)
			return
		}

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		fmt.Fprintf(w, "%s", currentTime)
	}
}

func StartServer(port string) error {
	http.HandleFunc("/datetime", GetTimeHandler())
	return http.ListenAndServe(":"+port, nil)
}
