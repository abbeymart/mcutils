package mcutils

import (
	"fmt"
	"strings"
)

// Index functions returns the index of the specified value (for string, it's case-insensitive)
func Index[T ValueType](arr []T, val T) int {
	// types - string, int, float, bool
	for i, value := range arr {
		if value == val || strings.ToLower(fmt.Sprintf("%v", value)) == strings.ToLower(fmt.Sprintf("%v", val)) {
			return i
		}
	}
	return -1
}

// IndexCaseSensitive functions returns the index of the specified value (for string, it's case-sensitive)
func IndexCaseSensitive[T ValueType](arr []T, val T) int {
	// types - string, int, float, bool
	for i, value := range arr {
		if value == val {
			return i
		}
	}
	return -1
}

// ArrayContains check if a slice of generic type T contains/includes a value of type T  (for string, it's case-insensitive)
func ArrayContains[T ValueType](arr []T, str T) bool {
	for _, a := range arr {
		if a == str || strings.ToLower(fmt.Sprintf("%v", a)) == strings.ToLower(fmt.Sprintf("%v", str)) {
			return true
		}
	}
	return false
}

// ArrayContainsCaseSensitive check if a slice of generic type T contains/includes a value of type T  (for string, it's case-sensitive)
func ArrayContainsCaseSensitive[T ValueType](arr []T, str T) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// ArrayStringContains check if a slice of string contains/includes a string value, case-insensitive
func ArrayStringContains(arr []string, val string) bool {
	for _, a := range arr {
		if strings.ToLower(a) == strings.ToLower(val) {
			return true
		}
	}
	return false
}

// ArrayStringContainsCaseSensitive check if a slice of string contains/includes a string value, case-sensitive
func ArrayStringContainsCaseSensitive(arr []string, val string) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}

// ArrayIntContains check if a slice of int contains/includes an int value
func ArrayIntContains(arr []int, val int) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}

// ArrayFloatContains check if a slice of int contains/includes a float value
func ArrayFloatContains(arr []float64, str float64) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func Any[T ValueType](arr []T, testFunc TestFuncType[T]) bool {
	for _, value := range arr {
		if testFunc(value) {
			return true
		}
	}
	return false
}

func All[T ValueType](arr []T, testFunc TestFuncType[T]) bool {
	for _, value := range arr {
		if !testFunc(value) {
			return false
		}
	}
	return true
}

func Map[T ValueType](arr []T, mapFunc func(T) T) []T {
	var mapResult []T
	for _, v := range arr {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func MapGen[T ValueType](arr []T, mapFunc func(T) T, mapChan chan<- T) {
	for _, v := range arr {
		mapChan <- mapFunc(v)
	}
	if mapChan != nil {
		close(mapChan)
	}
}

func MapInt(arr []int, mapFunc func(int) int) []int {
	var mapResult []int
	for _, v := range arr {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func MapFloat(arr []float64, mapFunc func(float64) float64) []float64 {
	var mapResult []float64
	for _, v := range arr {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func MapString(arr []string, mapFunc func(string) string) []string {
	var mapResult []string
	for _, v := range arr {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func Filter[T ValueType](arr []T, filterFunc func(T) bool) []T {
	var mapResult []T
	for _, v := range arr {
		if filterFunc(v) {
			mapResult = append(mapResult, v)
		}
	}
	return mapResult
}

func FilterGen[T ValueType](arr []T, filterFunc func(T) bool, filterChan chan<- T) {
	for _, v := range arr {
		if filterFunc(v) {
			filterChan <- v
		}
	}
	if filterChan != nil {
		close(filterChan)
	}

}

func Take[T ValueType](num uint, arr []T) []T {
	var takeResult []T
	var cnt uint = 0
	for _, v := range arr {
		if cnt == num {
			break
		}
		takeResult = append(takeResult, v)
		cnt++
	}
	return takeResult
}

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
	if takeChan != nil {
		close(takeChan)
	}
}
