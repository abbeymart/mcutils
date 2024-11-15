package mcutils

import (
	"fmt"
	"strings"
)

// Index functions returns the index of the specified Value (for string, it's case-insensitive)
func Index[T ValueType](arr []T, val T) int {
	// types - string, int, float, bool
	for i, value := range arr {
		if value == val || strings.ToLower(fmt.Sprintf("%v", value)) == strings.ToLower(fmt.Sprintf("%v", val)) {
			return i
		}
	}
	return -1
}

// IndexCaseSensitive functions returns the index of the specified Value (for string, it's case-sensitive)
func IndexCaseSensitive[T ValueType](arr []T, val T) int {
	// types - string, int, float, bool
	for i, value := range arr {
		if value == val {
			return i
		}
	}
	return -1
}

// ArrayContains check if a slice of generic type T contains/includes a Value of type T  (for string, it's case-insensitive)
func ArrayContains[T ValueType](arr []T, str T) bool {
	for _, a := range arr {
		if a == str || strings.ToLower(fmt.Sprintf("%v", a)) == strings.ToLower(fmt.Sprintf("%v", str)) {
			return true
		}
	}
	return false
}

// ArrayContainsCaseSensitive check if a slice of generic type T contains/includes a Value of type T  (for string, it's case-sensitive)
func ArrayContainsCaseSensitive[T ValueType](arr []T, str T) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// ArrayStringContains check if a slice of string contains/includes a string Value, case-insensitive
func ArrayStringContains(arr []string, val string) bool {
	for _, a := range arr {
		if strings.ToLower(a) == strings.ToLower(val) {
			return true
		}
	}
	return false
}

// ArrayStringContainsCaseSensitive check if a slice of string contains/includes a string Value, case-sensitive
func ArrayStringContainsCaseSensitive(arr []string, val string) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}

// ArrayIntContains check if a slice of int contains/includes an int Value
func ArrayIntContains(arr []int, val int) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}

// ArrayFloatContains check if a slice of int contains/includes a float Value
func ArrayFloatContains(arr []float64, str float64) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// Any function determines if one or more of the slice-values satisfies the testFunc.
func Any[T ValueType](arr []T, testFunc TestFuncType[T]) bool {
	for _, value := range arr {
		if testFunc(value) {
			return true
		}
	}
	return false
}

// All function determines if all the slice-values satisfies the testFunc.
func All[T ValueType](arr []T, testFunc TestFuncType[T]) bool {
	for _, value := range arr {
		if !testFunc(value) {
			return false
		}
	}
	return true
}

// include result address/pointer as function parameter to improve performance

// Map function map slice of type T, by the mapFunc.
func Map[T ValueType](arr []T, mapFunc func(T) T, result []T) {
	// reset the result Value
	result = []T{}
	for _, v := range arr {
		result = append(result, mapFunc(v))
	}
}

// MapGen function generates series of map Value of type T and send to the mapChan channel.
func MapGen[T ValueType](arr []T, mapFunc func(T) T, mapChan chan<- T) {
	for _, v := range arr {
		mapChan <- mapFunc(v)
	}
	// ends send task to the channel
	if mapChan != nil {
		close(mapChan)
	}
}

// MapInt function returns the mapped-slice-Value of type int, based on the mapFunc.
func MapInt(arr []int, mapFunc func(int) int, result []int) {
	result = []int{}
	for _, v := range arr {
		result = append(result, mapFunc(v))
	}
}

// MapFloat function returns the mapped-slice-Value, of type float64, based on the mapFunc.
func MapFloat(arr []float64, mapFunc func(float64) float64, result []float64) {
	result = []float64{}
	for _, v := range arr {
		result = append(result, mapFunc(v))
	}
}

// MapString function returns the mapped-slice-Value, of type string, based on the mapFunc.
func MapString(arr []string, mapFunc func(string) string, result []string) {
	result = []string{}
	for _, v := range arr {
		result = append(result, mapFunc(v))
	}
}

