package senbay

import (
	"testing"
)

func TestNewSenbayWriterStart(t *testing.T) {
	senbayWriter := NewSenbayWriter()
	if senbayWriter == nil {
		t.Fatalf("failed test")
	}

}
