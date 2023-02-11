package data

import "github.com/abbeymart/mcutils"

// stats

const MeanResult = 5.5
const MedianResult = 5.5
const MinResult = 1
const MaxResult = 10

var MinMaxResult = []float64{1, 10}

const StdDeviationResult = 3.0276503540974917 // 16 decimal places
const StdDeviationResultEst = 3.02765

// populationStandardDeviation

var ArrayOfNumber2 = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5, 9.5}

const StdDeviationResultEst2 = 2.74

// geometric mean

const GeoMeanPrecision2Result = 4.53    // precision 5
const GeoMeanPrecision5Result = 4.52872 // precision 5

// variance

const VariancePrecision2Result = 9.17    // precision of 2
const VariancePrecision5Result = 9.16667 // precision of 2

// interval

const Interval = 1

// frequency

var ArrayOfNumberFreq = []int{1, 2, 2, 2, 1, 4, 4, 3, 5, 10}

// frequencyStat

var FrequencyStatResult = mcutils.StatFrequencyResult{
	Interval: 1,
	Result: []mcutils.StatFrequencyValue{
		{
			Label:                       "",
			Value:                       1,
			Frequency:                   2,
			RelativeFrequency:           1.00,
			CumulativeFrequency:         1,
			CumulativeRelativeFrequency: 1.00,
		},
		{
			Label:                       "",
			Value:                       2,
			Frequency:                   3,
			RelativeFrequency:           1.00,
			CumulativeFrequency:         1,
			CumulativeRelativeFrequency: 1.00,
		},
		{
			Label:                       "",
			Value:                       3,
			Frequency:                   1,
			RelativeFrequency:           1.00,
			CumulativeFrequency:         1,
			CumulativeRelativeFrequency: 1.00,
		},
		{
			Label:                       "",
			Value:                       4,
			Frequency:                   2,
			RelativeFrequency:           1.00,
			CumulativeFrequency:         1,
			CumulativeRelativeFrequency: 1.00,
		},
		{
			Label:                       "",
			Value:                       5,
			Frequency:                   1,
			RelativeFrequency:           1.00,
			CumulativeFrequency:         1,
			CumulativeRelativeFrequency: 1.00,
		},
		{
			Label:                       "",
			Value:                       10,
			Frequency:                   1,
			RelativeFrequency:           1.00,
			CumulativeFrequency:         1,
			CumulativeRelativeFrequency: 1.00,
		},
	},
}

var FrequencyStatResultLabel = mcutils.StatFrequencyResult{
	Interval: 2,
	Result: []mcutils.StatFrequencyValue{
		{
			Label:                       "1<=value<3",
			Value:                       1,
			Frequency:                   2,
			RelativeFrequency:           1.00,
			CumulativeFrequency:         1,
			CumulativeRelativeFrequency: 1.00,
		},
		{
			Label:                       "",
			Value:                       2,
			Frequency:                   3,
			RelativeFrequency:           1.00,
			CumulativeFrequency:         1,
			CumulativeRelativeFrequency: 1.00,
		},
		{
			Label:                       "",
			Value:                       3,
			Frequency:                   1,
			RelativeFrequency:           1.00,
			CumulativeFrequency:         1,
			CumulativeRelativeFrequency: 1.00,
		},
		{
			Label:                       "",
			Value:                       4,
			Frequency:                   2,
			RelativeFrequency:           1.00,
			CumulativeFrequency:         1,
			CumulativeRelativeFrequency: 1.00,
		},
		{
			Label:                       "",
			Value:                       5,
			Frequency:                   1,
			RelativeFrequency:           1.00,
			CumulativeFrequency:         1,
			CumulativeRelativeFrequency: 1.00,
		},
		{
			Label:                       "",
			Value:                       10,
			Frequency:                   1,
			RelativeFrequency:           1.00,
			CumulativeFrequency:         1,
			CumulativeRelativeFrequency: 1.00,
		},
	},
}

// IQRange

var IQRange = mcutils.QuartilesType{
	Min:   1,
	Max:   10, // Q4
	Range: 9,
	Q1:    2.5,
	Q2:    5.5, // Median
	Q3:    7.7,
	Q4:    10,
	IQR:   1,
}
