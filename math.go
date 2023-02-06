// @Author: abbeymart | Abi Akindele | @Created: 2021-03-03 | @Updated: 2021-03-03
// @Company: mConnect.biz | @License: MIT
// @Description: math functions

package mcutils

import (
	"math"
)

// Fibos function returns slice of fibonacci numbers from 1 up to the specified limit (num).
func Fibos(num uint) []uint {
	var fiboArray = []uint{1, 1}
	var i uint = 0
	for i < num {
		var prev = fiboArray[len(fiboArray)-1]
		var prev2 = fiboArray[len(fiboArray)-2]
		fiboArray = append(fiboArray, prev+prev2)
		i++
	}
	return fiboArray
}

// NaturalNumbers function generates series positive numbers from 0 up to specified limit (num).
func NaturalNumbers(num uint, cntChannel chan<- uint) {
	// use channels to implement generator to send/yield/generate natural numbers
	var cnt uint
	for cnt = 0; cnt < num; cnt++ {
		cntChannel <- cnt
	}
	if cntChannel != nil {
		close(cntChannel)
	}
}

// FactorialTail function returns the factorial Value from 1 to the specified limit (num).
// Accumulator Value should be set to 1 (default)
func FactorialTail(num uint, acc uint) uint {
	if acc != 1 {
		acc = 1
	}
	if num <= 1 {
		return acc
	}
	// using the tail call optimization
	return FactorialTail(num-1, num*acc)
}

// FactNumGen function generates series of the factorial Value of 1 to num.
func FactNumGen(num uint, factChannel chan<- int) {
	var x int
	for x = 1; x <= int(num); x++ {
		factChannel <- x
	}
	if factChannel != nil {
		close(factChannel)
	}
}

// Factorial function returns the factorial Value of 1 to num, using number-generator via channel.
func Factorial(num uint) int {
	// using the generator function, via channel, no recursion
	result := 1
	factChannel := make(chan int, num)
	go FactNumGen(num, factChannel)
	for v := range factChannel {
		result *= v
	}
	return result
}

// FactorialGen2 function returns the factorial Value of 1 to num, using simple iteration method.
func FactorialGen2(num uint) int {
	// using number-series, no recursion
	result := 1
	var v int
	for v = 1; v <= int(num); v++ {
		result *= v
	}
	return result
}

// FiboTail function returns last fibonacci numbers up to the limit (num).
// current and next parameters should be set to 0/1? (default) - TODO: review/clarify/test.
func FiboTail(num int, current int, next int) int {
	if current != 0 {
		current = 0
	}
	if num == 0 {
		return current
	}

	// using the tail call optimization
	return FiboTail(num-1, current, current+next)
}

// FiboArray function returns the slice of pairs of fibonacci numbers.
func FiboArray(num uint, result [][]uint) {
	// no recursion, memoization using array
	result = [][]uint{}
	var a, b uint = 0, 1
	for i := 0; i < int(num); i++ {
		a, b = b, a+b
		result = append(result, []uint{a, b})
	}
}

// FiboSeries function generates series of fibonacci numbers up to the specified limit (num).
func FiboSeries(num uint, fiboChannel chan<- uint) {
	// initial pairs / values
	var a, b uint = 0, 1
	var i uint = 0
	for i < num {
		fiboChannel <- b
		a, b = b, a+b
		i++
	}
	if fiboChannel != nil {
		close(fiboChannel)
	}
}

// PrimeNumbers function returns the prime numbers up to the limit if num.
func PrimeNumbers(num int) (pNums []int) {
next:
	for outer := 2; outer < num; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
			pNums = append(pNums, outer)
		}
	}
	return pNums
}

// IsPrime function determines if the number(n) is a prime number.
func IsPrime(n int) bool {
	// prime number count algorithm condition
	s := math.Floor(math.Sqrt(float64(n)))
	for x := 2; x <= int(s); x++ {
		//Perform remainder of n for all numbers from 2 to s(short-algorithm-Value)/n-1
		if n%x == 0 {
			return false
		}
	}
	return n > 1
}

// Pythagoras function returns the number pairs and hypothesis values.
func Pythagoras(limit uint) [][]uint {
	var pResult [][]uint
	var a, b uint
	for a = 1; a <= limit; a++ {
		for b = a; b <= limit; b++ {
			itemSqrt := math.Sqrt(float64(a*a + b*b))
			if uint(itemSqrt)%1.00 == 0 || uint(itemSqrt)%1.00 == 0.00 {
				pResult = append(pResult, []uint{a, b, uint(itemSqrt)})
			}
		}
	}
	return pResult
}

// PythagorasGen function generates series the number pairs and hypothesis values.
func PythagorasGen(limit uint, pythagorasChan chan []uint) {
	var a, b uint
	for a = 1; a <= limit; a++ {
		for b = a; b <= limit; b++ {
			itemSqrt := math.Sqrt(float64(a*a + b*b))
			if uint(itemSqrt)%1.00 == 0 || uint(itemSqrt)%1.00 == 0.00 {
				pythagorasChan <- []uint{a, b, uint(itemSqrt)}
			}
		}
	}
	if pythagorasChan != nil {
		close(pythagorasChan)
	}
}

// NaturalNumbersGen function generates finite natural numbers.
func NaturalNumbersGen(num uint, naturalChan chan<- uint) {
	// use channels to implement generator to yield/generate finite natural numbers
	var cnt uint
	for cnt = 0; cnt < num; cnt++ {
		naturalChan <- cnt
	}
	if naturalChan != nil {
		close(naturalChan)
	}
}

// NaturalNumbersGenInf function generates infinite natural numbers.
func NaturalNumbersGenInf(naturalChan chan<- uint, stopFunc func() bool) {
	// use channels to implement generator to yield/generate infinite natural numbers
	for cnt := 0; ; cnt++ {
		naturalChan <- uint(cnt)
		if stopFunc() {
			break
		}
	}
	if naturalChan != nil {
		close(naturalChan)
	}
}
