package ml

import (
	"github.com/abbeymart/mcutils"
	"math"
)

// types

type IterationLog struct {
	Centroids           [][]float64
	Iteration           uint
	Error               int
	DidReachSteadyState bool
}

type IterationLogs = []IterationLog

type SolutionType struct {
	IterationLog
	K            int
	CurrentTrial int
}

type MinMaxType struct {
	Min float64
	Max float64
}

type KMeansType struct {
	K                   int
	Error               int
	Iterations          int
	IterationLogs       IterationLogs
	Centroids           [][]float64
	CentroidAssignments []float64
	Data                [][]float64
}

// KMeans construction function

func KMeans(k int, data [][]float64) KMeansType {
	km := KMeansType{}
	km.K = k
	km.Data = data

	return km
}

// Reset method resets the solver state.
// Use this if you wish to run the same solver instance again with the same data points,
// with different initial conditions.
func (km *KMeansType) Reset() {
	km.Error = 0
	km.Iterations = 0
	km.IterationLogs = []IterationLog{}
	km.Centroids = [][]float64{} // TODO: create initRandomCentroids method
	km.CentroidAssignments = []float64{}
}

// GetDimensionality method determines the number of dimensions in the data set.
func (km *KMeansType) GetDimensionality() int {
	point := km.Data[0]
	return len(point)
}

// GetRangeForDimension method, for a given dimension in the data set, determines the minimum and maximum value.
// This is used during random initialization to make sure the random centroids are in the same range as the data.
func (km *KMeansType) GetRangeForDimension(n uint) MinMaxType {
	var values []float64
	for _, point := range km.Data {
		values = append(values, point[n])
	}
	return MinMaxType{
		Min: mcutils.Min(values),
		Max: mcutils.Max(values),
	}
}

// GetAllDimensionRanges method get ranges for all dimensions.
// It returns array whose indices are the dimension number and whose members are the output of getRangeForDimension.
func (km *KMeansType) GetAllDimensionRanges() []MinMaxType {
	var dimensionRanges []MinMaxType
	dimensionality := km.GetDimensionality()
	for dimension := 0; dimension < dimensionality; dimension++ {
		dimensionRanges[dimension] = km.GetRangeForDimension(uint(dimension))
	}
	return dimensionRanges
}

// initRandomCentroids method initializes random centroids, using the ranges of the data to set minimum and maximum bounds for the centroids.
// You may inspect the output of this method if you need to debug random initialization, otherwise this is an internal method.
func (km *KMeansType) initRandomCentroids() [][]float64 {
	dimensionality := km.GetDimensionality()
	dimensionRanges := km.GetAllDimensionRanges()
	var centroids [][]float64

	// create centroids
	for i := 0; i < km.K; i++ {
		// for each dimension with its own range, create a placeholder
		var point []float64
		// For each dimension in the data find the min/max range of that dimension,
		// and choose a random value that lies within that range.
		for dimension := 0; dimension < dimensionality; dimension++ {
			minMax := dimensionRanges[dimension]
			point[dimension] = minMax.Min + (math.random() * (minMax.Max - minMax.Min))
		}

		centroids = append(centroids, point)
	}
	return centroids
}