// Filter function returns the filtered-slice-Value, of type T, based on the filterFunc [generic].
func Filter[T ValueType](arr []T, filterFunc func(T) bool, result []T) {
	result = []T{}
	for _, v := range arr {
		if filterFunc(v) {
			result = append(result, v)
		}
	}
}

// FilterGen function returns series filtered-Value, of type T, based on the filterFunc [generic].
func FilterGen[T ValueType](arr []T, filterFunc func(T) bool, filterChan chan<- T) {
	for _, v := range arr {
		if filterFunc(v) {
			filterChan <- v
		}
	}
	// ends send task to the channel
	if filterChan != nil {
		close(filterChan)
	}
}

// Take function returns num of the arr slice of type T [generic].
func Take[T ValueType](num uint, arr []T, result []T) {
	result = []T{}
	var cnt uint = 0
	for _, v := range arr {
		if cnt == num {
			break
		}
		result = append(result, v)
		cnt++
	}
}

// TakeGen function returns num series of values, of type T [generic].
func TakeGen[T ValueType](num uint, arr []T, takeChan chan<- T) {
	// use channels to implement generator to send/yield/generate num of values from arr
	var cnt uint = 0
	for _, v := range arr {
		if cnt == num {
			break
		}
		takeChan <- v
		cnt++
	}
	// ends send task to the channel
	if takeChan != nil {
		close(takeChan)
	}
}

// ReverseArray returns the reverse values of the specified array/slice [generic type]
func ReverseArray[T ValueType](arr []T, result []T) {
	// arr and arrChan must be of the same type: int, float
	result = []T{}
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}
}

// ReverseArrayInt returns the reverse values of the specified array/slice of int.
func ReverseArrayInt(arr []int, result []int) {
	result = []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}
}

// ReverseArrayFloat returns the reverse values of the specified array/slice of float64.
func ReverseArrayFloat(arr []float64, result []float64) {
	result = []float64{}
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}
}

// ReverseArrayGenerator sequentially generates reverse values of the specified array/slice [generic]
func ReverseArrayGenerator[T ValueType](arr []T, arrChan chan<- T) {
	// arr and arrChan must be of the same type: int, float
	for i := len(arr) - 1; i >= 0; i-- {
		arrChan <- arr[i]
	}
	// ends send task to the channel
	if arrChan != nil {
		close(arrChan)
	}
}

// ReverseArrayIntGen sequentially generates reverse values of the specified array/slice of int.
func ReverseArrayIntGen(arr []int, arrChan chan<- int) {
	for i := len(arr) - 1; i >= 0; i-- {
		arrChan <- arr[i]
	}
	// ends send task to the channel
	if arrChan != nil {
		close(arrChan)
	}
}

// ReverseArrayFloatGen sequentially generates reverse values of the specified array/slice of float64.
func ReverseArrayFloatGen(arr []float64, arrChan chan<- float64) {
	for i := len(arr) - 1; i >= 0; i-- {
		arrChan <- arr[i]
	}
	// ends send task to the channel
	if arrChan != nil {
		close(arrChan)
	}
}

// TODO: review generic type-infer error
//func sumOfArrayGeneric[T Number](values []T, accumulator T) T {
//	if len(values) == 0 {
//		return accumulator
//	}
//	head := values[0]
//	var tail []T
//	if len(values) > 1 {
//		tail = values[1:]
//	} else {
//		tail = []T{}
//	}
//	return sumOfArrayGeneric(tail, head+accumulator)
//}

func sumOfArray(values []float64, accumulator float64) float64 {
	if len(values) == 0 {
		return accumulator
	}
	head := values[0]
	var tail []float64
	if len(values) > 1 {
		tail = values[1:]
	} else {
		tail = []float64{}
	}
	return sumOfArray(tail, head+accumulator)
}
