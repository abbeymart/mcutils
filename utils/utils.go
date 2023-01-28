// @Author: abbeymart | Abi Akindele | @Created: 2020-11-20 | @Updated: 2020-12-02
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect shared utility functions

package utils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/abbeymart/mcresponse"
	"github.com/abbeymart/mcutils"
	"github.com/abbeymart/mcutils/collections"
	"github.com/asaskevich/govalidator"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func CamelCase(text string, sep string) (string, error) {
	// validate acceptable separators (" ", "_", "__", ".", "|")
	sepArr := mcutils.PermittedSeparators
	if !collections.ArrayContains(sepArr, sep) {
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
	sepArr := mcutils.PermittedSeparators
	if !ArrayStringContains(sepArr, sep) {
		textMsg := fmt.Sprintf("missing or unacceptable separator: %v | Acceptable-separators: %v", sep, strings.Join(sepArr, ", "))
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
	var defaultLang = mcutils.DefaultLanguage
	// Set defaultLang to current userLang, set from the UI
	if userLang != "" {
		defaultLang = userLang
	}
	return defaultLang
}

func getLocale(localeFiles mcutils.Locale, options mcutils.LocaleOptions) mcutils.LocaleContent {
	// localeType := options.LocaleType
	var language string
	if lang := options.Language; lang != "" {
		language = lang
	} else {
		language = mcutils.DefaultLanguage
	}
	// set the locale file contents
	myLocale := localeFiles[language]
	return myLocale
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

// counters

type ArrayValue []interface{}
type ArrayOfString []string
type ArrayOfInt []int
type ArrayOfFloat []float64
type DataCount map[string]int

func (val ArrayValue) counterGeneric() map[interface{}]int {
	var count = make(map[interface{}]int)
	for _, it := range val {
		if v, ok := count[it]; ok && v > 0 {
			count[it] = v + 1
		} else {
			count[it] = 1
		}
	}
	return count
}

func (val ArrayValue) counter() DataCount {
	var count = make(map[string]int)
	for _, it := range val {
		// stringify it=>keys
		var jsonVal, _ = json.Marshal(it)
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
	// refactor, using counter method
	var count = val.counter()
	// compute set values
	setValue := make([]string, len(count))
	for keyValue := range count {
		setValue = append(setValue, keyValue)
	}
	return setValue
}

func (val ArrayOfString) setOfString() []string {
	var count = make(map[string]int)
	for _, itVal := range val {
		if v, ok := count[itVal]; ok && v > 0 {
			count[itVal] = v + 1
		} else {
			count[itVal] = 1
		}
	}
	// compute set values
	setValue := make([]string, len(count))
	for keyValue := range count {
		setValue = append(setValue, keyValue)
	}
	return setValue
}

func (val ArrayOfInt) setOfInt() []int {
	var count = make(map[int]int)
	for _, itVal := range val {
		if v, ok := count[itVal]; ok && v > 0 {
			count[itVal] = v + 1
		} else {
			count[itVal] = 1
		}
	}
	// compute set values
	setValue := make([]int, len(count))
	for keyValue := range count {
		setValue = append(setValue, keyValue)
	}
	return setValue
}

func (val ArrayOfFloat) setOfFloat() []float64 {
	var count = make(map[float64]int)
	for _, itVal := range val {
		if v, ok := count[itVal]; ok && v > 0 {
			count[itVal] = v + 1
		} else {
			count[itVal] = 1
		}
	}
	// compute set values
	setValue := make([]float64, len(count))
	for keyValue := range count {
		setValue = append(setValue, keyValue)
	}
	return setValue
}

// Collections

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
	// channel may be closed on by the requester
	for cnt := 0; ; cnt++ {
		naturalChan <- uint(cnt)
	}
}

type EmailUserNameType struct {
	Email    string
	Username string
}

// EmailUsername processes and returns the loginName as email or username
func EmailUsername(loginName string) EmailUserNameType {
	if govalidator.IsEmail(loginName) {
		return EmailUserNameType{
			Email:    loginName,
			Username: "",
		}
	}

	return EmailUserNameType{
		Email:    "",
		Username: loginName,
	}

}

func TypeOf(rec interface{}) reflect.Type {
	return reflect.TypeOf(rec)
}

// ParseRawValues process the raw rows/records from SQL-query
func ParseRawValues(rawValues [][]byte) ([]interface{}, error) {
	// variables
	var value interface{}
	var values []interface{}
	// parse the current-raw-values
	for _, val := range rawValues {
		if err := json.Unmarshal(val, &value); err != nil {
			return nil, errors.New(fmt.Sprintf("Error parsing raw-row-value: %v", err.Error()))
		} else {
			values = append(values, value)
		}
	}
	return values, nil
}

// ArrayStringContains check if a slice of string contains/includes a string value
func ArrayStringContains(arr []string, val string) bool {
	for _, a := range arr {
		if strings.ToLower(a) == strings.ToLower(val) {
			return true
		}
	}
	return false
}

// ArrayIntContains check if a slice of int contains/includes an int value
func ArrayIntContains(arr []int, val int) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}

// ArrayToSQLStringValues transforms a slice of string to SQL-string-formatted-values
func ArrayToSQLStringValues(arr []string) string {
	result := ""
	for ind, val := range arr {
		result += "'" + val + "'"
		if ind < len(arr)-1 {
			result += ", "
		}
	}
	return result
}

// JsonToStruct converts json inputs to equivalent struct data type specification
// rec must be a pointer to a type matching the jsonRec
func JsonToStruct(jsonRec []byte, rec interface{}) error {
	if err := json.Unmarshal(jsonRec, &rec); err == nil {
		return nil
	} else {
		return errors.New(fmt.Sprintf("Error converting json-to-record-format: %v", err.Error()))
	}
}

type ActionParamType map[string]interface{}
type ActionParamsType []ActionParamType

// DataToValueParam accepts only a struct type/model and returns the ActionParamType
// data camel/Pascal-case keys are converted to underscore-keys to match table-field/columns specs
func DataToValueParam(rec interface{}) (ActionParamType, error) {
	// validate recs as struct{} type
	recType := fmt.Sprintf("%v", reflect.TypeOf(rec).Kind())
	switch recType {
	case "struct":
		dataValue := ActionParamType{}
		v := reflect.ValueOf(rec)
		typeOfS := v.Type()

		for i := 0; i < v.NumField(); i++ {
			dataValue[govalidator.CamelCaseToUnderscore(typeOfS.Field(i).Name)] = v.Field(i).Interface()
			//fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).ItemName, v.Field(i).Interface())
		}
		return dataValue, nil
	default:
		return nil, errors.New(fmt.Sprintf("rec parameter must be of type struct{}"))
	}
}

// StructToMap function converts struct to map
func StructToMap(rec interface{}) (map[string]interface{}, error) {
	// validate recs as struct{} type
	recType := fmt.Sprintf("%v", reflect.TypeOf(rec).Kind())
	switch recType {
	case "struct":
		break
	default:
		return nil, errors.New(fmt.Sprintf("rec parameter must be of type struct{}"))
	}
	var mapData map[string]interface{}
	// json record
	jsonRec, err := json.Marshal(rec)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error computing struct to map: %v", err.Error()))
	}
	// json-to-map
	err = json.Unmarshal(jsonRec, &mapData)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error computing struct to map: %v", err.Error()))
	}
	return mapData, nil
}

