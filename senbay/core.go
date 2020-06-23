package senbay

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// BaseX descrive base number based on PN
type BaseX struct {
	PN           int
	Table        []int
	ReverseTable []int
}

// NewBaseX returns a new BaseX based on PN
func NewBaseX(positionalNotation int) (*BaseX, error) {
	baseX := &BaseX{
		PN: 121,
		Table: []int{
			1, 2, 3, 4,
			5, 6, 7, 8, 9,
			10, 11, 12, 13, 14,
			15, 16, 17, 18, 19,
			20, 21, 22, 23, 24,
			25, 26, 27, 28, 29,
			30, 31, 32, 33, 34,
			35, 36, 37, 38,
			40, 41, 42, 43,
			47, 48, 49,
			50, 51, 52, 53, 54,
			55, 56, 57, 59,
			60, 61, 62, 63, 64,
			65, 66, 67, 68, 69,
			70, 71, 72, 73, 74,
			75, 76, 77, 78, 79,
			80, 81, 82, 83, 84,
			85, 86, 87, 88, 89,
			90, 91, 92, 93, 94,
			95, 96, 97, 98, 99,
			100, 101, 102, 103, 104,
			105, 106, 107, 108, 109,
			110, 111, 112, 113, 114,
			115, 116, 117, 118, 119,
			120, 121, 122, 123, 124,
			125, 126, 127,
		},
		ReverseTable: []int{
			0, 0, 1, 2, 3,
			4, 5, 6, 7, 8,
			9, 10, 11, 12, 13,
			14, 15, 16, 17, 18,
			19, 20, 21, 22, 23,
			24, 25, 26, 27, 28,
			29, 30, 31, 32, 33,
			34, 35, 36, 37, 0,
			38, 39, 40, 41, 0,
			0, 0, 42, 43, 44,
			45, 46, 47, 48, 49,
			50, 51, 52, 0, 53,
			54, 55, 56, 57, 58,
			59, 60, 61, 62, 63,
			64, 65, 66, 67, 68,
			69, 70, 71, 72, 73,
			74, 75, 76, 77, 78,
			79, 80, 81, 82, 83,
			84, 85, 86, 87, 88,
			89, 90, 91, 92, 93,
			94, 95, 96, 97, 98,
			99, 100, 101, 102, 103,
			104, 105, 106, 107, 108,
			109, 110, 111, 112, 113,
			114, 115, 116, 117, 118,
			119, 120, 121,
		},
	}

	if positionalNotation > len(baseX.Table) || positionalNotation < 2 {
		msg := fmt.Sprint("shinsu must be 2-", len(baseX.Table))
		return baseX, errors.New(msg)
	}
	baseX.PN = positionalNotation

	return baseX, nil
}

func (baseX BaseX) encodeLongValue(lVal int) []rune {
	var isNegative bool
	if lVal < 0 {
		isNegative = true
		lVal *= -1
	}

	var places []int
	if lVal == 0 {
		places = append(places, 0)
	} else {
		for lVal > 0 {
			remainder := int(math.Mod(float64(lVal), float64(baseX.PN)))
			places = append(places, baseX.Table[remainder])
			lVal = lVal / baseX.PN
		}
	}

	var muString []rune
	for _, place := range places {
		muString = append([]rune{int32(place)}, muString...)
	}

	if isNegative {
		muString = append([]rune{'-'}, muString...)
	}

	return muString
}

func (baseX BaseX) encodeDoubleValue(dVal float64) []rune {
	var isNegative bool
	if dVal < 0 {
		isNegative = true
		dVal *= -1
	}

	strVal := strconv.FormatFloat(dVal, 'f', -1, 64)
	vals := strings.Split(strVal, ".")

	sVal, err := strconv.Atoi(vals[0])
	if err != nil {
		panic(err)
	}
	runeIntVal := baseX.encodeLongValue(sVal)
	if len(vals) == 1 {
		if isNegative {
			runeIntVal = append([]rune{'-'}, runeIntVal...)
		}
		return runeIntVal
	}

	sVal, err = strconv.Atoi(vals[1])
	if err != nil {
		panic(err)
	}
	strDecVal := baseX.encodeLongValue(sVal)

	var zeros []rune
	aZero := baseX.encodeLongValue(0)
	for _, aVal := range vals[1] {
		if string(aVal) == "0" {
			zeros = append(zeros, aZero...)
		}
	}

	var encoded []rune
	if isNegative {
		encoded = append(encoded, '-')
	}
	encoded = append(encoded, runeIntVal...)
	encoded = append(encoded, '.')
	encoded = append(encoded, zeros...)
	encoded = append(encoded, []rune(strDecVal)...)

	return encoded
}

