package timegin

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestGetTimeHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.HandleMethodNotAllowed = true

	r.GET("/datetime", GetTimeHandler)

	t.Run("Valid Request JSON", func(t *testing.T) {
		expected := time.Now().Format(time.RFC822)
		req := httptest.NewRequest(http.MethodGet, "/datetime", nil)
		req.Header.Set("Accept", "application/json")
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		var timeResponse struct {
			DateTime string `json:"DateTime"`
		}

		err := json.Unmarshal(rr.Body.Bytes(), &timeResponse)
		if err != nil {
			t.Fatalf("error unmarshalling response: %v", err)
		}

		got := timeResponse.DateTime

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if got != expected {
			t.Errorf("handler returned unexpected datetime: got %v want %v", got, expected)
		}
	})

	t.Run("Valid Request Plain Text", func(t *testing.T) {
		expected := time.Now().Format(time.RFC822)
		req := httptest.NewRequest(http.MethodGet, "/datetime", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

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

		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
		}
	})
}
