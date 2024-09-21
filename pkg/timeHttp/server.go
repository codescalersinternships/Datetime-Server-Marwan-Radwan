package timehttp

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

var server *http.Server

// GetTimeHandler returns an HTTP handler function that responds with the current time formatted depending on accept header.
func GetTimeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		acceptHeader := r.Header.Get("Accept")
		if acceptHeader == "application/json" {
			getTimeJsonHandler()(w, r)
		} else {
			getTimePlainHandler()(w, r)
		}
	}
}

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

// StartServer starts an HTTP server on the specified port.
func StartServer(port string) error {
	server := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	http.HandleFunc("/datetime", getTimeJsonHandler())

	return server.ListenAndServe()
}

// ShutdownServer gracefully shuts down the HTTP server
func ShutdownServer(ctx context.Context) error {
	if server != nil {
		return server.Shutdown(ctx)
	}
	return nil
}