// TagField return the field-tag (e.g. table-column-name) for mcorm tag
func TagField(rec interface{}, fieldName string, tag string) (string, error) {
	// validate recs as struct{} type
	t := reflect.TypeOf(rec)
	recType := fmt.Sprintf("%v", t.Kind())
	switch recType {
	case "struct":
		break
	default:
		return "", errors.New(fmt.Sprintf("rec parameter must be of type struct{}"))
	}
	// convert the first-letter to upper-case (public field)
	field, found := t.FieldByName(strings.Title(fieldName))
	if !found {
		// check private field
		field, found = t.FieldByName(fieldName)
		if !found {
			return "", errors.New(fmt.Sprintf("error retrieving tag-field for field-name: %v", fieldName))
		}
	}
	//tagValue := field.Tag
	return field.Tag.Get(tag), nil
}

// StructToTagMap function converts struct to map (tag/underscore_field), for crud-db-table-record
func StructToTagMap(rec interface{}, tag string) (map[string]interface{}, error) {
	// validate recs as struct{} type
	recType := fmt.Sprintf("%v", reflect.TypeOf(rec).Kind())
	switch recType {
	case "struct":
		break
	default:
		return nil, errors.New(fmt.Sprintf("rec parameter must be of type struct{}"))
	}
	tagMapData := map[string]interface{}{}
	mapData, err := StructToMap(rec)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error computing struct to map: %v", err.Error()))
	}
	// compose tagMapData
	for key, val := range mapData {
		tagField, tagErr := TagField(rec, key, tag)
		if tagErr != nil {
			return nil, errors.New(fmt.Sprintf("error computing tag-field: %v", tagErr.Error()))
		}
		tagMapData[tagField] = val
	}
	return tagMapData, nil
}

