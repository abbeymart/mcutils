package mcutils

// Index method returns the index of the specified value (for string, it's case-insensitive)
func (arr *GenericSliceType[T]) Index(queryVal T) int {
	for i, value := range arr.value {
		if value == queryVal {
			return i
		}
	}
	return -1
}

// IndexCaseSensitive method returns the index of the specified value (for string, it's case-sensitive)
func (arr *GenericSliceType[T]) IndexCaseSensitive(val T) int {
	// types - string, int, float, bool
	for i, value := range arr.value {
		if value == val {
			return i
		}
	}
	return -1
}

// ArrayContains method check if a slice of generic type T contains/includes a value of type T  (for string, it's case-insensitive)
func (arr *GenericSliceType[T]) ArrayContains(queryVal T) bool {
	for _, a := range arr.value {
		if a == queryVal {
			return true
		}
	}
	return false
}

// ArrayContainsCaseSensitive check if a slice of generic type T contains/includes a value of type T  (for string, it's case-sensitive)
func (arr *GenericSliceType[T]) ArrayContainsCaseSensitive(str T) bool {
	for _, a := range arr.value {
		if a == str {
			return true
		}
	}
	return false
}

// ArrayStringContains method check if a slice of string contains/includes a string value, case-insensitive
func (arr *StringSliceType) ArrayStringContains(queryVal string) bool {
	for _, a := range arr.value {
		if a == queryVal {
			return true
		}
	}
	return false
}

// ArrayStringContainsCaseSensitive method check if a slice of string contains/includes a string value, case-sensitive
func (arr *StringSliceType) ArrayStringContainsCaseSensitive(val string) bool {
	for _, a := range arr.value {
		if a == val {
			return true
		}
	}
	return false
}

// ArrayIntContains method check if a slice of int contains/includes an int value
func (arr *IntSliceType) ArrayIntContains(queryVal int) bool {
	for _, a := range arr.value {
		if a == queryVal {
			return true
		}
	}
	return false
}

// ArrayFloatContains method check if a slice of int contains/includes a float value
func (arr *FloatSliceType) ArrayFloatContains(queryVal float64) bool {
	for _, a := range arr.value {
		if a == queryVal {
			return true
		}
	}
	return false
}

// Any method determines if one or more of the slice-values satisfies the testFunc.
func (arr *GenericSliceType[T]) Any(testFunc TestFuncType[T]) bool {
	for _, value := range arr.value {
		if testFunc(value) {
			return true
		}
	}
	return false
}

// All method determines if all the slice-values satisfies the testFunc.
func (arr *GenericSliceType[T]) All(testFunc TestFuncType[T]) bool {
	for _, value := range arr.value {
		if !testFunc(value) {
			return false
		}
	}
	return true
}

// include result address/pointer as method parameter to improve performance

// Map method returns the mapped-slice-value, of type T, based on the mapFunc [generic].
func (arr *GenericSliceType[T]) Map(mapFunc func(T) T, result []T) {
	result = []T{}
	for _, v := range arr.value {
		result = append(result, mapFunc(v))
	}
}

// MapGen returns series of the mapped-value, of type T, based on the mapFunc [generic].
func (arr *GenericSliceType[T]) MapGen(mapFunc func(T) T, mapChan chan<- T) {
	for _, v := range arr.value {
		mapChan <- mapFunc(v)
	}
	if mapChan != nil {
		close(mapChan)
	}
}

// MapInt method returns the mapped-slice-value of type int, based on the mapFunc.
func (arr *IntSliceType) MapInt(mapFunc func(int) int, result []int) {
	result = []int{}
	for _, v := range arr.value {
		result = append(result, mapFunc(v))
	}
}

// MapFloat method returns the mapped-slice-value, of type float64, based on the mapFunc.
func (arr *FloatSliceType) MapFloat(mapFunc func(float64) float64, result []float64) {
	result = []float64{}
	for _, v := range arr.value {
		result = append(result, mapFunc(v))
	}
}

// MapString method returns the mapped-slice-value, of type string, based on the mapFunc.
func (arr *StringSliceType) MapString(mapFunc func(string) string, result []string) {
	result = []string{}
	for _, v := range arr.value {
		result = append(result, mapFunc(v))
	}
}

// Filter method returns the filtered-slice-value, of type T, based on the filterFunc [generic].
func (arr *GenericSliceType[T]) Filter(filterFunc func(T) bool, result []T) {
	result = []T{}
	for _, v := range arr.value {
		if filterFunc(v) {
			result = append(result, v)
		}
	}
}

// FilterGen method returns series filtered-value, of type T, based on the filterFunc [generic].
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

// Take method returns num of the arr slice of type T [generic].
func (arr *GenericSliceType[T]) Take(num uint, result []T) {
	result = []T{}
	var cnt uint = 0
	for _, v := range arr.value {
		if cnt == num {
			break
		}
		result = append(result, v)
		cnt++
	}
}

// TakeGen method returns num series of values, of type T [generic].
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

// ReverseArray method returns the reverse values of the specified array/slice [generic type]
func (arr *GenericSliceType[T]) ReverseArray(result []T) {
	// arr and arrChan must be of the same type: int, float, string, bool
	result = []T{}
	for i := len(arr.value) - 1; i >= 0; i-- {
		result = append(result, arr.value[i])
	}
}

// ReverseArrayInt method returns the reverse values of the specified array/slice of int.
func (arr *IntSliceType) ReverseArrayInt(result []int) {
	result = []int{}
	for i := len(arr.value) - 1; i >= 0; i-- {
		result = append(result, arr.value[i])
	}
}

// ReverseArrayFloat returns the reverse values of the specified array/slice of float64.
func (arr *FloatSliceType) ReverseArrayFloat(result []float64) {
	result = []float64{}
	for i := len(arr.value) - 1; i >= 0; i-- {
		result = append(result, arr.value[i])
	}
}

// ReverseArrayGenerator sequentially generates reverse values of the specified array/slice [generic]
func (arr *GenericSliceType[T]) ReverseArrayGenerator(arrChan chan<- T) {
	// arr and arrChan must be of the same type: int, float
	for i := len(arr.value) - 1; i >= 0; i-- {
		arrChan <- arr.value[i]
	}
	// ends send task to the channel
	if arrChan != nil {
		close(arrChan)
	}
}

// ReverseArrayIntGen sequentially generates reverse values of the specified array/slice of int.
func (arr *IntSliceType) ReverseArrayIntGen(arrChan chan<- int) {
	for i := len(arr.value) - 1; i >= 0; i-- {
		arrChan <- arr.value[i]
	}
	// ends send task to the channel
	if arrChan != nil {
		close(arrChan)
	}
}

// ReverseArrayFloatGen sequentially generates reverse values of the specified array/slice of float64.
func (arr *FloatSliceType) ReverseArrayFloatGen(arrChan chan<- float64) {
	for i := len(arr.value) - 1; i >= 0; i-- {
		arrChan <- arr.value[i]
	}
	// ends send task to the channel
	if arrChan != nil {
		close(arrChan)
	}
}
