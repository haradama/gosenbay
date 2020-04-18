package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type BaseX struct {
	PN int
	TABLE []int
	REVERSE_TABLE []int
}

func NewBaseX(positionalNotation int) *BaseX {
	baseX := &BaseX{
		PN: 1,
		TABLE: []int{
			1, 2, 3, 4,
			5,   6,  7,  8,  9,
			10,  11,  12,  13,  14,
			15,  16,  17,  18,  19,
			20,  21,  22,  23,  24,
			25,  26,  27,  28,  29,
			30,  31,  32,  33,  34,
			35,  36,  37,  38,
			40,  41,  42,  43,
					47,  48,  49,
			50,  51,  52,  53,  54,
			55,  56,  57,       59,
			60,  61,  62,  63,  64,
			65,  66,  67,  68,  69,
			70,  71,  72,  73,  74,
			75,  76,  77,  78,  79,
			80,  81,  82,  83,  84,
			85,  86,  87,  88,  89,
			90,  91,  92,  93,  94,
			95,  96,  97,  98,  99,
			100,  101,  102,  103,  104,
			105,  106,  107,  108,  109,
			110,  111,  112,  113,  114,
			115,  116,  117,  118,  119,
			120,  121,  122,  123,  124,
			125,  126,  127,
		},
		REVERSE_TABLE: []int{
			0,  0,  1,  2,  3,
			4,  5,  6,  7,  8,
			9,  10,  11,  12,  13,
			14,  15,  16,  17,  18,
			19,  20,  21,  22,  23,
			24,  25,  26,  27,  28,
			29,  30, 31,  32,  33,
			34,  35,  36,  37, 0,
			38,  39,  40,  41, 0,
			0,  0,  42, 43, 44,
			45,  46,  47,  48,  49,
			50,  51,  52,   0,   53,
			54,  55,  56,  57,  58,
			59,  60,  61,  62,  63,
			64,  65,  66,  67,  68,
			69,  70,  71,  72,  73,
			74,  75,  76,  77,  78,
			79,  80,  81,  82,  83,
			84,  85,  86,  87,  88,
			89,  90,  91,  92,  93,
			94,  95,  96,  97,  98,
			99,  100,  101,  102,  103,
			104,  105,  106,  107,  108,
			109,  110,  111,  112,  113,
			114,  115,  116,  117,  118,
			119,  120,  121,
		},
	}

	if positionalNotation > len(baseX.TABLE) || positionalNotation < 2 {
		fmt.Println("shinsu must be 2-", len(baseX.TABLE))
	} else {
		baseX.PN = positionalNotation
	}

	return baseX
}

func (baseX BaseX) encodeLongValue(lVal int) string {
	var isNegative bool
	if lVal < 0 {
		isNegative = true
		lVal *= -1
	}

	var places []int
	if lVal == 0 {
		places = append(places, 0)
	} else {
		for (lVal > 0) {
			remainder := int(math.Mod(float64(lVal), float64(baseX.PN)))
			places = append(places, baseX.TABLE[remainder])
			lVal = lVal / baseX.PN
		}
	}

	var muString string
	for _, place := range places {
		muString = string(place) + muString
	}

	if isNegative {
		return "_" + muString
	}
	
	return muString
}

func (baseX BaseX) encodeDoubleValue(dVal float64) string {
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
	strIntVal := baseX.encodeLongValue(sVal)
	if len(vals) == 1 {
		if isNegative {
			return "_" + strIntVal
		}
		return strIntVal
	}

	sVal, err = strconv.Atoi(vals[1])
	if err != nil {
		panic(err)
	}
	strDecVal := baseX.encodeLongValue(sVal)

	var zeros string
	aZero := baseX.encodeLongValue(0)
	for _, aVal := range vals[1] {
		if aVal == 0 {
			zeros = zeros + aZero
		}
	}

	if isNegative {
		return "_" + strIntVal + "." + zeros + strDecVal
	}
	
	return strIntVal + "." + zeros + strDecVal
}

func (baseX BaseX) decodeLongValue(sVal string) int {
	var isNegative bool
	if sVal[0:1] == "_" {
		isNegative = true
	}

	var totalVal float64
	for i := len(sVal) - 1; i > 0; i-- {
		v := math.Pow(float64(baseX.PN), float64(len(sVal) - (i + 1))) * float64(baseX.REVERSE_TABLE[int(sVal[i])])
		totalVal += v
	}

	if isNegative {
		return int(totalVal * -1)
	}
	
	return int(totalVal)
}

func (baseX BaseX) decodeDoubleValue(sVal string) float64 {
	var isNegative bool
	if sVal[0:1] == "-" {
		isNegative = true
	}

	vals := strings.Split(sVal, ".")
	var intVal int
	if len(vals) == 1 {
		intVal = baseX.decodeLongValue(vals[0])
	} else if len(vals) == 2 {
		intVal = baseX.decodeLongValue(vals[0])
		aZero := baseX.encodeLongValue(0)
		var zeroNum int
		for _, aVal := range vals[1] {
			if string(aVal) == aZero {
				zeroNum++
			} else {
				break
			}
		}
		decVal := baseX.decodeLongValue(vals[1][zeroNum:len(vals[1])])

		result := float64(intVal) + math.Pow(0.1, float64(zeroNum + 1)) * float64(decVal)
		if isNegative && intVal >= 0 {
			return -1 * result
		}
		return result
	}

	return float64(baseX.decodeLongValue(sVal))
}

type SenbayFormat struct {
	RESERVED_KEYS map[string]string
	PN int
	basex BaseX
}

func main() {
	fmt.Println("hello")
}