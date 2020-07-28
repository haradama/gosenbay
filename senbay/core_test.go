package senbay

import (
	"math"
	"testing"
)

// Test BaseX
func TestNewBaseX(t *testing.T) {
	_, err := NewBaseX(121)
	if err != nil {
		t.Error("failed test", err)
	}

	_, err = NewBaseX(200)
	if err == nil {
		t.Error("failed test")
	}
}
func TestBaseXEncodeLongValue(t *testing.T) {
	baseX, err := NewBaseX(121)
	if err != nil {
		t.Error("failed test", err)
	}

	var expect []rune
	expect = []rune{0}
	result := baseX.encodeLongValue(0)
	if result[0] != expect[0] {
		t.Error("\nresult:", result)
	}
	result = baseX.encodeLongValue(100000)
	expect = []rune{7, 106, 60}
	for i, s := range result {
		if s != expect[i] {
			t.Error("\nresult:", result, "\nexpect:", expect)
		}
	}
	result = baseX.encodeLongValue(200)
	expect = []rune{2, 85}
	for i, s := range result {
		if s != expect[i] {
			t.Error("\nresult:", result, "\nexpect:", expect)
		}
	}
	result = baseX.encodeLongValue(-200)
	expect = []rune{45, 2, 85}
	for i, s := range result {
		if s != expect[i] {
			t.Error("\nresult:", result, "\nexpect:", expect)
		}
	}
}

func TestBaseXEncodeDoubleValue(t *testing.T) {
	baseX, err := NewBaseX(121)
	if err != nil {
		t.Error("failed test", err)
	}
	var expect []rune

	result := baseX.encodeDoubleValue(0)
	expect = []rune{0}
	if result[0] != expect[0] {
		t.Error("\nresult:", result, "\nexpect:", expect)
	}

	result = baseX.encodeDoubleValue(-3)
	expect = []rune{45, 4}
	for i, s := range result {
		if s != expect[i] {
			t.Error("\nresult:", result, "\nexpect:", expect)
		}
	}

	result = baseX.encodeDoubleValue(3.14)
	expect = []rune{4, 46, 15}
	for i, s := range result {
		if s != expect[i] {
			t.Error("\nresult:", result, "\nexpect:", expect)
		}
	}
	result = baseX.encodeDoubleValue(-3.04)
	expect = []rune{45, 4, 46, 0, 5}
	for i, s := range result {
		if s != expect[i] {
			t.Error("\nresult:", result, "\nexpect:", expect)
		}
	}
}

func TestBaseXDecodeLongValue(t *testing.T) {
	baseX, err := NewBaseX(121)
	if err != nil {
		t.Error("failed test", err)
	}

	var indata []rune
	var expect int

	indata = []rune{0}
	totalVal := baseX.decodeLongValue(indata)
	expect = 0
	if totalVal != expect {
		t.Error("\nresult:", totalVal, "\nexpect:", expect)
	}

	indata = []rune{2, 85}
	totalVal = baseX.decodeLongValue(indata)
	expect = 200
	if totalVal != expect {
		t.Error("\nresult:", totalVal, "\nexpect:", expect)
	}

	indata = []rune{45, 2, 85}
	totalVal = baseX.decodeLongValue(indata)
	expect = -200
	if totalVal != expect {
		t.Error("\nresult:", totalVal, "\nexpect:", expect)
	}

	indata = []rune{7, 106, 60}
	totalVal = baseX.decodeLongValue(indata)
	expect = 100000
	if totalVal != expect {
		t.Error("\nresult:", totalVal, "\nexpect:", expect)
	}
}

func TestBaseXDecodeDoubleValue(t *testing.T) {
	baseX, err := NewBaseX(121)
	if err != nil {
		t.Error("failed test", err)
	}

	var indata []rune
	var totalVal float64
	var expect float64

	indata = []rune{4, 46, 15}
	totalVal = baseX.decodeDoubleValue(indata)
	expect = 3.14
	if totalVal != expect {
		t.Error("\nresult:", totalVal, "\nexpect:", expect)
	}

	indata = []rune{4, 46, 0, 5}
	totalVal = baseX.decodeDoubleValue(indata)
	expect = 3.04
	if totalVal != expect {
		t.Error("\nresult:", totalVal, "\nexpect:", expect)
	}
	indata = []rune{4}
	totalVal = baseX.decodeDoubleValue(indata)
	expect = 3
	if totalVal != expect {
		t.Error("\nresult:", totalVal, "\nexpect:", expect)
	}

	indata = []rune{45, 4, 46, 0, 5}
	totalVal = baseX.decodeDoubleValue(indata)
	expect = -3.04
	if totalVal != expect {
		t.Error("\nresult:", totalVal, "\nexpect:", expect)
	}

	indata = []rune{45, 46, 72, 123, 57}
	totalVal = baseX.decodeDoubleValue(indata)
	expect = -0.980515
	totalVal = math.Round(totalVal*1000000) / 1000000

	if totalVal != expect {
		t.Error("\nresult:", totalVal, "\nexpect:", expect)
	}
}

