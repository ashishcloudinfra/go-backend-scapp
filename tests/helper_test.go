package tests

import (
	"testing"

	"server.simplifycontrol.com/helpers"
)

func TestHandleRequest(t *testing.T) {
	// Assuming you have a function `HandleRequest` in `handlers/handler.go`
	result := helpers.HandleRequest("test input")

	// Check if the result is expected
	expected := "expected output"
	if result != expected {
		t.Errorf("expected %s, but got %s", expected, result)
	}
}
