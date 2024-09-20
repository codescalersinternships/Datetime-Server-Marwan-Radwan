package timehttp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetTimePlainHandler(t *testing.T) {
	handler := getTimePlainHandler()

	t.Run("Valid Request", func(t *testing.T) {
		expected := time.Now().Format(time.RFC822)

		req := httptest.NewRequest(http.MethodGet, "/datetime/plain", nil)
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		got := rr.Body.String()

		if got != expected {
			t.Errorf("handler returned unexpected datetime: got %v want %v", got, expected)
		}
	})

	t.Run("Invalid Method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/datetime/plain", nil)
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
		}
	})
}

func TestGetTimeJsonHandler(t *testing.T) {
	handler := getTimeJsonHandler()

	t.Run("Valid Request", func(t *testing.T) {
		expected := time.Now().Format(time.RFC822)
		req := httptest.NewRequest(http.MethodGet, "/datetime/json", nil)
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		var timeResponse struct {
			Datetime string `json:"datetime"`
		}

		err := json.Unmarshal(rr.Body.Bytes(), &timeResponse)
		if err != nil {
			t.Fatalf("error unmarshalling response: %v", err)
		}

		got := timeResponse.Datetime

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if got != expected {
			t.Errorf("handler returned unexpected datetime: got %v want %v", got, expected)
		}
	})

	t.Run("Invalid Method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/datetime/json", nil)
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
		}

	})

}
