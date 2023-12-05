package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey 123456")
	_, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

