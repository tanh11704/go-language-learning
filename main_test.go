package main

import (
	"testing"

	"rsc.io/quote"
)

func TestHello(t *testing.T) {
	expected := "Don't communicate by sharing memory, share memory by communicating."
	if got := quote.Go(); got != expected {
		t.Errorf("quote.Go() = %q, want %q", got, expected)
	}
}