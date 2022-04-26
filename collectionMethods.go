package mcutils

// types

type McObjectString struct {
	val string
}
type McObjectFloat struct {
	val float64
}
type McObjectBool struct {
	val bool
}
type McObjectInt struct {
	val int
}
type McObjectInterface struct {
	val interface{}
}
type McObjectStringSlice struct {
	val []string
}
type McObjectFloatSlice struct {
	val []float64
}
type McObjectBoolSlice struct {
	val []bool
}
type McObjectIntSlice struct {
	val []int
}
type McObjectInterfaceSlice struct {
	val []interface{}
}

// methods

func (arr *McObjectInterfaceSlice) Index(queryVal interface{}) int {
	for i, value := range arr.val {
		if value == queryVal {
			return i
		}
	}
	return -1
}

func (arr *McObjectInterfaceSlice) ArrayContains(queryVal interface{}) bool {
	for _, a := range arr.val {
		if a == queryVal {
			return true
		}
	}
	return false
}

func (arr *McObjectStringSlice) ArrayStringContains(queryVal string) bool {
	for _, a := range arr.val {
		if a == queryVal {
			return true
		}
	}
	return false
}

func (arr *McObjectIntSlice) ArrayIntContains(queryVal int) bool {
	for _, a := range arr.val {
		if a == queryVal {
			return true
		}
	}
	return false
}

func (arr *McObjectFloatSlice) ArrayFloatContains(queryVal float64) bool {
	for _, a := range arr.val {
		if a == queryVal {
			return true
		}
	}
	return false
}

func (arr *McObjectInterfaceSlice) Any(queryVal interface{}) bool {
	for _, value := range arr.val {
		if value == queryVal {
			return true
		}
	}
	return false
}

func (arr *McObjectInterfaceSlice) All(val interface{}) bool {
	for _, value := range arr.val {
		if value != val {
			return true
		}
	}
	return false
}

func (arr *McObjectInterfaceSlice) Map(mapFunc func(interface{}) interface{}) []interface{} {
	var mapResult []interface{}
	for _, v := range arr.val {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func (arr McObjectInterfaceSlice) MapGen(mapFunc func(interface{}) interface{}, mapChan chan<- interface{}) {
	for _, v := range arr.val {
		mapChan <- mapFunc(v)
	}
	if mapChan != nil {
		close(mapChan)
	}
}

func (arr McObjectIntSlice) MapInt(mapFunc func(int) int) []int {
	var mapResult []int
	for _, v := range arr.val {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func (arr McObjectFloatSlice) MapFloat(mapFunc func(float64) float64) []float64 {
	var mapResult []float64
	for _, v := range arr.val {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func (arr McObjectStringSlice) MapString(mapFunc func(string) string) []string {
	var mapResult []string
	for _, v := range arr.val {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func (arr McObjectInterfaceSlice) Filter(filterFunc func(interface{}) bool) []interface{} {
	var mapResult []interface{}
	for _, v := range arr.val {
		if filterFunc(v) {
			mapResult = append(mapResult, v)
		}
	}
	return mapResult
}

func (arr McObjectInterfaceSlice) FilterGen(filterFunc func(interface{}) bool, filterChan chan<- interface{}) {
	for _, v := range arr.val {
		if filterFunc(v) {
			filterChan <- v
		}
	}
	if filterChan != nil {
		close(filterChan)
	}

}

func (arr McObjectInterfaceSlice) Take(num uint) chan<- interface{} {
	// use channels to implement generator to send/yield/generate num of values from arr
	// buffered channel with capacity of number of values to take
	var takeChannel = make(chan interface{}, num)
	var cnt uint = 0
	for _, v := range arr.val {
		if cnt == num {
			break
		}
		takeChannel <- v
		cnt++
	}
	close(takeChannel)
	return takeChannel
}

func (arr McObjectInterfaceSlice) TakeGen(num uint, takeChan chan<- interface{}) {
	// use channels to implement generator to send/yield/generate num of values from arr
	var cnt uint = 0
	for _, v := range arr.val {
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
