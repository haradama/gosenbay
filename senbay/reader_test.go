package senbay

import (
	"testing"
)

func TestNewSenbayReader(t *testing.T) {
	senbayReader := NewSenbayReader(0, "", 0, 0, false)
	if senbayReader == nil {
		t.Fatalf("failed test")
	}
}
