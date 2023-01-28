package collections

// methods

func (arr *mcutils.McObjectInterfaceSlice[T]) Index(queryVal T) int {
	for i, value := range arr.value {
		if value == queryVal {
			return i
		}
	}
	return -1
}

func (arr *mcutils.McObjectInterfaceSlice[T]) ArrayContains(queryVal T) bool {
	for _, a := range arr.value {
		if a == queryVal {
			return true
		}
	}
	return false
}

func (arr *mcutils.McObjectStringSlice) ArrayStringContains(queryVal string) bool {
	for _, a := range arr.value {
		if a == queryVal {
			return true
		}
	}
	return false
}

func (arr *mcutils.McObjectIntSlice) ArrayIntContains(queryVal int) bool {
	for _, a := range arr.value {
		if a == queryVal {
			return true
		}
	}
	return false
}

func (arr *mcutils.McObjectFloatSlice) ArrayFloatContains(queryVal float64) bool {
	for _, a := range arr.value {
		if a == queryVal {
			return true
		}
	}
	return false
}

func (arr *mcutils.McObjectInterfaceSlice[T]) Any(testFunc mcutils.TestFuncType[T]) bool {
	for _, value := range arr.value {
		if testFunc(value) {
			return true
		}
	}
	return false
}

func (arr *mcutils.McObjectInterfaceSlice[T]) All(testFunc mcutils.TestFuncType[T]) bool {
	for _, value := range arr.value {
		if !testFunc(value) {
			return false
		}
	}
	return true
}

func (arr *mcutils.McObjectInterfaceSlice[T]) Map(mapFunc func(T) T) []T {
	var mapResult []T
	for _, v := range arr.value {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func (arr *mcutils.McObjectInterfaceSlice[T]) MapGen(mapFunc func(T) T, mapChan chan<- T) {
	for _, v := range arr.value {
		mapChan <- mapFunc(v)
	}
	if mapChan != nil {
		close(mapChan)
	}
}

func (arr *mcutils.McObjectIntSlice) MapInt(mapFunc func(int) int) []int {
	var mapResult []int
	for _, v := range arr.value {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func (arr *mcutils.McObjectFloatSlice) MapFloat(mapFunc func(float64) float64) []float64 {
	var mapResult []float64
	for _, v := range arr.value {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func (arr *mcutils.McObjectStringSlice) MapString(mapFunc func(string) string) []string {
	var mapResult []string
	for _, v := range arr.value {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func (arr *mcutils.McObjectInterfaceSlice[T]) Filter(filterFunc func(T) bool) []T {
	var mapResult []T
	for _, v := range arr.value {
		if filterFunc(v) {
			mapResult = append(mapResult, v)
		}
	}
	return mapResult
}

func (arr *mcutils.McObjectInterfaceSlice[T]) FilterGen(filterFunc func(T) bool, filterChan chan<- T) {
	for _, v := range arr.value {
		if filterFunc(v) {
			filterChan <- v
		}
	}
	if filterChan != nil {
		close(filterChan)
	}

}

func (arr *mcutils.McObjectInterfaceSlice[T]) Take(num uint) []T {
	var takeResult []T
	var cnt uint = 0
	for _, v := range arr.value {
		if cnt == num {
			break
		}
		takeResult = append(takeResult, v)
		cnt++
	}
	return takeResult
}

func (arr *mcutils.McObjectInterfaceSlice[T]) TakeGen(num uint, takeChan chan<- T) {
	// use channels to implement generator to send/yield/generate num of values from arr
	var cnt uint = 0
	for _, v := range arr.value {
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
