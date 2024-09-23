package timehttp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetTimeHandler(t *testing.T) {
	handler := GetTimeHandler()

	t.Run("Valid Request with JSON Accept Header", func(t *testing.T) {
		expected := time.Now().Format(time.RFC822)
		req := httptest.NewRequest(http.MethodGet, "/datetime", nil)
		req.Header.Set("Accept", "application/json")
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

	t.Run("Valid Request with Plain Text Accept Header", func(t *testing.T) {
		expected := time.Now().Format(time.RFC822)
		req := httptest.NewRequest(http.MethodGet, "/datetime", nil)
		req.Header.Set("Accept", "text/plain")
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		got := rr.Body.String()

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if got != expected {
			t.Errorf("handler returned unexpected datetime: got %v want %v", got, expected)
		}
	})

	t.Run("Valid Request with No Accept Header", func(t *testing.T) {
		expected := time.Now().Format(time.RFC822)
		req := httptest.NewRequest(http.MethodGet, "/datetime", nil)
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		got := rr.Body.String()

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if got != expected {
			t.Errorf("handler returned unexpected datetime: got %v want %v", got, expected)
		}
	})

	t.Run("Invalid Method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/datetime", nil)
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
		}
	})
}
