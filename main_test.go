package main

import (
	"testing"

	"rsc.io/quote"
)

func TestGetMessage(t *testing.T) {
	expected := quote.Go()
	if got := GetMessage(); got != expected {
		t.Errorf("GetMessage() = %q, want %q", got, expected)
	}
}

func TestMainExecution(t *testing.T) {
	// Gọi hàm main()
	main()
}
