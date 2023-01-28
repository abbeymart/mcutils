// @Author: abbeymart | Abi Akindele | @Created: 2020-12-03 | @Updated: 2020-12-03
// @Company: mConnect.biz | @License: MIT
// @Description: statistical functions

package stats

import (
	"math"
	"sort"
)

func Mean(arr []float64) float64 {
	var sum = 0.00
	arrLength := len(arr)
	for _, v := range arr {
		sum += v
	}
	return sum / float64(arrLength)
}

func Median(arr []float64) float64 {
	// sort numbers, ascending order
	sort.Float64s(arr)
	arrLength := len(arr)
	if arrLength%2 != 0 {
		// if count is odd
		medianIndex := math.Floor(float64(arrLength / 2))
		return arr[uint(medianIndex)]
	}
	// if count is even
	medianIndex1 := (arrLength / 2) - 1
	medianIndex2 := arrLength / 2
	return (arr[medianIndex1] + arr[medianIndex2]) / 2
}

func StandardDeviation(arr []float64) float64 {
	deltaSquareSum := 0.00
	arrLength := len(arr)
	mean := Mean(arr)
	for _, val := range arr {
		deltaSquareSum += math.Pow(val-mean, 2)
	}
	return math.Sqrt(deltaSquareSum / float64(arrLength-1))
}

func Min(arr []float64) float64 {
	// sort numbers, ascending order
	sort.Float64s(arr)
	return arr[0]
}

func Max(arr []float64) float64 {
	// sort numbers, ascending order
	sort.Float64s(arr)
	return arr[len(arr)-1]
}

func MinMax(arr []float64) []float64 {
	// sort numbers, ascending order
	sort.Float64s(arr)
	return []float64{arr[0], arr[len(arr)-1]}
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
