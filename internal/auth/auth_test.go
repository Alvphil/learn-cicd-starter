package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Set("Authorization", "ApiKey")
	_, err := GetAPIKey(req.Header)
	if err != nil {
		t.Fatalf("%v", err)
	}
}