func (baseX BaseX) decodeLongValue(sVal []rune) int {
	var isNegative bool
	if sVal[0] == '-' {
		isNegative = true
	}

	var totalVal float64
	for i := len(sVal) - 1; i >= 0; i-- {
		key := int(sVal[i])
		v := math.Pow(
			float64(baseX.PN),
			float64(len(sVal)-i-1),
		) * float64(baseX.ReverseTable[key])
		totalVal += v
	}

	if isNegative {
		return int(totalVal * -1)
	}

	return int(totalVal)
}

func (baseX BaseX) decodeDoubleValue(sVal []rune) float64 {
	var isNegative bool
	if sVal[0] == '-' {
		isNegative = true

		sVal = sVal[1:]
	}

	runeIntNum := []rune{}
	runeFloatNum := []rune{}

	var isFloat bool
	for _, num := range sVal {
		if num != '.' {
			if isFloat {
				runeFloatNum = append(runeFloatNum, num)
			} else {
				runeIntNum = append(runeIntNum, num)
			}
		} else {
			isFloat = true
		}
	}

	var intVal int
	if len(runeIntNum) != 0 {
		intVal = baseX.decodeLongValue(runeIntNum)
	}

	if len(runeFloatNum) == 0 {
		return float64(intVal)
	}

	var zeros []rune
	for _, aVal := range runeFloatNum {
		if aVal == 0 {
			zeros = append(zeros, aVal)
		} else {
			break
		}
	}
	decVal := baseX.decodeLongValue(runeFloatNum[len(zeros):])

	floatDigit := math.Floor(math.Log10(float64(decVal))+1.0) + float64(len(zeros))
	floatNum := float64(decVal) * math.Pow(0.1, floatDigit)
	if isNegative {
		return (float64(intVal) + floatNum) * -1.0
	}
	return float64(intVal) + floatNum
}

// Format is
type Format struct {
	ReversedKeys map[string]string
	PN           int
	baseX        *BaseX
}

// NewSenbayFormat returns a new Format based on PN
func NewSenbayFormat(PN int) (*Format, error) {
	baseX, err := NewBaseX(PN)
	if err != nil {
		return nil, err
	}
	senbayFormat := &Format{
		ReversedKeys: map[string]string{
			"TIME": "0", "LONG": "1", "LATI": "2",
			"ALTI": "3", "ACCX": "4", "ACCY": "5",
			"ACCZ": "6", "YAW": "7", "ROLL": "8",
			"PITC": "9", "HEAD": "A", "SPEE": "B",
			"BRIG": "C", "AIRP": "D", "HTBT": "E",
		},
		PN:    PN,
		baseX: baseX,
	}

	return senbayFormat, nil
}

func (senbayFormat Format) getReservedShortKey(key string) string {
	for k, v := range senbayFormat.ReversedKeys {
		if k == key {
			return v
		}
	}

	return ""
}

func (senbayFormat Format) getReservedOriginalKey(key string) string {
	for k, v := range senbayFormat.ReversedKeys {
		if v == key {
			return k
		}
	}

	return ""
}

func (senbayFormat Format) encode(text string) string {
	var encodedText string
	elements := strings.Split(text, ",")
	var count int

	for _, element := range elements {
		contents := strings.Split(element, ":")
		if len(contents) > 1 {
			key := contents[0]
			var val string
			for _, con := range contents[1:] {
				if val == "" {
					val = con
				} else {
					val = val + ":" + con
				}
			}
			reservedKey := senbayFormat.getReservedShortKey(key)
			var isReservedKey bool
			if len(reservedKey) != 0 {
				isReservedKey = true
				key = reservedKey
			}
			if len(val) > 0 {
				if val[:1] != "'" {
					floatVal, err := strconv.ParseFloat(val, 64)
					if err != nil {
						panic(err)
					}
					if isReservedKey {
						encodedText = encodedText + key + string(senbayFormat.baseX.encodeDoubleValue(floatVal))
					} else {
						encodedText = encodedText + key + ":" + string(senbayFormat.baseX.encodeDoubleValue(floatVal))
					}
				} else {
					if isReservedKey {
						encodedText = encodedText + key + val
					} else {
						encodedText = encodedText + key + ":" + val
					}
				}
			}
		}
		if count < len(elements)-1 {
			count++
			encodedText = encodedText + ","
		}
	}
	return encodedText
}

