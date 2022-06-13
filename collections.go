package mcutils

type ValueType interface {
	string | int64 | float64 | bool
}

type Number interface {
	int64 | float64
}

func Index[T ValueType](arr []T, val T) int {
	// types - string, int, float, bool
	for i, value := range arr {
		if value == val {
			return i
		}
	}
	return -1
}

func ArrayContains[T ValueType](arr []T, str T) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func ArrayFloatContains(arr []float64, str float64) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func Any[T ValueType](arr []T, val T) bool {
	for _, value := range arr {
		if value == val {
			return true
		}
	}
	return false
}

func All[T ValueType](arr []T, val T) bool {
	for _, value := range arr {
		if value != val {
			return true
		}
	}
	return false
}

func Map[T ValueType](arr []T, mapFunc func(T) T) []T {
	var mapResult []T
	for _, v := range arr {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func MapGen[T ValueType](arr []T, mapFunc func(T) T, mapChan chan<- T) {
	for _, v := range arr {
		mapChan <- mapFunc(v)
	}
	if mapChan != nil {
		close(mapChan)
	}
}

func MapInt(arr []int, mapFunc func(int) int) []int {
	var mapResult []int
	for _, v := range arr {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func MapFloat(arr []float64, mapFunc func(float64) float64) []float64 {
	var mapResult []float64
	for _, v := range arr {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func MapString(arr []string, mapFunc func(string) string) []string {
	var mapResult []string
	for _, v := range arr {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func Filter[T ValueType](arr []T, filterFunc func(T) bool) []T {
	var mapResult []T
	for _, v := range arr {
		if filterFunc(v) {
			mapResult = append(mapResult, v)
		}
	}
	return mapResult
}

func FilterGen[T ValueType](arr []T, filterFunc func(T) bool, filterChan chan<- T) {
	for _, v := range arr {
		if filterFunc(v) {
			filterChan <- v
		}
	}
	if filterChan != nil {
		close(filterChan)
	}

}

func Take[T ValueType](num uint, arr []T) []T {
	var takeResult []T
	var cnt uint = 0
	for _, v := range arr {
		if cnt == num {
			break
		}
		takeResult = append(takeResult, v)
		cnt++
	}
	return takeResult
}

func TakeGen[T ValueType](num uint, arr []T, takeChan chan<- T) {
	// use channels to implement generator to send/yield/generate num of values from arr
	var cnt uint = 0
	for _, v := range arr {
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
