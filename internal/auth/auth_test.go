package auth

import (
	"errors"
	"net/http/httptest"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("Provide Empty Auth header request", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		r.Header.Add("Authorization", "")
		_, err := GetAPIKey(r.Header)

		if errors.Is(err, ErrNoAuthHeaderIncluded) {
			t.Error("Expected AuthHeaderNotIncluded Error but got none")
		}
	})
	t.Run("malformed Authorization header", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		r.Header.Add("Authorization", "APIKEY nigger")

		_, err := GetAPIKey(r.Header)

		if err == nil {
			t.Error("Expected Error but got none")
		}
	})
}