func (senbayFormat Format) decode(text string) string {
	var decodedText string
	var count int
	elements := strings.Split(text, ",")
	for _, element := range elements {
		var key []rune
		var val []rune
		contents := strings.Split(element, ":")
		if len(contents) > 1 {
			key = []rune(contents[0])
			for _, con := range contents[1:] {
				if len(val) == 0 {
					val = []rune(con)
				} else {
					val = append(val, ':')
					val = append(val, []rune(con)...)
				}
			}
		} else {
			key = []rune(contents[0])[:1]
			val = []rune(contents[0])[1:]
		}

		reservedKey := senbayFormat.getReservedOriginalKey(string(key))
		if reservedKey != "" {
			key = []rune(reservedKey)
		}

		if val[0] != '\'' {
			decodedDoubleValue := senbayFormat.baseX.decodeDoubleValue(val)
			decodedText = decodedText + string(key) + ":" + strconv.FormatFloat(decodedDoubleValue, 'f', -1, 64)
		} else {
			decodedText = decodedText + string(key) + ":" + string(val)
		}

		if count < len(elements)-1 {
			count++
			decodedText = decodedText + ";"
		}
	}
	return decodedText
}

// Frame is
type Frame struct {
	Data map[string]string
	PN   int
	SF   *Format
}

// NewSenbayFrame returns a new Frame based on PN
func NewSenbayFrame(PN int) (*Frame, error) {
	senbayFrame := &Frame{
		PN: PN,
	}
	SF, err := NewSenbayFormat(PN)
	if err != nil {
		return senbayFrame, err
	}
	senbayFrame.Data = map[string]string{}
	senbayFrame.SF = SF
	return senbayFrame, err
}

// AddNumber add number to data in senbayFrame
func (senbayFrame Frame) AddNumber(key string, value int) {
	senbayFrame.Data[key] = strconv.Itoa(value)
}

// AddText add text to data in senbayFrame
func (senbayFrame Frame) AddText(key string, value string) {
	senbayFrame.Data[key] = "'" + value + "'"
}

// Clear wil clear the data in senbayFrame
func (senbayFrame Frame) Clear() {
	// senbayFrame.Data = make(map[string]string)
	for key := range senbayFrame.Data {
		delete(senbayFrame.Data, key)
	}
}

// Encode will encode the data in senbayFrame
func (senbayFrame Frame) Encode(compress bool) string {
	var formattedData string
	var count int
	for k, v := range senbayFrame.Data {
		formattedData = formattedData + k + ":" + v
		if count < len(senbayFrame.Data)-1 {
			count++
			formattedData = formattedData + ","
		}
	}
	if compress {
		return "V:4," + senbayFrame.SF.encode(formattedData)
	}
	return "V:3," + formattedData
}

// Decode will decode the encoded text
func (senbayFrame Frame) Decode(text string) map[string]string {
	// "V:4,0YU97.+>H,16,2$."
	senbayMap := map[string]string{}
	elements := strings.Split(text, ",")
	var isCompress bool
	for _, element := range elements {
		contents := strings.Split(element, ":")
		if len(contents) > 1 && contents[0] == "V" && contents[1] == "4" {
			isCompress = true
			break
		}
	}
	if isCompress {
		text = senbayFrame.SF.decode(text)
	}
	elements = strings.Split(text, ";")
	for _, element := range elements {
		contents := strings.Split(element, ":")
		if len(contents) > 1 {
			key := contents[0]
			var value string
			for _, con := range contents[1:] {
				if value == "" {
					value = con
				} else {
					value = value + ":" + con
				}
			}

			if key != "V" {
				if value != "None" {
					if value[:1] == "'" {
						senbayMap[key] = value[1 : len(value)-1]
					} else {
						senbayMap[key] = value
					}
				}
			}
		}
	}
	return senbayMap
}
