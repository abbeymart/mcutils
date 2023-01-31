// @Author: abbeymart | Abi Akindele | @Created: 2021-03-03 | @Updated: 2021-03-03
// @Company: mConnect.biz | @License: MIT
// @Description: math functions

package mcutils

import (
	"fmt"
	"math"
)

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

func NaturalNumbers(count uint, cntChannel chan<- uint) {
	// use channels to implement generator to send/yield/generate natural numbers
	var cnt uint
	for cnt = 0; cnt < count; cnt++ {
		cntChannel <- cnt
	}
	if cntChannel != nil {
		close(cntChannel)
	}
}

func FactorialTail(num uint, acc uint) uint {
	acc = 1
	if num <= 1 {
		return acc
	}
	// using the tail call optimization
	return FactorialTail(num-1, num*acc)
}

func FactNumGen(num uint) chan uint {
	var factRes = make(chan uint, num)
	var x uint
	for x = 1; x <= num; x++ {
		factRes <- x
	}
	return factRes
}

func FactorialGen(num uint) uint {
	// using the generator function, via channel, no recursion
	var result uint = 1
	for v := range FactNumGen(num) {
		result *= v
	}
	return result
}

func FactorialGen2(num uint) uint {
	// using number-series, no recursion
	var result uint = 1
	var v uint
	for v = 1; v < num+1; v++ {
		result *= v
	}
	return result
}

func FiboTail(n int, current int, next int) int {
	if n == 0 {
		return current
	}
	// using the tail call optimization
	return FiboTail(n-1, current, current+next)
}

func FiboArray(num uint) [][]uint {
	// no recursion, memoization using array
	var c, d uint = 0, 1
	var result [][]uint
	var fibRes uint = 0 // track current fibo-value
	for i := 0; i < int(num); i++ {
		c, d = d, c+d
		result = append(result, []uint{c, d})
		fibRes = c
	}
	fmt.Printf("fib-result: %v", fibRes)
	return result
}

func FiboSeries(num uint) chan<- uint {
	// initial pairs / values
	var fiboChannel = make(chan uint, num) // buffered channel
	var a, b uint = 0, 1
	var i uint = 0
	for i < num {
		fiboChannel <- b
		a, b = b, a+b
		i++
	}
	return fiboChannel
}

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

func IsPrime(n int) bool {
	// prime number count algorithm condition
	s := math.Floor(math.Sqrt(float64(n)))
	for x := 2; x <= int(s); x++ {
		//Perform remainder of n for all numbers from 2 to s(short-algorithm-value)/n-1
		if n%x == 0 {
			return false
		}
	}
	return n > 1
}

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

// NaturalNumbersGen generates finite natural numbers.
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

// NaturalNumbersGenInf generates infinite natural numbers.
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
