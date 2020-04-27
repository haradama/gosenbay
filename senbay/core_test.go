package senbay

import (
	"testing"
)

func TestNewBaseX(t *testing.T) {
	_, err := NewBaseX(121)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	_, err = NewBaseX(200)
	if err == nil {
		t.Fatalf("failed test")
	}
}

func TestBaseXEncodeLongValue(t *testing.T) {
	baseX, err := NewBaseX(121)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	var expected []rune
	expected = []rune{0}
	result := baseX.encodeLongValue(0)
	if result[0] != expected[0] {
		t.Error("\nresult is", result)
	}
	result = baseX.encodeLongValue(100000)
	expected = []rune{7, 106, 60}
	for i, s := range result {
		if s != expected[i] {
			t.Error("\nresult is", result, "\nexpected： ", expected)
		}
	}
	result = baseX.encodeLongValue(200)
	expected = []rune{2, 85}
	for i, s := range result {
		if s != expected[i] {
			t.Error("\nresult is", result, "\nexpected： ", expected)
		}
	}
	result = baseX.encodeLongValue(-200)
	expected = []rune{45, 2, 85}
	for i, s := range result {
		if s != expected[i] {
			t.Error("\nresult is", result, "\nexpected： ", expected)
		}
	}
}

func TestBaseXEncodeDoubleValue(t *testing.T) {
	baseX, err := NewBaseX(121)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	var expected []rune

	result := baseX.encodeDoubleValue(0)
	expected = []rune{0}
	if result[0] != expected[0] {
		t.Error("\nresult is", result, "\nexpected： ", expected)
	}

	result = baseX.encodeDoubleValue(-3)
	expected = []rune{45, 4}
	for i, s := range result {
		if s != expected[i] {
			t.Error("\nresult is", result, "\nexpected： ", expected)
		}
	}

	result = baseX.encodeDoubleValue(3.14)
	expected = []rune{4, 46, 15}
	for i, s := range result {
		if s != expected[i] {
			t.Error("\nresult is", result, "\nexpected： ", expected)
		}
	}
	result = baseX.encodeDoubleValue(-3.04)
	expected = []rune{45, 4, 46, 0, 5}
	for i, s := range result {
		if s != expected[i] {
			t.Error("\nresult is", result, "\nexpected： ", expected)
		}
	}
}

func TestBaseXDecodeLongValue(t *testing.T) {
	baseX, err := NewBaseX(121)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	var indata []rune
	var expected int

	indata = []rune{0}
	totalVal := baseX.decodeLongValue(indata)
	expected = 0
	if totalVal != expected {
		t.Error("\nresult is", totalVal, "\nexpected： ", expected)
	}

	indata = []rune{2, 85}
	totalVal = baseX.decodeLongValue(indata)
	expected = 200
	if totalVal != expected {
		t.Error("\nresult is", totalVal, "\nexpected： ", expected)
	}

	indata = []rune{45, 2, 85}
	totalVal = baseX.decodeLongValue(indata)
	expected = -200
	if totalVal != expected {
		t.Error("\nresult is", totalVal, "\nexpected： ", expected)
	}

	indata = []rune{7, 106, 60}
	totalVal = baseX.decodeLongValue(indata)
	expected = 100000
	if totalVal != expected {
		t.Error("\nresult is", totalVal, "\nexpected： ", expected)
	}
}

func TestBaseXDecodeDoubleValue(t *testing.T) {
	baseX, err := NewBaseX(121)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	indata := []rune{4, 46, 15}
	totalVal := baseX.decodeDoubleValue(indata)
	expected := 3.14
	if totalVal != expected {
		t.Error("\nresult is", totalVal, "\nexpected： ", expected)
	}

	indata = []rune{4, 46, 0, 5}
	totalVal = baseX.decodeDoubleValue(indata)
	expected = 3.04
	if totalVal != expected {
		t.Error("\nresult is", totalVal, "\nexpected： ", expected)
	}
	indata = []rune{4}
	totalVal = baseX.decodeDoubleValue(indata)
	expected = 3
	if totalVal != expected {
		t.Error("\nresult is", totalVal, "\nexpected： ", expected)
	}

	indata = []rune{45, 4, 46, 0, 5}
	totalVal = baseX.decodeDoubleValue(indata)
	expected = -3.04
	if totalVal != expected {
		t.Error("\nresult is", totalVal, "\nexpected： ", expected)
	}

	indata = []rune{45, 46, 72, 123, 57}
	totalVal = baseX.decodeDoubleValue(indata)
	expected = -0.980515
	if totalVal != expected {
		t.Error("\nresult is", totalVal, "\nexpected： ", expected)
	}
}

func TestNewSenbayFormat(t *testing.T) {
	_, err := NewSenbayFormat(121)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	_, err = NewSenbayFormat(200)
	if err == nil {
		t.Fatalf("failed test")
	}
}

func TestGetReservedShortKey(t *testing.T) {
	senbayFormat, err := NewSenbayFormat(121)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	expected := "0"
	result := senbayFormat.getReservedShortKey("TIME")
	if result != expected {
		t.Error("\nresult is", result, "\nexpected： ", expected)
	}
}

func TestGetReservedOriginalKey(t *testing.T) {
	senbayFormat, err := NewSenbayFormat(121)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	expected := "TIME"
	result := senbayFormat.getReservedOriginalKey("0")
	if result != expected {
		t.Error("\nresult is", result, "\nexpected： ", expected)
	}
}

func TestDecode(t *testing.T) {
	senbayFormat, err := NewSenbayFormat(121)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	indata := "3+."
	expected := "ALTI:41"
	result := senbayFormat.decode(indata)
	if result != expected {
		t.Error("\nresult is", result, "\nexpected： ", expected)
	}

	indata = "4-.H{9,6-.|"
	expected = "ACCX:-0.980515,ACCZ:-0.118"
	result = senbayFormat.decode(indata)
	if result != expected {
		t.Error("\nresult is", result, "\nexpected： ", expected)
	}
}
