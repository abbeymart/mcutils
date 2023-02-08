package data

import "strings"

// test-data
var arrString = []string{"a", "b", "k", "m", "z", "q"}
var arrString2 = []string{"Abi", "boy", "king", "Mango", "zOo", "Quote"}

// var arrString3 = []string{"abi, boy, king, mango, zoo, quote"}
var arrInt = []int{10, 9, 4, 2, 5, 1}
var arrInt64 = []int64{100, 90, 45, 28, 49, 19}
var arrFloat = []float64{10.00, 9.50, 4.5, 4.2, 2.5, 1.5}

const arrIndexOfa = 0  // index of "a" = 0
const arrIndexOft = -1 // index of "t"
const arrCaseIndexOfB = -1
const arrCaseIndexOfb = 1

const arrStringContainsAbi = true
const arrStringCaseContainsabi = false

const arrIntContains10 = true
const arrIntContains44 = false
const arrFloatContains5Pt5 = true
const arrFloatContains15Pt5 = false

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
func mapDoubleIntValue(val int) int {
	return val * 2
}

var mapDoubleIntResult = []int{20, 18, 8, 4, 10, 2}

func mapDoubleFloatValue(val float64) float64 {
	return val * 2
}

var mapDoubleFloatResult = []float64{20.00, 19.00, 9.00, 8.4, 5.00, 3.00}

func mapStringAppendA(val string) string {
	return val + "A"
}

// []string{"Abi", "boy", "king", "Mango", "zOo", "Quote"}
var mapStringAppendAResult = []string{"AbiA", "boyA", "kingA", "MangoA", "zOoA", "QuoteA"}

// filter
func filterIntGreaterThan2(val int) bool {
	return val > 2
}

var filterIntGreaterThan2Result = []int{10, 9, 4, 5}

func filterFloatGreaterThan5(val float64) bool {
	return val > 5.00
}

var filterFloatGreaterThan5Result = []float64{10.00, 9.50}

func filterStringIncludeA(val string) bool {
	return strings.Contains(val, "b")
}

// []string{"Abi", "boy", "king", "Mango", "zOo", "Quote"}
var filterStringIncludeAResult = []string{"Abi", "boy"}

// take
var take2IntResult = []int{10, 9}
var take2StringResult = []string{"Abi", "boy"}

// reverse | []int{10, 9, 4, 2, 5, 1}
var reverseIntResult = []int{1, 5, 2, 4, 9, 10}

// []float64{10.00, 9.50, 4.5, 4.2, 2.5, 1.5}
var reverseFloatResult = []float64{1.5, 2.5, 4.2, 4.5, 9.50, 10.00}

// []string{"Abi", "boy", "king", "Mango", "zOo", "Quote"}
var reverseStringResult = []string{"Quote", "zOo", "Mango", "king", "boy", "Abi"}
