package senbay

import (
	"testing"
)

func TestNewSenbayWriterStart(t *testing.T) {
	senbayWriter := NewSenbayWriter("test", 480, 640, 45, 0, "a", 60)
	if senbayWriter == nil {
		t.Fatalf("failed test")
	}

}
