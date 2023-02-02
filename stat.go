// @Author: abbeymart | Abi Akindele | @Created: 2020-12-03 | @Updated: 2020-12-03
// @Company: mConnect.biz | @License: MIT
// @Description: statistical functions

package mcutils

import (
	"math"
	"sort"
)

// Mean function returns the mean or average value from a slice of type float.
func Mean[T Number](arr []T) float64 {
	var sum = 0.00
	arrLength := len(arr)
	for _, v := range arr {
		sum += float64(v)
	}
	return sum / float64(arrLength)
}

// GeometricMean function returns the geometric mean of a slice of type float.
func GeometricMean[T Number](arr []T) float64 {
	var multi = 0.00
	arrLength := len(arr)
	for _, v := range arr {
		multi *= float64(v)
	}
	return math.Pow(multi, float64(1/arrLength))
}

// Median function returns the mid or median value from a slice of type float.
func Median[T Number](arr []T) float64 {
	// sort numbers, ascending order
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })
	arrLength := len(arr)
	// if slice-items count is odd
	if arrLength%2 != 0 {
		medianIndex := math.Floor(float64(arrLength / 2))
		return float64(arr[uint(medianIndex)])
	}
	// if slice-items count is even
	medianIndex2 := arrLength / 2
	medianIndex1 := medianIndex2 - 1
	return float64((arr[medianIndex1] + arr[medianIndex2]) / 2)
}

// Mode function returns the mode(most frequently occurring value(s)) of a slice of type float.
func Mode[T Number](arr []T) []CounterValue[T] {
	// Obtain the counter values for the arr items
	result := ArrayValue[T](arr)
	arrCounters := result.Counter()
	var modes []CounterValue[T]
	// get the maximum or the highest occurrence of the arrCounter values
	var counters []int
	for _, cVal := range arrCounters {
		counters = append(counters, cVal.Count)
	}
	max := Max(counters)
	// compute mode/modes, i.e. modal value greater than 1
	if max > 1 {
		for _, cVal := range arrCounters {
			if cVal.Count == max {
				modes = append(modes, cVal)
			}
		}
	}
	return modes
}

// Frequency function returns the frequency / occurrence of a slice of type float.
func Frequency[T Number](arr []T) []CounterValue[T] {
	// Obtain the counter values for the arr items
	result := ArrayValue[T](arr)
	arrCounters := result.Counter()
	var modes []CounterValue[T]
	// compute the frequency/occurrence
	for _, cVal := range arrCounters {
		modes = append(modes, cVal)
	}
	return modes
}

// Range function returns the range of the slice type of float.
func Range[T Number](arr []T) T {
	min, max := MinMax(arr)
	return max - min
}

// Variance function returns the variance of the mean-square-value from a slice of type float.
func Variance[T Number](arr []T) float64 {
	meanSquareSum := 0.00
	arrLength := len(arr)
	mean := Mean(arr)
	for _, val := range arr {
		meanSquareSum += math.Pow(float64(val)-mean, 2)
	}
	return meanSquareSum / float64(arrLength)
}

// SampleStandardDeviation function returns the standard-deviation value from a sample-data of slice of type float.
func SampleStandardDeviation[T Number](arr []T) float64 {
	deltaSquareSum := 0.00
	arrLength := len(arr)
	mean := Mean(arr)
	for _, val := range arr {
		deltaSquareSum += math.Pow(float64(val)-mean, 2)
	}
	return math.Sqrt(deltaSquareSum / float64(arrLength-1))
}

// PopulationStandardDeviation function returns the standard-deviation value from a population/complete-data of slice of type float.
func PopulationStandardDeviation[T Number](arr []T) float64 {
	deltaSquareSum := 0.00
	arrLength := len(arr)
	mean := Mean(arr)
	for _, val := range arr {
		deltaSquareSum += math.Pow(float64(val)-mean, 2)
	}
	return math.Sqrt(deltaSquareSum / float64(arrLength))
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

// Quartiles returns slice-values that separate the data into four equal parts.
// Q0, Q1, Q2, Q3 & Q4
func Quartiles[T Number](arr []T) []T {

	return []T{}
}

// Percentiles returns slice-values that separate the data into 100 equal parts.
// Examples: 25%[Q1], 50%[Q2], 75%[Q3]
func Percentiles[T Number](arr []T) []T {

	return []T{}
}

// IQRange InterQuartileRange returns the difference between the first and third quartiles (Q1 and Q3)
func IQRange[T Number](arr []T) float64 {
	// sort numbers, ascending order
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })
	// Determine the Q1, Q2, Q3 and Q4 values from arr
	// IQR = Q3 - Q1
	// Determine the numbers of elements
	arrLength := len(arr)
	// Determine if the arr is even or odd
	isEven := false
	if arrLength%2 == 0 {
		isEven = true
	}

	if isEven {

	} else {

	}

	return 0.00
}

func MeanSquareError() {

}

func MeanSquareRootError() {

}

func Knn() {

}

func NBayes() {

}

func Classify() {

}
