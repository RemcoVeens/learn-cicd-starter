package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Fatalf("fix later")
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
}
