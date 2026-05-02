// Package senbay provides the functions to encode and
// decode to the senbay format.
package senbay

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const keyCodeEsc = 27

// BaseX descrive base number based on PN
type BaseX struct {
	PN           int
	Table        []int
	ReverseTable []int
}

var (
	reversedKeys = map[string]string{
		"TIME": "0", "LONG": "1", "LATI": "2",
		"ALTI": "3", "ACCX": "4", "ACCY": "5",
		"ACCZ": "6", "YAW": "7", "ROLL": "8",
		"PITC": "9", "HEAD": "A", "SPEE": "B",
		"BRIG": "C", "AIRP": "D", "HTBT": "E",
	}
)

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

// encodeLongValue returns
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
	if len(sVal) == 0 {
		return 0
	}

	var isNegative bool
	if sVal[0] == '-' {
		isNegative = true
		sVal = sVal[1:]
	}

	var totalVal float64
	for i := len(sVal) - 1; i >= 0; i-- {
		key := int(sVal[i])

		if key < 0 || key >= len(baseX.ReverseTable) {
			continue
		}

		v := math.Pow(
			float64(baseX.PN),
			float64(len(sVal)-i-1),
		) * float64(baseX.ReverseTable[key])
		totalVal += v
	}

	if isNegative {
		return int(totalVal) * -1
	}

	return int(totalVal)
}

func (baseX BaseX) decodeDoubleValue(sVal []rune) float64 {
	if len(sVal) == 0 {
		return 0
	}

	var isNegative bool
	if sVal[0] == '-' {
		isNegative = true
		sVal = sVal[1:]
	}

	vals := strings.Split(string(sVal), ".")

	switch len(vals) {
	case 1:
		intVal := baseX.decodeLongValue([]rune(vals[0]))
		if isNegative {
			return float64(intVal) * -1
		}
		return float64(intVal)

	case 2:
		intVal := baseX.decodeLongValue([]rune(vals[0]))

		frac := []rune(vals[1])
		aZero := baseX.encodeLongValue(0)

		zeroCount := 0
		if len(aZero) > 0 {
			for zeroCount < len(frac) && frac[zeroCount] == aZero[0] {
				zeroCount++
			}
		}

		decVal := baseX.decodeLongValue(frac[zeroCount:])

		numText := strconv.Itoa(intVal) + "." + strings.Repeat("0", zeroCount) + strconv.Itoa(decVal)
		floatVal, err := strconv.ParseFloat(numText, 64)
		if err != nil {
			return 0
		}

		if isNegative && floatVal >= 0 {
			return floatVal * -1
		}
		return floatVal

	default:
		intVal := baseX.decodeLongValue(sVal)
		if isNegative {
			return float64(intVal) * -1
		}
		return float64(intVal)
	}
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
		ReversedKeys: reversedKeys,
		PN:           PN,
		baseX:        baseX,
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

// decode
// decode decodes compressed Senbay text.
func (senbayFormat Format) decode(text string) string {
	var decodedElements []string

	elements := strings.Split(text, ",")
	for _, element := range elements {
		if element == "" {
			continue
		}

		var key string
		var val string

		contents := strings.Split(element, ":")
		if len(contents) > 1 {
			key = contents[0]
			val = strings.Join(contents[1:], ":")
		} else {
			runes := []rune(contents[0])
			if len(runes) == 0 {
				continue
			}

			key = string(runes[:1])
			val = string(runes[1:])
		}

		if key == "" || val == "" {
			continue
		}

		if key == "V" {
			decodedElements = append(decodedElements, key+":"+val)
			continue
		}

		reservedKey := senbayFormat.getReservedOriginalKey(key)
		if reservedKey != "" {
			key = reservedKey
		}

		if !strings.HasPrefix(val, "'") {
			decodedDoubleValue := senbayFormat.baseX.decodeDoubleValue([]rune(val))
			decodedElements = append(
				decodedElements,
				key+":"+strconv.FormatFloat(decodedDoubleValue, 'f', -1, 64),
			)
		} else {
			decodedElements = append(decodedElements, key+":"+val)
		}
	}

	return strings.Join(decodedElements, ",")
}

// A Data is
type Data struct {
	senbayData map[string]string
	PN         int
	SF         *Format
}

// NewSenbayData returns a new SenbayData struct based on PN
func NewSenbayData(PN int) (*Data, error) {
	SD := &Data{
		PN: PN,
	}
	SF, err := NewSenbayFormat(PN)
	if err != nil {
		return SD, err
	}
	SD.senbayData = map[string]string{}
	SD.SF = SF
	return SD, err
}

// AddInt add int value to senbayData
func (SD Data) AddInt(key string, value int) {
	SD.senbayData[key] = strconv.Itoa(value)
}

// AddInt64 add int64 value to senbayData
func (SD Data) AddInt64(key string, value int64) {
	SD.senbayData[key] = strconv.FormatInt(value, 10)
}

// AddFloat add float value to senbayData
func (SD Data) AddFloat(key string, value float32) {
	SD.senbayData[key] = strconv.FormatFloat(float64(value), 'f', -1, 64)
}

// AddFloat64 add float64 value to senbayData
func (SD Data) AddFloat64(key string, value float64) {
	SD.senbayData[key] = strconv.FormatFloat(value, 'f', -1, 64)
}

// AddText add string value to senbayData
func (SD Data) AddText(key string, value string) {
	SD.senbayData[key] = "'" + value + "'"
}

// Clear empties the contents of Data
func (SD Data) Clear() {
	for key := range SD.senbayData {
		delete(SD.senbayData, key)
	}
}

// Encode converts the data to decoded.
func (SD Data) Encode(compress bool) string {
	var formattedData string
	var count int
	for k, v := range SD.senbayData {
		formattedData = formattedData + k + ":" + v
		if count < len(SD.senbayData)-1 {
			count++
			formattedData = formattedData + ","
		}
	}
	if compress {
		return "V:4," + SD.SF.encode(formattedData)
	}
	return "V:3," + formattedData
}

// Decode converts the decoded text to the original data.
func (SD Data) Decode(text string) map[string]string {
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
		text = SD.SF.decode(text)
	}

	elements = strings.Split(text, ",")
	for _, element := range elements {
		if element == "" {
			continue
		}

		contents := strings.Split(element, ":")
		if len(contents) <= 1 {
			continue
		}

		key := contents[0]
		value := strings.Join(contents[1:], ":")

		if key == "" || key == "V" {
			continue
		}
		if value == "" || value == "None" {
			continue
		}

		if strings.HasPrefix(value, "'") {
			if len(value) >= 2 {
				senbayMap[key] = value[1 : len(value)-1]
			} else {
				senbayMap[key] = ""
			}
		} else {
			senbayMap[key] = value
		}
	}

	return senbayMap
}
