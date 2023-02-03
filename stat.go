// @Author: abbeymart | Abi Akindele | @Created: 2020-12-03 | @Updated: 2020-12-03
// @Company: mConnect.biz | @License: MIT
// @Description: statistical functions

package mcutils

import (
	"fmt"
	"math"
	"sort"
)

type FrequencyValue struct {
	Label     string  `json:"label"`
	Frequency int     `json:"frequency"`
	Value     float64 `json:"value"`
}

type FrequencyResult struct {
	Result   []FrequencyValue `json:"result"`
	Interval uint             `json:"interval"`
}

type StatFrequencyValue struct {
	Label                       string  `json:"label"`
	Value                       float64 `json:"value"`
	Frequency                   int     `json:"frequency"`
	RelativeFrequency           float64 `json:"relativeFrequency"`
	CumulativeFrequency         int     `json:"cumulativeFrequency"`
	CumulativeRelativeFrequency float64 `json:"cumulativeRelativeFrequency"`
}

type StatFrequencyResult struct {
	Result   []StatFrequencyValue `json:"result"`
	Interval uint                 `json:"interval"`
}

// Mean function returns the mean or average Value from a slice of type float.
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

// Median function returns the mid or median Value from a slice of type float.
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

// Mode function returns the mode(most frequently occurring Value(s)) of a slice of type float.
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
	// compute mode/modes, i.e. modal Value greater than 1
	if max > 1 {
		for _, cVal := range arrCounters {
			if cVal.Count == max {
				modes = append(modes, cVal)
			}
		}
	}
	return modes
}

// Range function returns the range of the slice type of float.
func Range[T Number](arr []T) T {
	min, max := MinMax(arr)
	return max - min
}

// Variance function returns the variance of the mean-square-Value from a slice of type float.
func Variance[T Number](arr []T) float64 {
	meanSquareSum := 0.00
	arrLength := len(arr)
	mean := Mean(arr)
	for _, val := range arr {
		meanSquareSum += math.Pow(float64(val)-mean, 2)
	}
	return meanSquareSum / float64(arrLength)
}

// SampleStandardDeviation function returns the standard-deviation Value from a sample-data of slice of type float.
func SampleStandardDeviation[T Number](arr []T) float64 {
	deltaSquareSum := 0.00
	arrLength := len(arr)
	mean := Mean(arr)
	for _, val := range arr {
		deltaSquareSum += math.Pow(float64(val)-mean, 2)
	}
	return math.Sqrt(deltaSquareSum / float64(arrLength-1))
}

// PopulationStandardDeviation function returns the standard-deviation Value from a population/complete-data of slice of type float.
func PopulationStandardDeviation[T Number](arr []T) float64 {
	deltaSquareSum := 0.00
	arrLength := len(arr)
	mean := Mean(arr)
	for _, val := range arr {
		deltaSquareSum += math.Pow(float64(val)-mean, 2)
	}
	return math.Sqrt(deltaSquareSum / float64(arrLength))
}

// Min function returns the minimum Value from a slice of type T (int or float).
func Min[T Number](arr []T) T {
	// set initial max Value
	max := arr[0]
	// compute min and max values
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

// Max function returns the maximum Value from a slice of type T (int or float).
func Max[T Number](arr []T) T {
	// set initial max Value
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

// Interval calculates the width/interval of the sample data size
func Interval[T Number](arr []T) float64 {
	// sort numbers, ascending order
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })
	arrLength := len(arr)
	min := arr[0]
	max := arr[arrLength-1]
	rangeValue := max - min
	intervalVal := float64(rangeValue) / float64(arrLength)
	return math.Ceil(intervalVal)
}

// Frequency function returns the frequency / occurrence of a slice of type float.
func Frequency[T Number](arr []T, interval float64, valueLabel string) FrequencyResult {
	// sort numbers, ascending order
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })
	arrLength := len(arr)
	min := float64(arr[0])
	max := float64(arr[arrLength-1])
	// TODO: compose range and counts/frequency/occurrence
	if valueLabel == "" {
		valueLabel = "value"
	}
	if interval < 1 {
		interval = 1
	}
	if interval == 1 {
		// Obtain the counter values for the arr items
		result := ArrayValue[T](arr)
		arrCounters := result.Counter()
		var freqValue []FrequencyValue
		// compute the frequency/occurrence
		for _, cVal := range arrCounters {
			freqValue = append(freqValue, FrequencyValue{
				Value:     float64(cVal.Value),
				Frequency: cVal.Count,
			})
		}
	} else {
		var freqValue []FrequencyValue
		start := min
		for start <= max {
			end := start + interval
			rangeValue := fmt.Sprintf("%v<=%v<%v", start, valueLabel, end)
			// compute counts of arr values that fall within the rangeValue(start-end)
			count := 0
			for _, arrVal := range arr {
				if float64(arrVal) >= start || float64(arrVal) < end {
					count += 1
				}
			}
			freqValue = append(freqValue, FrequencyValue{
				Label:     rangeValue,
				Frequency: count,
			})
			// next range start
			start += interval
		}
	}
	return FrequencyResult{}
}

