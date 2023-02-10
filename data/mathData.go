package data

import "github.com/abbeymart/mcutils"

const ALeapYear = 2000
const NotLeapYear = 2022

const FactorialParam = 4
const FactorialValue = 24

const FibSeriesParam = 8

var FibSeriesResult = []int{1, 1, 2, 3, 5, 8, 13, 21}

const PrimeNumParam = 10

var PrimeNums = []int{2, 3, 5, 7}

const PythagorasParam = 10

var PythagorasResult = [][]int{{3, 4, 5}, {6, 8, 10}}

const PythagorasParam100 = 100

var PythagorasResult100 = [][]int{
	{3, 4, 5}, {5, 12, 13}, {6, 8, 10},
	{7, 24, 25}, {8, 15, 17}, {9, 12, 15},
	{9, 40, 41}, {10, 24, 26}, {11, 60, 61},
	{12, 16, 20}, {12, 35, 37}, {13, 84, 85},
	{14, 48, 50}, {15, 20, 25}, {15, 36, 39},
	{16, 30, 34}, {16, 63, 65}, {18, 24, 30},
	{18, 80, 82}, {20, 21, 29}, {20, 48, 52},
	{20, 99, 101}, {21, 28, 35}, {21, 72, 75},
	{24, 32, 40}, {24, 45, 51}, {24, 70, 74},
	{25, 60, 65}, {27, 36, 45}, {28, 45, 53},
	{28, 96, 100}, {30, 40, 50}, {30, 72, 78},
	{32, 60, 68}, {33, 44, 55}, {33, 56, 65},
	{35, 84, 91}, {36, 48, 60}, {36, 77, 85},
	{39, 52, 65}, {39, 80, 89}, {40, 42, 58},
	{40, 75, 85}, {40, 96, 104}, {42, 56, 70},
	{45, 60, 75}, {48, 55, 73}, {48, 64, 80},
	{48, 90, 102}, {51, 68, 85}, {54, 72, 90},
	{56, 90, 106}, {57, 76, 95}, {60, 63, 87},
	{60, 80, 100}, {60, 91, 109}, {63, 84, 105},
	{65, 72, 97}, {66, 88, 110}, {69, 92, 115},
	{72, 96, 120}, {75, 100, 125}, {80, 84, 116},
}

var NumParams = []int{2, 5, 3, 5, 3, 5, 2, 3, 5}
var StringParams = []string{"a", "b", "a", "a", "a", "a"}
var BooleanParams = []bool{true, false, true, true, true, false, true, true}

//var symbolParams = [Symbol("abc"), Symbol("bcd"), Symbol("abc"), Symbol("bcd"), Symbol("abc")]
// {2: 2, 3: 3, 5: 4,}

var CountNumResult = mcutils.CounterResult[int]{
	"2": mcutils.CounterValue[int]{
		Count: 2,
		Value: 2,
	},
	"3": mcutils.CounterValue[int]{
		Count: 3,
		Value: 3,
	},
	"5": mcutils.CounterValue[int]{
		Count: 4,
		Value: 5,
	},
}

func CountNumResultKeys() []string {
	var result []string
	for k, _ := range CountNumResult {
		result = append(result, k)
	}
	return result
}

// {"a": 5, "b": 1,}

var CountStringResult = mcutils.CounterResult[string]{
	"a": mcutils.CounterValue[string]{
		Count: 5,
		Value: "a",
	},
	"b": mcutils.CounterValue[string]{
		Count: 1,
		Value: "b",
	},
}

func CountStringResultKeys() []string {
	var result []string
	for k, _ := range CountStringResult {
		result = append(result, k)
	}
	return result
}

var SetNumResult = []int{2, 5, 3}
var SetStingResult = []string{"a", "b"}
var SetBooleanResult = []bool{true, false}

//var setSymbolResult = [Symbol("abc"), Symbol("bcd"), Symbol("abc"), Symbol("bcd"), Symbol("abc")]
