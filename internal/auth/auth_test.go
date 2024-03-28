package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Set("Authorization", "ApiKey 123")
	_, err := GetAPIKey(req.Header)
	if err != nil {
		t.Fatalf("%v", err)
	}
}
