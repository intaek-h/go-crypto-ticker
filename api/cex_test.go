package api_test

import (
	"testing"

	"github.com/intaek-h/go-crypto-ticker/api"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("error was not found")
	}
}
