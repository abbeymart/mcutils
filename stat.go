// @Author: abbeymart | Abi Akindele | @Created: 2020-12-03 | @Updated: 2020-12-03
// @Company: mConnect.biz | @License: MIT
// @Description: statistical functions

package mcutils

import (
	"math"
	"sort"
)

// Mean function returns the mean or average value from a slice of type float.
func Mean(arr []float64) float64 {
	var sum = 0.00
	arrLength := len(arr)
	for _, v := range arr {
		sum += v
	}
	return sum / float64(arrLength)
}

// Median function returns the mid or median value from a slice of type float.
func Median(arr []float64) float64 {
	// sort numbers, ascending order
	sort.Float64s(arr)
	arrLength := len(arr)
	// if slice-items count is odd
	if arrLength%2 != 0 {
		medianIndex := math.Floor(float64(arrLength / 2))
		return arr[uint(medianIndex)]
	}
	// if slice-items count is even
	medianIndex2 := arrLength / 2
	medianIndex1 := medianIndex2 - 1
	return (arr[medianIndex1] + arr[medianIndex2]) / 2
}

// StandardDeviation function returns the standard-deviation value from a slice of type float.
func StandardDeviation(arr []float64) float64 {
	deltaSquareSum := 0.00
	arrLength := len(arr)
	mean := Mean(arr)
	for _, val := range arr {
		deltaSquareSum += math.Pow(val-mean, 2)
	}
	return math.Sqrt(deltaSquareSum / float64(arrLength-1))
}

// Min function returns the minimum value from a slice of type T (int or float).
func Min[T Number](arr []T) T {
	// set initial max value
	max := arr[0]
	// compute min and max values
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

// Max function returns the maximum value from a slice of type T (int or float).
func Max[T Number](arr []T) T {
	// set initial max value
	max := arr[0]
	// compute min and max values
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

// MinMax function returns the minimum amd maximum values from a slice of type T (int or float).
func MinMax[T Number](arr []T) (min T, max T) {
	// set initial min and max values
	min = arr[0]
	max = arr[0]
	// compute min and max values
	for _, val := range arr {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return
}

func MeanSquareError() {

}

func MeanSquareRootError() {

}

func Variance() {

}

func Knn() {

}

func NBayes() {

}

func Classify() {

}