//　Test SenbayFormat
func TestNewSenbayFormat(t *testing.T) {
	_, err := NewSenbayFormat(121)
	if err != nil {
		t.Error("failed test", err)
	}

	_, err = NewSenbayFormat(200)
	if err == nil {
		t.Error("failed test")
	}
}

func TestSenbayFormatGetReservedShortKey(t *testing.T) {
	senbayFormat, err := NewSenbayFormat(121)
	if err != nil {
		t.Error("failed test", err)
	}

	expect := "0"
	result := senbayFormat.getReservedShortKey("TIME")
	if result != expect {
		t.Error("\nresult:", result, "\nexpect:", expect)
	}
}

func TestSenbayFormatGetReservedOriginalKey(t *testing.T) {
	senbayFormat, err := NewSenbayFormat(121)
	if err != nil {
		t.Error("failed test", err)
	}

	expect := "TIME"
	result := senbayFormat.getReservedOriginalKey("0")
	if result != expect {
		t.Error("\nresult:", result, "\nexpect:", expect)
	}
}

// Test SenbayData
func TestSenbayDataEncode(t *testing.T) {
	senbayData, err := NewSenbayData(121)
	if err != nil {
		t.Error("failed test", err)
	}

	senbayData.AddInt("KEY1", 234)
	senbayData.AddText("KEY2", "value2")

	result := senbayData.Encode(true)
	expect := "V:4,KEY1:w,KEY2:'value2'"
	if result != expect {
		t.Error("\nresult:", result, "\nexpect:", expect)
	}

	senbayData, err = NewSenbayData(121)
	if err != nil {
		t.Error("failed test", err)
	}
	senbayData.AddFloat("KEY1", 1000.34)
	senbayData.AddText("KEY2", "value2")

	result = senbayData.Encode(true)
	expect = "V:4,KEY1:	!.#,KEY2:'value2'"
	if result != expect {
		t.Error("\nresult:", result, "\nexpect:", expect)
	}
}

func TestSenbayDataClear(t *testing.T) {
	senbayData, err := NewSenbayData(121)
	if err != nil {
		t.Error("failed test", err)
	}

	senbayData.AddInt("KEY1", 123)
	senbayData.AddText("KEY2", "hello")
	senbayData.Clear()

	if len(senbayData.senbayData) != 0 {
		t.Error("\nresult:", senbayData)
	}
}

func TestSenbayFormatDecode(t *testing.T) {
	senbayFormat, err := NewSenbayFormat(121)
	if err != nil {
		t.Error("failed test", err)
	}

	indata := "3+."
	expect := "ALTI:41"
	result := senbayFormat.decode(indata)
	if result != expect {
		t.Error("\nresult:", result, "\nexpect:", expect)
	}

	indata = "4-.H{9,6-.|"
	expect = "ACCX:-0.9805150000000006;ACCZ:-0.11800000000000002"
	result = senbayFormat.decode(indata)
	if result != expect {
		t.Error("\nresult:", result, "\nexpect:", expect)
	}
}

func TestNewSenbayData(t *testing.T) {
	_, err := NewSenbayData(121)
	if err != nil {
		t.Error("failed test", err)
	}

	_, err = NewSenbayData(200)
	if err == nil {
		t.Error("failed test")
	}
}

func TestSenbayFormatEncode(t *testing.T) {
	SenbayData, err := NewSenbayData(121)
	if err != nil {
		t.Error("failed test", err)
	}

	indata := "V:4,0YU97.+>H,16,2$."
	expect := "ALTI:41"
	result := SenbayData.Decode(indata)
	if len(result) != 3 {
		t.Error("\nresult:", result, "\nexpect:", expect)
	}
}

func TestSenbayDataDecode(t *testing.T) {
	SenbayData, err := NewSenbayData(121)
	if err != nil {
		t.Error("failed test", err)
	}

	indata := "V:4,0YU97.+>H,16,2$."
	expect := "ALTI:41"
	result := SenbayData.Decode(indata)
	if len(result) != 3 {
		t.Error("\nresult:", result, "\nexpect:", expect)
	}
}
