// @Author: abbeymart | Abi Akindele | @Created: 2020-11-20 | @Updated: 2020-12-02
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect shared utility functions

package mcutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/abbeymart/mcresponse"
	"github.com/abbeymart/mctypes"
	"math"
	"strconv"
	"strings"
	"time"
)

func ArrayContains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func ArrayContain(arr []interface{}, str interface{}) bool {
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

func CamelCase(text string, sep string) (string, error) {
	// validate acceptable separators (" ", "_", "__", ".", "|")
	sepArr := []string{" ", "_", "__", ".", "|"}
	if !ArrayContains(sepArr, sep) {
		textMsg := fmt.Sprintf("missing or unacceptable separator: %v", sep)
		return text, errors.New(textMsg)
	}
	// split text by separator/sep
	textArray := strings.Split(text, sep)
	// convert the first word to lowercase
	firstWord := strings.ToLower(textArray[0])
	// convert other words: first letter to upper case and other letters to lowercase
	remainingWords := textArray[1:]
	var otherWords string
	for _, v := range remainingWords {
		firstLetterUpper := strings.ToUpper(string(v[0]))
		remainLetterLower := strings.ToLower(v[1:])
		otherWords += firstLetterUpper + remainLetterLower
	}

	return fmt.Sprintf("%v%v", firstWord, otherWords), nil
}

func PascalCase(text string, sep string) (string, error) {
	// validate acceptable separators (" ", "_", "__", ".", "|")
	sepArr := []string{" ", "_", "__", ".", "|"}
	if !ArrayContains(sepArr, sep) {
		textMsg := fmt.Sprintf("missing or unacceptable separator: %v", sep)
		return text, errors.New(textMsg)
	}
	// split text by separator/sep
	textArray := strings.Split(text, sep)
	// convert all words: first letter to upper case and other letters to lowercase
	var allWords string
	for _, v := range textArray {
		firstLetterUpper := strings.ToUpper(string(v[0]))
		remainLetterLower := strings.ToLower(v[1:])
		allWords += firstLetterUpper + remainLetterLower
	}

	return fmt.Sprintf("%v", allWords), nil
}

func LeapYear(year int) bool {
	// by setting the day to the 29th and checking if the day remains
	return year%400 == 0 || (year%4 == 0 && year%100 != 0) || time.Date(year, time.February, 29, 23, 0, 0, 0, time.UTC).Day() == 29
}

func GetLanguage(userLang string) string {
	// Define/set default language variable
	var defaultLang = DefaultLanguage
	// Set defaultLang to current userLang, set from the UI
	if userLang != "" {
		defaultLang = userLang
	}
	return defaultLang
}

func getLocale(localeFiles Locale, options LocaleOptions) LocaleContent {
	// localeType := options.LocaleType
	var language string
	if lang := options.Language; lang != "" {
		language = lang
	} else {
		language = DefaultLanguage
	}
	// set the locale file contents
	myLocale := localeFiles[language]
	return myLocale
}

func GetParamsMessage(msgObject mctypes.MessageObject, msgType string) mcresponse.ResponseMessage {
	var messages = ""

	for key, val := range msgObject {
		if messages != "" {
			messages = fmt.Sprintf("%v | %v : %v", messages, key, val)
		} else {
			messages = fmt.Sprintf("%v : %v", key, val)
		}
	}
	if msgType == "" {
		msgType = "unknown"
	}
	return mcresponse.GetResMessage(msgType, mcresponse.ResponseMessageOptions{
		Message: messages,
		Value:   nil,
	})
}

func ShortString(str string, maxLength uint) string {
	if len(str) > int(maxLength) {
		// return slice of the string, up to/including the maxLength, and append "..."
		return str[:int(maxLength)+1] + "..."
	}
	// return whole string
	return str
}

func StringToBool(val string) bool {
	// convert val to lowercase
	strVal := strings.ToLower(val)
	// perform the conversion
	if strVal == "true" || strVal == "t" || strVal == "yes" || strVal == "y" {
		return true
	} else if intVal, err := strconv.Atoi(strVal); err == nil && intVal > 0 {
		return true
	} else {
		return false
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

func FactGen(num uint) chan uint {
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
	for v := range FactGen(num) {
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
	var fibArr []uint
	var fibRes uint = 0 // track current fibo-value
	for i := 0; i < int(num); i++ {
		c, d = d, c+d
		fibArr = append(fibArr, c, d)
		result = append(result, fibArr)
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

func IsPrime(n uint) bool {
	// prime number count algorithm condition
	s := math.Floor(math.Sqrt(float64(n)))
	for x := 2; x <= int(s); x++ {
		//Perform remainder of n for all numbers from 2 to s(short-algorithm-value)/n-1
		if n%uint(x) == 0 {
			return false
		}
	}
	return n > 1
}

func ReverseArray(arr []interface{}) []interface{} {
	// arr and arrChan must be of the same type: int, float
	var reverseArray []interface{}
	for i := len(arr) - 1; i >= 0; i-- {
		reverseArray = append(reverseArray, arr[i])
	}
	return reverseArray
}

func ReverseArrayInt(arr []int) []int {
	var reverseArray []int
	for i := len(arr) - 1; i >= 0; i-- {
		reverseArray = append(reverseArray, arr[i])
	}
	return reverseArray
}

func ReverseArrayFloat(arr []float64) []float64 {
	var reverseArray []float64
	for i := len(arr) - 1; i >= 0; i-- {
		reverseArray = append(reverseArray, arr[i])
	}
	return reverseArray
}

func ReverseArrayGen(arr []interface{}, arrChan chan interface{}) {
	// arr and arrChan must be of the same type: int, float
	for i := len(arr) - 1; i >= 0; i-- {
		arrChan <- arr[i]
	}
}

func ReverseArrayIntGen(arr []int, arrChan chan int) {
	for i := len(arr) - 1; i >= 0; i-- {
		arrChan <- arr[i]
	}
}

func ReverseArrayFloatGen(arr []float64, arrChan chan float64) {
	for i := len(arr) - 1; i >= 0; i-- {
		arrChan <- arr[i]
	}
}

func Pythagoras(limit uint) [][]uint {
	var pResult [][]uint
	var pArr []uint
	var a, b uint
	for a = 1; a <= limit; a++ {
		for b = a; b <= limit; b++ {
			itemSqrt := math.Sqrt(float64(a*a + b*b))
			if uint(itemSqrt)%1.00 == 0 {
				pArr := append(pArr, a, b, uint(itemSqrt))
				pResult = append(pResult, pArr)
			}
		}
	}
	return pResult
}

func PythagorasGen(limit uint, pythagorasChan chan []uint) {
	var pArr []uint
	var a, b uint
	for a = 1; a <= limit; a++ {
		for b = a; b <= limit; b++ {
			itemSqrt := math.Sqrt(float64(a*a + b*b))
			if uint(itemSqrt)%1.00 == 0 {
				pArr := append(pArr, a, b, uint(itemSqrt))
				pythagorasChan <- pArr
			}
		}
	}
	if pythagorasChan != nil {
		close(pythagorasChan)
	}
}

// counter
type ArrayValue []interface{}
type ArrayOfString []string
type ArrayOfInt []int
type ArrayOfFloat []float64
type DataCount map[string]int

func (val ArrayValue) counter() DataCount {
	var count = make(map[string]int)
	for _, val := range val {
		var jsonVal, _ = json.Marshal(val)
		var countKey = string(jsonVal)
		if v, ok := count[countKey]; ok && v > 0 {
			count[countKey] = v + 1
		} else {
			count[countKey] = 1
		}
	}
	return count
}

func (val ArrayValue) set() []string {
	var count = make(map[string]int)
	for _, itVal := range val {
		var jsonVal, _ = json.Marshal(itVal)
		var countKey = string(jsonVal)
		if v, ok := count[countKey]; ok && v > 0 {
			count[countKey] = v + 1
		} else {
			count[countKey] = 1
		}
	}
	// compute set values
	setValue := make([]string, len(count))
	for keyValue, _ := range count {
		setValue = append(setValue, keyValue)
	}
	return setValue
}

func (val ArrayOfString) setOfString() []string {
	var count = make(map[string]int)
	for _, itVal := range val {
		var jsonVal, _ = json.Marshal(itVal)
		var countKey = string(jsonVal)
		if v, ok := count[countKey]; ok && v > 0 {
			count[countKey] = v + 1
		} else {
			count[countKey] = 1
		}
	}
	// compute set values
	setValue := make([]string, len(count))
	for keyValue, _ := range count {
		setValue = append(setValue, keyValue)
	}
	return setValue
}

func (val ArrayOfInt) setOfInt() []int {
	var count = make(map[int]int)
	for _, itVal := range val {
		//var jsonVal, _ = json.Marshal(itVal)
		//var countKey = string(jsonVal)
		if v, ok := count[itVal]; ok && v > 0 {
			count[itVal] = v + 1
		} else {
			count[itVal] = 1
		}
	}
	// compute set values
	setValue := make([]int, len(count))
	for keyValue, _ := range count {
		setValue = append(setValue, keyValue)
	}
	return setValue
}

func (val ArrayOfFloat) setOfFloat() []float64 {
	var count = make(map[float64]int)
	for _, itVal := range val {
		//var jsonVal, _ = json.Marshal(itVal)
		//var countKey = string(jsonVal)
		if v, ok := count[itVal]; ok && v > 0 {
			count[itVal] = v + 1
		} else {
			count[itVal] = 1
		}
	}
	// compute set values
	setValue := make([]float64, len(count))
	for keyValue, _ := range count {
		setValue = append(setValue, keyValue)
	}
	return setValue
}

// Collections
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

func Take(num uint, arr []interface{}) chan<- interface{} {
	// use channels to implement generator to send/yield/generate num of values from arr
	// buffered channel with capacity of number of values to take
	var takeChannel = make(chan interface{}, num)
	var cnt uint = 0
	for _, v := range arr {
		if cnt == num {
			break
		}
		takeChannel <- v
		cnt++
	}
	close(takeChannel)
	return takeChannel
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

func NaturalNumbers(count uint) chan<- uint {
	// use channels to implement generator to send/yield/generate natural numbers
	// buffered channel with capacity of the count of natural numbers to generate
	var cntChannel = make(chan uint, count)
	var cnt uint
	for cnt = 0; cnt < count; cnt++ {
		cntChannel <- cnt
	}
	close(cntChannel)
	return cntChannel
}

// Finite natural numbers generation
func NaturalNumbersGen(count uint, naturalChan chan<- uint) {
	// use channels to implement generator to yield/generate natural numbers
	var cnt uint
	for cnt = 0; cnt < count; cnt++ {
		naturalChan <- cnt
	}
	if naturalChan != nil {
		close(naturalChan)
	}
}

// Infinite natural numbers generation
func NaturalNumbersGenInf(naturalChan chan<- uint) {
	// use channels to implement generator to yield/generate natural numbers
	for cnt := 0; ; cnt++ {
		naturalChan <- uint(cnt)
	}
}
