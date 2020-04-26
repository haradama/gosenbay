package senbay

import (
	"testing"
)

func TestNewSenbayReader(t *testing.T) {
	senbayReader := NewSenbayReader(0, "", 0, 0)
	if senbayReader == nil {
		t.Fatalf("failed test")
	}
}