// StatFrequency function returns the frequency / relative / cumulative / relative-cumulative frequencies of a slice of type float.
func StatFrequency[T Number](arr []T, interval float64, valueLabel string) StatFrequencyResult {
	// Compute frequency values
	freqRes := Frequency(arr, interval, valueLabel)
	freqResult := freqRes.Result
	//freqResultLength := len(freqResult)
	var result []StatFrequencyValue
	// compute relative / cumulative / relative-cumulative frequencies
	// frequency/occurrence summation
	freqSum := 0
	for _, fVal := range freqResult {
		freqSum += fVal.Frequency
	}
	cumFreq := 0
	for _, val := range freqResult {
		cumFreq += val.Frequency
		result = append(result, StatFrequencyValue{
			Label:                       val.Label,
			Value:                       val.Value,
			Frequency:                   val.Frequency,
			RelativeFrequency:           float64(val.Frequency) / float64(freqSum),
			CumulativeFrequency:         cumFreq,
			CumulativeRelativeFrequency: float64(cumFreq) / float64(freqSum),
		})
	}
	return StatFrequencyResult{
		Result:   result,
		Interval: freqRes.Interval,
	}
}

// IQRange InterQuartileRange returns the difference between the first and third quartiles (Q1 and Q3),
// including quartile-values[Q0/min, Q1/25%, Q2/50%(median), Q3/75% & Q4/max].
func IQRange[T Number](arr []T) QuartilesType {
	// sort numbers, ascending order
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })
	// Determine the numbers of elements
	arrLength := len(arr)
	// minimum and maximum
	min := arr[0]
	max := arr[arrLength-1]
	// Determine the Q1, Q2, Q3 and Q4 values from arr
	Q2 := Median(arr)
	Q1 := 0.00
	Q3 := 0.00
	// Determine if the arr is even or odd
	isEven := false
	if arrLength%2 == 0 {
		isEven = true
	}
	// IQR = Q3 - Q1
	IQR := 0.00
	if isEven {
		Q1 = Median(arr[:arrLength/2])
		Q3 = Median(arr[arrLength/2:])
		IQR = Q3 - Q1
	} else {
		halfDataLength := arrLength / 2 // the ceiling value, i.e.  11, 5
		// compute medians (Q1 and Q3) to be inclusive of Q2(arr-median)
		Q1 = Median(arr[:halfDataLength+1])
		Q3 = Median(arr[halfDataLength:])
		IQR = Q3 - Q1
	}
	return QuartilesType{
		Min:   float64(min),
		Max:   float64(max), // Q4
		Range: float64(max - min),
		Q1:    Q1,
		Q2:    Q2, // Median
		Q3:    Q3,
		Q4:    float64(max),
		IQR:   IQR,
	}
}

// Deciles returns slice-values that separate the data into 10 equal parts (quantiles). TODO: review/complete.
// Examples: 10%, 20%[Q2], 30%[Q3]... 100%
func Deciles[T Number](arr []T) QuartilesType {
	// sort numbers, ascending order
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })
	// Determine the numbers of elements
	arrLength := len(arr)
	// minimum and maximum
	min := arr[0]
	max := arr[arrLength-1]
	// Determine the Q1, Q2, Q3 and Q4 values from arr
	Q2 := Median(arr)
	Q1 := 0.00
	Q3 := 0.00
	// Determine if the arr is even or odd
	isEven := false
	if arrLength%2 == 0 {
		isEven = true
	}
	// IQR = Q3 - Q1
	IQR := 0.00
	if isEven {
		Q1 = Median(arr[:arrLength/2])
		Q3 = Median(arr[arrLength/2:])
		IQR = Q3 - Q1
	} else {
		halfDataLength := arrLength / 2 // the ceiling value, i.e.  11, 5
		// compute medians (Q1 and Q3) to be inclusive of Q2(arr-median)
		Q1 = Median(arr[:halfDataLength+1])
		Q3 = Median(arr[halfDataLength:])
		IQR = Q3 - Q1
	}
	return QuartilesType{
		Min:   float64(min),
		Max:   float64(max),
		Range: float64(max - min),
		Q1:    Q1,
		Q2:    Q2,
		Q3:    Q3,
		IQR:   IQR,
	}
}

// Percentiles returns slice-values that separate the data into 100 equal parts (quantiles). TODO: review/complete.
// Examples: 1%, 2%, 3%... 100%.
func Percentiles[T Number](arr []T) QuartilesType {
	// sort numbers, ascending order
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })
	// Determine the numbers of elements
	arrLength := len(arr)
	// minimum and maximum
	min := arr[0]
	max := arr[arrLength-1]
	// Determine the Q1, Q2, Q3 and Q4 values from arr
	Q2 := Median(arr)
	Q1 := 0.00
	Q3 := 0.00
	// Determine if the arr is even or odd
	isEven := false
	if arrLength%2 == 0 {
		isEven = true
	}
	// IQR = Q3 - Q1
	IQR := 0.00
	if isEven {
		Q1 = Median(arr[:arrLength/2])
		Q3 = Median(arr[arrLength/2:])
		IQR = Q3 - Q1
	} else {
		halfDataLength := arrLength / 2 // the ceiling value, i.e.  11, 5
		// compute medians (Q1 and Q3) to be inclusive of Q2(arr-median)
		Q1 = Median(arr[:halfDataLength+1])
		Q3 = Median(arr[halfDataLength:])
		IQR = Q3 - Q1
	}
	return QuartilesType{
		Min:   float64(min),
		Max:   float64(max),
		Range: float64(max - min),
		Q1:    Q1,
		Q2:    Q2,
		Q3:    Q3,
		IQR:   IQR,
	}
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
