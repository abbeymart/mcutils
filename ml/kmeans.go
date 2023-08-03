package ml

import (
	"github.com/abbeymart/mcutils"
	"math"
	"math/rand"
	"time"
)

// types | TODO: review types

type IterationLog struct {
	Centroids           [][]float64
	Iteration           int
	Error               float64
	DidReachSteadyState bool
	K                   float64 // TODO: review the type for K
	CurrentTrial        int
}

type IterationLogs = []IterationLog

type MinMaxType struct {
	Min float64
	Max float64
}

type KMeansType struct {
	K                   int
	Error               float64
	Iterations          int
	IterationLogs       IterationLogs
	Centroids           [][]float64
	CentroidAssignments []int
	Data                [][]float64
}

// helper function

// Distance function returns the distance between two points
func Distance(a []float64, b []float64) float64 {
	sumOfSquares := 0.00
	for i, aPoint := range a {
		diff := b[i] - aPoint
		sumOfSquares += diff * diff
	}
	return math.Pow(sumOfSquares, 0.5)
}

// KMeans construction function
func KMeans(k int, data [][]float64) KMeansType {
	km := KMeansType{}
	km.K = k
	km.Data = data
	km.Reset()
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
	km.CentroidAssignments = []int{}
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
		rand.NewSource(time.Now().UnixNano())
		//rand.Seed(time.Now().UnixNano())
		// For each dimension in the data find the min/max range of that dimension,
		// and choose a random value that lies within that range.
		for dimension := 0; dimension < dimensionality; dimension++ {
			minMax := dimensionRanges[dimension]
			point[dimension] = minMax.Min + (rand.Float64() * (minMax.Max - minMax.Min))
		}
		centroids = append(centroids, point)
	}
	return centroids
}

// AssignPointToCentroid method: given a point in the data to consider,
// determines the closest centroid and assign the point to that centroid.
// The return value of this method is a boolean which represents whether
// the point's centroid assignment has changed.
// It is used to determine the termination condition for the algorithm.
func (km *KMeansType) AssignPointToCentroid(pointIndex int) bool {
	lastAssignedCentroid := km.CentroidAssignments[pointIndex]
	point := km.Data[pointIndex]
	minDistance := 0.00
	assignedCentroid := 0

	for i := 0; i < len(km.Centroids); i++ {
		centroid := km.Centroids[i]
		distanceToCentroid := Distance(point, centroid)

		if minDistance == 0.00 || distanceToCentroid < minDistance {
			minDistance = distanceToCentroid
			assignedCentroid = i
		}

	}
	if assignedCentroid > 0 {
		km.CentroidAssignments[pointIndex] = assignedCentroid
	}

	return lastAssignedCentroid != assignedCentroid
}

// AssignPointsToCentroids method, for all points in the data, call AssignPointsToCentroids.
// It returns whether _any_ point's centroid assignment has been updated.
func (km *KMeansType) AssignPointsToCentroids() bool {
	didAnyPointsGetReassigned := false
	for i := 0; i < len(km.Data); i++ {
		wasReassigned := km.AssignPointToCentroid(i)
		if wasReassigned {
			didAnyPointsGetReassigned = true
		}
	}
	return didAnyPointsGetReassigned
}

// GetPointsForCentroid method, given a centroid to consider, returns an array of all points assigned to that centroid.
func (km *KMeansType) GetPointsForCentroid(centroidIndex int) [][]float64 {
	var points [][]float64
	for i := 0; i < len(km.Data); i++ {
		assignment := km.CentroidAssignments[i]
		if assignment == centroidIndex {
			points = append(points, km.Data[i])
		}
	}
	return points
}

// UpdateCentroidLocation method, given a centroid to consider, update its location to
// the mean value of the positions of points assigned to it.
func (km *KMeansType) UpdateCentroidLocation(centroidIndex int) []float64 {
	thisCentroidPoints := km.GetPointsForCentroid(centroidIndex)
	dimensionality := km.GetDimensionality()
	var newCentroid []float64

	for dimension := 0; dimension < dimensionality; dimension++ {
		var mappedValue []float64
		for _, point := range thisCentroidPoints {
			mappedValue = append(mappedValue, point[dimension])
		}
		newCentroid[dimension] = mcutils.Mean(mappedValue)
	}
	km.Centroids[centroidIndex] = newCentroid
	return newCentroid
}

// UpdateCentroidLocations method updateCentroidLocation, for all centroids.
func (km *KMeansType) UpdateCentroidLocations() {
	for i := 0; i < len(km.Centroids); i++ {
		km.UpdateCentroidLocation(i)
	}
}

// Calculates the total "error" for the current state
//of centroid positions and assignments.
//Here, error is defined as the root-mean-squared distance
//of all points to their centroids.

func (km *KMeansType) CalculateError() float64 {
	sumDistanceSquared := 0.00
	for i := 0; i < len(km.Data); i++ {
		centroidIndex := km.CentroidAssignments[i]
		centroid := km.Centroids[centroidIndex]
		point := km.Data[i]

		// Un-comment this one to do a purely geometrical error calculation
		// const thisDistance = distance(point, centroid);

		// This version also considers the number of clusters; helpful for
		// our auto-solver so that it doesn't over-fit.
		thisDistance := Distance(point, centroid) + float64(km.K)
		sumDistanceSquared += thisDistance * thisDistance
	}

	km.Error = math.Pow(sumDistanceSquared/float64(len(km.Data)), 0.5)
	return km.Error
}

// Solve method runs the k-means algorithm until either the solver reaches steady-state, or
// the maxIterations allowed has been exceeded. It returns IterationLog:
func (km *KMeansType) Solve(maxIterations int) IterationLog {
	if maxIterations < 100 {
		maxIterations = 1000
	}
	for km.Iterations < maxIterations {
		didAssignmentsChange := km.AssignPointsToCentroids()
		km.UpdateCentroidLocations()
		km.CalculateError()

		var centroidsCopy [][]float64
		_ = copy(centroidsCopy, km.Centroids)
		km.IterationLogs[km.Iterations] = IterationLog{
			Centroids:           centroidsCopy,
			Iteration:           km.Iterations,
			Error:               km.Error,
			DidReachSteadyState: !didAssignmentsChange,
		}

		if !didAssignmentsChange {
			break
		}

		km.Iterations++
	}

	return km.IterationLogs[len(km.IterationLogs)-1]

}
