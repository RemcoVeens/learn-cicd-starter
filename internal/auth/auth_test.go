package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("basic good test", func(t *testing.T) {
		want := "pindas"
		header := http.Header{}
		header.Set("Authorization", fmt.Sprintf("ApiKey %s", want))
		res, err := GetAPIKey(header)
		if res != want {
			t.Fatalf("expected %s, got %s", want, res)
		}
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})
	t.Run("no auth header", func(t *testing.T) {
		header := http.Header{}
		res, err := GetAPIKey(header)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if res != "" {
			t.Fatalf("expected empty string, got %s", res)
		}
	})
	t.Run("malformed authorization header", func(t *testing.T) {
		header := http.Header{}
		header.Set("Authorization", "Bearer pindas")
		res, err := GetAPIKey(header)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if res != "" {
			t.Fatalf("expected empty string, got %s", res)
		}
	})
}
