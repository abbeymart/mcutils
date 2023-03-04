package ml

type KMeansAutoSolverType struct {
	KMin      float64
	KMax      float64
	MaxTrials int
	Data      [][]float64
	Best      IterationLog
	Log       []IterationLog
}

// KMeansAutoSolver construction function
func KMeansAutoSolver(kMin float64, kMax float64, maxTrials int, data [][]float64) KMeansAutoSolverType {
	if kMin < 1.00 {
		kMin = 1.00
	}
	if kMax < 5.00 {
		kMax = 5.00
	}
	if maxTrials < 5 {
		maxTrials = 5
	}
	kmAuto := KMeansAutoSolverType{}
	kmAuto.KMin = kMin
	kmAuto.KMax = kMax
	kmAuto.MaxTrials = maxTrials
	kmAuto.Data = data
	kmAuto.Reset()
	return kmAuto
}

func (kmAuto *KMeansAutoSolverType) Reset() {
	kmAuto.Best = IterationLog{}
	kmAuto.Log = []IterationLog{}
}

func (kmAuto *KMeansAutoSolverType) Solve(maxIterations int) IterationLog {
	if maxIterations < 100 {
		maxIterations = 1000
	}

	for k := kmAuto.KMin; k < kmAuto.KMax; k++ {
		for currentTrial := 0; currentTrial < kmAuto.MaxTrials; currentTrial++ {
			solver := KMeans(int(k), kmAuto.Data)
			// Add k and currentTrial number to the solution before logging
			solution := solver.Solve(maxIterations)
			solution.K = k
			solution.CurrentTrial = currentTrial
			kmAuto.Log = append(kmAuto.Log, solution)

			if kmAuto.Best.Error == 0.00 || solution.Error < kmAuto.Best.Error {
				kmAuto.Best = solution
			}
		}
	}

	return kmAuto.Best

}
