package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	headers := http.Header{
		"Content-Type": {"application/json"},
	}
	headers.Set("Authorization", "ApiKey your-token-here")
	_, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Unexpected nil %v", err)
	}

	headers.Set("Authorization", "BearerToken your-token-here")
	_, err2 := GetAPIKey(headers)
	if err2 == nil {
		t.Errorf("Expected error when invalid key given: %v", err2)
	}
}
