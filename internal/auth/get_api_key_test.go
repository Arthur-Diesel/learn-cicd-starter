package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey abc123")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("GetAPIKey returned an unexpected error: %v", err)
	}
	if apiKey != "abc123" {
		t.Errorf("GetAPIKey returned %q, want %q", apiKey, "abc123")
	}
}

func TestGetAPIKeyWithoutAuthorizationHeader(t *testing.T) {
	apiKey, err := GetAPIKey(http.Header{})
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("GetAPIKey error = %v, want %v", err, ErrNoAuthHeaderIncluded)
	}
	if apiKey != "" {
		t.Errorf("GetAPIKey returned %q, want an empty API key", apiKey)
	}
}
