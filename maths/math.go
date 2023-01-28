// @Author: abbeymart | Abi Akindele | @Created: 2021-03-03 | @Updated: 2021-03-03
// @Company: mConnect.biz | @License: MIT
// @Description: math functions

package maths

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