func ToCamelCase(text string, sep string) string {
	// accept words/text and separator(' ', '_', '__', '.')
	textArray := strings.Split(text, sep)
	// convert the first word to lowercase
	firstWord := strings.ToLower(textArray[0])
	// convert other words: first letter to upper case and other letters to lowercase
	remWords := textArray[1:]
	var otherWords []string
	for _, item := range remWords {
		// convert first letter to upper case
		item0 := strings.ToUpper(string(item[0]))
		// convert other letters to lowercase
		item1N := strings.ToLower(item[1:])
		itemString := fmt.Sprintf("%v%v", item0, item1N)
		otherWords = append(otherWords, itemString)
	}
	return fmt.Sprintf("%v%v", firstWord, strings.Join(otherWords, ""))
}

// StructToMapUnderscore converts struct to map (underscore_fields), for crud-db-table-record
func StructToMapUnderscore(rec interface{}) (map[string]interface{}, error) {
	// validate recs as struct{} type
	recType := fmt.Sprintf("%v", reflect.TypeOf(rec).Kind())
	switch recType {
	case "struct":
		break
	default:
		return nil, errors.New(fmt.Sprintf("rec parameter must be of type struct{}"))
	}

	caseUnderscoreMapData := map[string]interface{}{}
	mapData, err := StructToMap(rec)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error computing struct to map: %v", err.Error()))
	}
	// compose caseUnderscoreMapData
	for key, val := range mapData {
		caseUnderscoreMapData[govalidator.CamelCaseToUnderscore(key)] = val
	}
	return caseUnderscoreMapData, nil
}

// MapToMapUnderscore converts map camelCase-fields to underscore-fields
func MapToMapUnderscore(rec interface{}) (map[string]interface{}, error) {
	// validate recs as map type
	recMap, ok := rec.(map[string]interface{})
	if !ok || recMap == nil {
		return nil, errors.New(fmt.Sprintf("rec parameter must be of type map[string]interface{}"))
	}

	uMapData := map[string]interface{}{}
	// compose uMapData
	for key, val := range recMap {
		uMapData[govalidator.CamelCaseToUnderscore(key)] = val
	}
	return uMapData, nil
}

// MapToMapCamelCase converts map underscore-fields to camelCase-fields
func MapToMapCamelCase(rec interface{}, sep string) (map[string]interface{}, error) {
	// validate recs as map type
	recMap, ok := rec.(map[string]interface{})
	if !ok || recMap == nil {
		return nil, errors.New(fmt.Sprintf("rec parameter must be of type map[string]interface{}"))
	}

	uMapData := map[string]interface{}{}
	// compose uMapData
	for key, val := range recMap {
		uMapData[ToCamelCase(key, sep)] = val
	}
	return uMapData, nil
}

// ArrayMapToMapUnderscore converts []map-fields to underscore
func ArrayMapToMapUnderscore(rec interface{}) ([]map[string]interface{}, error) {
	// validate recs as []map type
	arrayMap, ok := rec.([]map[string]interface{})
	if !ok || arrayMap == nil {
		return nil, errors.New(fmt.Sprintf("rec parameter must be of type []map[string]interface{}"))
	}

	var uArrayMapData []map[string]interface{}
	// compose underscoreMapData
	for _, mapRec := range arrayMap {
		uMapData, err := MapToMapUnderscore(mapRec)
		if err != nil {
			return nil, err
		}
		uArrayMapData = append(uArrayMapData, uMapData)
	}

	return uArrayMapData, nil
}

