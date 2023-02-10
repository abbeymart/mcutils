package data

import "strings"

// test-data

var ArrString = []string{"a", "b", "k", "m", "z", "q"}
var ArrString2 = []string{"Abi", "boy", "king", "Mango", "zOo", "Quote"}

// var arrString3 = []string{"abi, boy, king, mango, zoo, quote"}

var ArrInt = []int{10, 9, 4, 2, 5, 1}
var ArrInt64 = []int64{100, 90, 45, 28, 49, 19}
var ArrFloat = []float64{10.00, 9.50, 4.5, 4.2, 2.5, 1.5}

const ArrIndexOfa = 0  // index of "a" = 0
const ArrIndexOft = -1 // index of "t"
const ArrCaseIndexOfB = -1
const ArrCaseIndexOfb = 1

const ArrStringContainsAbi = true
const ArrStringCaseContainsabi = false

const ArrIntContains10 = true
const ArrIntContains44 = false
const ArrFloatContains5Pt5 = true
const ArrFloatContains15Pt5 = false

func AnyIntFunc(val int) bool {
	return val > 5
}
func NotAnyIntFunc(val int) bool {
	return val > 500
}
func AllIntFunc(val int) bool {
	return val > 1
}
func NotAllIntFunc(val int) bool {
	return val > 1000
}

func AnyFloatFunc(val float64) bool {
	return val > 5.00
}
func NotAnyFloatFunc(val float64) bool {
	return val > 500.00
}
func AllFloatFunc(val float64) bool {
	return val > 1.25
}
func NotAllFloatFunc(val float64) bool {
	return val > 1000.00
}

// map

func MapDoubleIntValue(val int) int {
	return val * 2
}

var MapDoubleIntResult = []int{20, 18, 8, 4, 10, 2}

func MapDoubleFloatValue(val float64) float64 {
	return val * 2
}

var MapDoubleFloatResult = []float64{20.00, 19.00, 9.00, 8.4, 5.00, 3.00}

func MapStringAppendA(val string) string {
	return val + "A"
}

// []string{"Abi", "boy", "king", "Mango", "zOo", "Quote"}

var MapStringAppendAResult = []string{"AbiA", "boyA", "kingA", "MangoA", "zOoA", "QuoteA"}

// filter

func FilterIntGreaterThan2(val int) bool {
	return val > 2
}

var FilterIntGreaterThan2Result = []int{10, 9, 4, 5}

func FilterFloatGreaterThan5(val float64) bool {
	return val > 5.00
}

var FilterFloatGreaterThan5Result = []float64{10.00, 9.50}

func FilterStringIncludeA(val string) bool {
	return strings.Contains(val, "b")
}

// []string{"Abi", "boy", "king", "Mango", "zOo", "Quote"}

var FilterStringIncludeAResult = []string{"Abi", "boy"}

// take

var Take2IntResult = []int{10, 9}
var Take2StringResult = []string{"Abi", "boy"}

// reverse | []int{10, 9, 4, 2, 5, 1}

var ReverseIntResult = []int{1, 5, 2, 4, 9, 10}

// []float64{10.00, 9.50, 4.5, 4.2, 2.5, 1.5}

var ReverseFloatResult = []float64{1.5, 2.5, 4.2, 4.5, 9.50, 10.00}

// []string{"Abi", "boy", "king", "Mango", "zOo", "Quote"}

var ReverseStringResult = []string{"Quote", "zOo", "Mango", "king", "boy", "Abi"}
