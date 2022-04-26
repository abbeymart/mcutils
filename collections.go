package mcutils

func Index(arr []interface{}, val interface{}) int {
	// types - string, int, float, bool
	for i, value := range arr {
		if value == val {
			return i
		}
	}
	return -1
}

func ArrayContains(arr []interface{}, str interface{}) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func ArrayStringContains(arr []string, val string) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}

func ArrayIntContains(arr []int, val int) bool {
	for _, a := range arr {
		if a == val {
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

func Any(arr []interface{}, val interface{}) bool {
	for _, value := range arr {
		if value == val {
			return true
		}
	}
	return false
}

func All(arr []interface{}, val interface{}) bool {
	for _, value := range arr {
		if value != val {
			return true
		}
	}
	return false
}

func Map(arr []interface{}, mapFunc func(interface{}) interface{}) []interface{} {
	var mapResult []interface{}
	for _, v := range arr {
		mapResult = append(mapResult, mapFunc(v))
	}
	return mapResult
}

func MapGen(arr []interface{}, mapFunc func(interface{}) interface{}, mapChan chan<- interface{}) {
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

func Filter(arr []interface{}, filterFunc func(interface{}) bool) []interface{} {
	var mapResult []interface{}
	for _, v := range arr {
		if filterFunc(v) {
			mapResult = append(mapResult, v)
		}
	}
	return mapResult
}

func FilterGen(arr []interface{}, filterFunc func(interface{}) bool, filterChan chan<- interface{}) {
	for _, v := range arr {
		if filterFunc(v) {
			filterChan <- v
		}
	}
	if filterChan != nil {
		close(filterChan)
	}

}

func Take(num uint, arr []interface{}) []interface{} {
	var takeResult []interface{}
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

func TakeGen(num uint, arr []interface{}, takeChan chan<- interface{}) {
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