// StructToFieldValues converts struct to record fields(underscore) and associated values (columns and values)
func StructToFieldValues(rec interface{}) ([]string, []interface{}, error) {
	// validate recs as struct{} type
	recType := fmt.Sprintf("%v", reflect.TypeOf(rec).Kind())
	switch recType {
	case "struct":
		break
	default:
		return nil, nil, errors.New(fmt.Sprintf("rec parameter must be of type struct{}"))
	}
	var tableFields []string
	var fieldValues []interface{}
	mapDataValue, err := StructToMap(rec)
	if err != nil {
		return nil, nil, errors.New("error computing struct to map")
	}
	// compose table fields/column(underscore) and values
	for key, val := range mapDataValue {
		tableFields = append(tableFields, govalidator.CamelCaseToUnderscore(key))
		fieldValues = append(fieldValues, val)
	}
	return tableFields, fieldValues, nil
}

// ArrayMapToStruct converts []map/actParams to []struct/model-type
func ArrayMapToStruct(actParams ActionParamsType, recs interface{}) (interface{}, error) {
	// validate recs as slice / []struct{} type
	recsType := fmt.Sprintf("%v", reflect.TypeOf(recs).Kind())
	switch recsType {
	case "slice":
		break
	default:
		return nil, errors.New(fmt.Sprintf("recs parameter must be of type []struct{}: %v", recsType))
	}
	switch rType := recs.(type) {
	case []interface{}:
		for i, val := range rType {
			// validate each record as struct type
			recType := fmt.Sprintf("%v", reflect.TypeOf(val).Kind())
			switch recType {
			case "struct":
				break
			default:
				return nil, errors.New(fmt.Sprintf("recs[%v] parameter must be of type struct{}: %v", i, recType))
			}
		}
	default:
		return nil, errors.New(fmt.Sprintf("rec parameter must be of type []struct{}: %v", rType))
	}
	// compute json records from actParams
	jsonRec, err := json.Marshal(actParams)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error computing map to struct records: %v", err.Error()))
	}
	// transform json records to []struct{} (recs)
	err = json.Unmarshal(jsonRec, &recs)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error computing map to struct records: %v", err.Error()))
	}
	return recs, nil
}

// MapToStruct converts map to struct
func MapToStruct(mapRecord map[string]interface{}, rec interface{}) (interface{}, error) {
	// validate rec as struct{} type
	recType := fmt.Sprintf("%v", reflect.TypeOf(rec).Kind())
	switch recType {
	case "struct":
		break
	default:
		return nil, errors.New(fmt.Sprintf("rec parameter must be of type struct{}"))
	}
	// compute json records from actParams (map-record)
	jsonRec, err := json.Marshal(mapRecord)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error computing map to struct records: %v", err.Error()))
	}
	// transform json record to struct{} (rec)
	err = json.Unmarshal(jsonRec, &rec)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error computing map to struct records: %v", err.Error()))
	}
	return rec, nil
}

// GetParamsMessage compose the message-object into mcresponse.ResponseMessage
func GetParamsMessage(msgObject mcutils.MessageObject) mcresponse.ResponseMessage {
	var messages = ""

	for key, val := range msgObject {
		if messages != "" {
			messages = fmt.Sprintf("%v | %v: %v", messages, key, val)
		} else {
			messages = fmt.Sprintf("%v: %v", key, val)
		}
	}
	return mcresponse.GetResMessage("validateError", mcresponse.ResponseMessageOptions{
		Message: messages,
		Value:   nil,
	})
}

// ConvertJsonStringToMapValue converts the db-json-string-value to the map-type
func ConvertJsonStringToMapValue(jsonStr string) (map[string]interface{}, error) {
	mapVal := map[string]interface{}{}
	jErr := json.Unmarshal([]byte(jsonStr), &mapVal)
	if jErr != nil {
		return nil, jErr
	}
	return mapVal, nil
}

// ConvertJsonStringToTypeValue converts the db-json-string-value to the base-type
func ConvertJsonStringToTypeValue(jsonStr string, typePointer interface{}) (interface{}, error) {
	jErr := json.Unmarshal([]byte(jsonStr), typePointer)
	if jErr != nil {
		return nil, jErr
	}
	return typePointer, nil
}

