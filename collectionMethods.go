package mcutils

// methods

func (arr *GenericSliceType[T]) Index(queryVal T) int {
	for i, value := range arr.value {
		if value == queryVal {
			return i
		}
	}
	return -1
}

func (arr *GenericSliceType[T]) ArrayContains(queryVal T) bool {
	for _, a := range arr.value {
		if a == queryVal {
			return true
		}
	}
	return false
}

func (arr *StringSliceType) ArrayStringContains(queryVal string) bool {
	for _, a := range arr.value {
		if a == queryVal {
			return true
		}
	}
	return false
}

func (arr *IntSliceType) ArrayIntContains(queryVal int) bool {
	for _, a := range arr.value {
		if a == queryVal {
			return true
		}
	}
	return false
}

func (arr *FloatSliceType) ArrayFloatContains(queryVal float64) bool {
	for _, a := range arr.value {
		if a == queryVal {
			return true
		}
	}
	return false
}

func (arr *GenericSliceType[T]) Any(testFunc TestFuncType[T]) bool {
	for _, value := range arr.value {
		if testFunc(value) {
			return true
		}
	}
	return false
}

func (arr *GenericSliceType[T]) All(testFunc TestFuncType[T]) bool {
	for _, value := range arr.value {
		if !testFunc(value) {
			return false
		}
	}
	return true
}

// TODO: include result address/pointer as method parameter to improve performance

func (arr *GenericSliceType[T]) Map(mapFunc func(T) T) []T {
	var mapResult []T
	for _, v := range arr.value {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func (arr *GenericSliceType[T]) MapGen(mapFunc func(T) T, mapChan chan<- T) {
	for _, v := range arr.value {
		mapChan <- mapFunc(v)
	}
	if mapChan != nil {
		close(mapChan)
	}
}

func (arr *IntSliceType) MapInt(mapFunc func(int) int) []int {
	var mapResult []int
	for _, v := range arr.value {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func (arr *FloatSliceType) MapFloat(mapFunc func(float64) float64) []float64 {
	var mapResult []float64
	for _, v := range arr.value {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func (arr *StringSliceType) MapString(mapFunc func(string) string) []string {
	var mapResult []string
	for _, v := range arr.value {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func (arr *GenericSliceType[T]) Filter(filterFunc func(T) bool) []T {
	var mapResult []T
	for _, v := range arr.value {
		if filterFunc(v) {
			mapResult = append(mapResult, v)
		}
	}
	return mapResult
}

func (arr *GenericSliceType[T]) FilterGen(filterFunc func(T) bool, filterChan chan<- T) {
	for _, v := range arr.value {
		if filterFunc(v) {
			filterChan <- v
		}
	}
	if filterChan != nil {
		close(filterChan)
	}

}

func (arr *GenericSliceType[T]) Take(num uint) []T {
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

func (arr *GenericSliceType[T]) TakeGen(num uint, takeChan chan<- T) {
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

// TODO: reverse array methods
