package senbay

import (
	"math"
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
		t.Error("\nresult:", result)
	}
	result = baseX.encodeLongValue(100000)
	expected = []rune{7, 106, 60}
	for i, s := range result {
		if s != expected[i] {
			t.Error("\nresult:", result, "\nexpected:", expected)
		}
	}
	result = baseX.encodeLongValue(200)
	expected = []rune{2, 85}
	for i, s := range result {
		if s != expected[i] {
			t.Error("\nresult:", result, "\nexpected:", expected)
		}
	}
	result = baseX.encodeLongValue(-200)
	expected = []rune{45, 2, 85}
	for i, s := range result {
		if s != expected[i] {
			t.Error("\nresult:", result, "\nexpected:", expected)
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
		t.Error("\nresult:", result, "\nexpected:", expected)
	}

	result = baseX.encodeDoubleValue(-3)
	expected = []rune{45, 4}
	for i, s := range result {
		if s != expected[i] {
			t.Error("\nresult:", result, "\nexpected:", expected)
		}
	}

	result = baseX.encodeDoubleValue(3.14)
	expected = []rune{4, 46, 15}
	for i, s := range result {
		if s != expected[i] {
			t.Error("\nresult:", result, "\nexpected:", expected)
		}
	}
	result = baseX.encodeDoubleValue(-3.04)
	expected = []rune{45, 4, 46, 0, 5}
	for i, s := range result {
		if s != expected[i] {
			t.Error("\nresult:", result, "\nexpected:", expected)
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
		t.Error("\nresult:", totalVal, "\nexpected:", expected)
	}

	indata = []rune{2, 85}
	totalVal = baseX.decodeLongValue(indata)
	expected = 200
	if totalVal != expected {
		t.Error("\nresult:", totalVal, "\nexpected:", expected)
	}

	indata = []rune{45, 2, 85}
	totalVal = baseX.decodeLongValue(indata)
	expected = -200
	if totalVal != expected {
		t.Error("\nresult:", totalVal, "\nexpected:", expected)
	}

	indata = []rune{7, 106, 60}
	totalVal = baseX.decodeLongValue(indata)
	expected = 100000
	if totalVal != expected {
		t.Error("\nresult:", totalVal, "\nexpected:", expected)
	}
}

func TestBaseXDecodeDoubleValue(t *testing.T) {
	baseX, err := NewBaseX(121)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	var indata []rune
	var totalVal float64
	var expected float64

	indata = []rune{4, 46, 15}
	totalVal = baseX.decodeDoubleValue(indata)
	expected = 3.14
	if totalVal != expected {
		t.Error("\nresult:", totalVal, "\nexpected:", expected)
	}

	indata = []rune{4, 46, 0, 5}
	totalVal = baseX.decodeDoubleValue(indata)
	expected = 3.04
	if totalVal != expected {
		t.Error("\nresult:", totalVal, "\nexpected:", expected)
	}
	indata = []rune{4}
	totalVal = baseX.decodeDoubleValue(indata)
	expected = 3
	if totalVal != expected {
		t.Error("\nresult:", totalVal, "\nexpected:", expected)
	}

	indata = []rune{45, 4, 46, 0, 5}
	totalVal = baseX.decodeDoubleValue(indata)
	expected = -3.04
	if totalVal != expected {
		t.Error("\nresult:", totalVal, "\nexpected:", expected)
	}

	indata = []rune{45, 46, 72, 123, 57}
	totalVal = baseX.decodeDoubleValue(indata)
	expected = -0.980515
	totalVal = math.Round(totalVal*1000000) / 1000000

	if totalVal != expected {
		t.Error("\nresult:", totalVal, "\nexpected:", expected)
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

func TestSenbayFormatGetReservedShortKey(t *testing.T) {
	senbayFormat, err := NewSenbayFormat(121)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	expected := "0"
	result := senbayFormat.getReservedShortKey("TIME")
	if result != expected {
		t.Error("\nresult:", result, "\nexpected:", expected)
	}
}

func TestSenbayFormatGetReservedOriginalKey(t *testing.T) {
	senbayFormat, err := NewSenbayFormat(121)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	expected := "TIME"
	result := senbayFormat.getReservedOriginalKey("0")
	if result != expected {
		t.Error("\nresult:", result, "\nexpected:", expected)
	}
}

func TestSenbayFormatEncode(t *testing.T) {
	senbayFrame, err := NewSenbayFrame(121)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	senbayFrame.AddNumber("KEY1", 123)
	senbayFrame.AddText("KEY2", "hello")

	result := senbayFrame.Encode(false)
	if len(result) != 3 {
		t.Error("\nresult:", result)
	}

	result = senbayFrame.Encode(true)
	if len(result) != 3 {
		t.Error("\nresult:", result)
	}
}

func TestSenbayFormatClear(t *testing.T) {
	senbayFrame, err := NewSenbayFrame(121)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	senbayFrame.AddNumber("KEY1", 123)
	senbayFrame.AddText("KEY2", "hello")
	senbayFrame.Clear()

	if len(senbayFrame.Data) != 0 {
		t.Error("\nresult:", senbayFrame)
	}
}

func TestSenbayFormatDecode(t *testing.T) {
	senbayFormat, err := NewSenbayFormat(121)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	indata := "3+."
	expected := "ALTI:41"
	result := senbayFormat.decode(indata)
	if result != expected {
		t.Error("\nresult:", result, "\nexpected:", expected)
	}

	indata = "4-.H{9,6-.|"
	expected = "ACCX:-0.9805150000000006;ACCZ:-0.11800000000000002"
	result = senbayFormat.decode(indata)
	if result != expected {
		t.Error("\nresult:", result, "\nexpected:", expected)
	}
}

func TestNewSenbayFrame(t *testing.T) {
	_, err := NewSenbayFrame(121)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	_, err = NewSenbayFrame(200)
	if err == nil {
		t.Fatalf("failed test")
	}
}

func TestSenbayFrameDecode(t *testing.T) {
	senbayFrame, err := NewSenbayFrame(121)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	indata := "V:4,0YU97.+>H,16,2$."
	expected := "ALTI:41"
	result := senbayFrame.Decode(indata)
	if len(result) != 3 {
		t.Error("\nresult:", result, "\nexpected:", expected)
	}
}