// ConvertJsonBase64StringToTypeValue converts the db-json-string-value to the base-type
func ConvertJsonBase64StringToTypeValue(base64Str interface{}, typePointer interface{}) (interface{}, error) {
	// assert the base64String value as of string-type
	strVal, ok := base64Str.(string)
	if !ok {
		return nil, errors.New(fmt.Sprintf("unable to convert base64-string [%v] to string", base64Str))
	}
	// decode the base64StringValue
	decoded, err := base64.StdEncoding.DecodeString(strVal)
	if err != nil {
		return nil, err
	}
	// transform/un-marshal the decoded value to the base-type
	jErr := json.Unmarshal(decoded, typePointer)
	if jErr != nil {
		return nil, jErr
	}
	return typePointer, nil
}

// ConvertJsonBase64StringToMap converts the db-json-string-value to the map-type
func ConvertJsonBase64StringToMap(base64Str interface{}) (map[string]interface{}, error) {
	mapVal := map[string]interface{}{}
	strVal, ok := base64Str.(string)
	if !ok {
		return nil, errors.New(fmt.Sprintf("unable to convert base64-string [%v] to string", base64Str))
	}
	decoded, err := base64.StdEncoding.DecodeString(strVal)
	if err != nil {
		return nil, err
	}
	jErr := json.Unmarshal(decoded, &mapVal)
	if jErr != nil {
		return nil, jErr
	}
	return mapVal, nil
}

func ConvertByteSliceToBase64Str(fileContent []byte) string {
	return base64.StdEncoding.EncodeToString(fileContent)
}

func ConvertStringToBase64Str(fileContent string) string {
	return base64.StdEncoding.EncodeToString([]byte(fileContent))
}

func ExcludeEmptyIdFromMapRecord(rec ActionParamType) ActionParamType {
	mapVal := ActionParamType{}
	for key, val := range rec {
		if key == "id" && val == "" {
			continue
		}
		mapVal[key] = val
	}
	return mapVal
}

// ExcludeFieldFromMapRecord exclude id and accessKey fields
func ExcludeFieldFromMapRecord(rec ActionParamType, field string) ActionParamType {
	mapVal := ActionParamType{}
	for key, val := range rec {
		if key == field {
			continue
		}
		mapVal[key] = val
	}
	return mapVal
}

func ExcludeEmptyIdFields(recs []ActionParamType) []ActionParamType {
	var mapValues []ActionParamType
	for _, rec := range recs {
		mapVal := ActionParamType{}
		for key, val := range rec {
			if (key == "id" || strings.HasSuffix(key, "Id")) && (val == nil || val == "") {
				continue
			}
			mapVal[key] = val
		}
		mapValues = append(mapValues, mapVal)
	}
	return mapValues
}

func StructToMapToCamelCase(rec interface{}, sep string) (map[string]interface{}, error) {
	mapVal, mErr := StructToMap(rec)
	if mErr != nil {
		return nil, mErr
	}
	val, err := MapToMapCamelCase(mapVal, sep)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// ComputeTaskDuration computes the task interval in microseconds
func ComputeTaskDuration(start time.Time, end time.Time) int64 {
	return end.Sub(start).Microseconds()
}

// RandomString generates random string of characters and numbers
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

// RandomNumbers generates random numbers using rand.Perm and returns []int as string
func RandomNumbers(n int) string {
	rand.Seed(time.Now().UnixNano())
	v := rand.Perm(n)
	var vString []string
	for _, item := range v {
		vString = append(vString, fmt.Sprintf("%v", item))
	}
	return fmt.Sprintf("%v", strings.Join(vString, ""))
}

// ValidateSubActionParams validates that subscriber-appIds includes actionParam-appId, for save - create/update tasks
func ValidateSubActionParams(actParams ActionParamsType, subAppIds []string) bool {
	result := false
	for _, rec := range actParams {
		id, idOk := rec["appId"].(string)
		if !idOk || !ArrayStringContains(subAppIds, id) {
			return false
		}
		result = true
	}
	return result
}
