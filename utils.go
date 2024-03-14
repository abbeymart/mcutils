// @Author: abbeymart | Abi Akindele | @Created: 2020-11-20 | @Updated: 2020-12-02
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect shared utility functions

package mcutils

import (
	cryptoRand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/abbeymart/mcresponse"
	"github.com/leekchan/accounting"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/big"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// counters

type CounterValue[T ValueType] struct {
	Count int `json:"count"`
	Value T   `json:"value"`
}

type CounterObjectValue[T map[string]interface{} | struct{}] struct {
	Count int
	Value T
}

type ArrayValue[T ValueType] []T

//type ArrayValue[T ValueType] struct {
//	Value []T
//}

type ArrayOfString []string
type ArrayOfInt []int
type ArrayOfFloat []float64
type DataCount map[string]int
type SliceObjectType[T map[string]interface{}] []T
type CounterResult[T ValueType] map[string]CounterValue[T]
type ObjectCounterResult[T map[string]interface{} | struct{}] map[string]CounterObjectValue[T]

// Counter method returns the unique counts of the specified array/slice values[int, float, string and bool]
func (val ArrayValue[T]) Counter() CounterResult[T] {
	var count = make(CounterResult[T])
	for _, it := range val {
		// stringify it=>key
		var itStr = fmt.Sprintf("%v", it)
		if v, ok := count[itStr]; ok && v.Count > 0 {
			count[itStr] = CounterValue[T]{
				Count: v.Count + 1,
				Value: it,
			}
		} else {
			count[itStr] = CounterValue[T]{
				Count: 1,
				Value: it,
			}
		}
	}
	return count
}

// SliceObjectCounter method returns the unique counts of the specified array/slice of map[string]interface{} values
func (val SliceObjectType[T]) SliceObjectCounter() ObjectCounterResult[T] {
	var count ObjectCounterResult[T]
	for _, it := range val {
		// stringify it=>key
		jsonVal, _ := json.Marshal(it)
		var itStr = string(jsonVal)
		if v, ok := count[itStr]; ok && v.Count > 0 {
			count[itStr] = CounterObjectValue[T]{
				Count: v.Count + 1,
				Value: it,
			}
		} else {
			count[itStr] = CounterObjectValue[T]{
				Count: 1,
				Value: it,
			}
		}
	}
	return count
}

// Set method returns the slice of set values for a generic type T
func (val ArrayValue[T]) Set() []T {
	// refactor, using counter method
	var count = val.Counter()
	// compute set values
	setValue := make([]T, len(count))
	for _, value := range count {
		setValue = append(setValue, value.Value)
	}
	return setValue
}

func (val ArrayOfString) SetOfString() []string {
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

func (val ArrayOfInt) SetOfInt() []int {
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

func (val ArrayOfFloat) SetOfFloat() []float64 {
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

// SeparatorFieldToCamelCase transforms a separated/underscore name to camel-case name
func SeparatorFieldToCamelCase(text string, sep string) (string, error) {
	// validate acceptable separators [" ", "_", "__", ".", "|"]
	sepArr := PermittedSeparators
	if !ArrayContains(sepArr, sep) {
		return "", errors.New(fmt.Sprintf("missing or unacceptable separator: %v | Acceptable-separators: %v", sep, strings.Join(sepArr, ", ")))
	}
	// split text by separator/sep
	textArray := strings.Split(text, sep)
	// convert the first word to lowercase
	firstWord := strings.ToLower(textArray[0])
	// convert other words: first letter to upper case and other letters to lowercase
	remainingWords := textArray[1:]
	otherWords := ""
	for _, v := range remainingWords {
		// transform first letter to upper case and other letters to lowercase
		otherWords += strings.ToUpper(string(v[0])) + strings.ToLower(v[1:])
	}

	return fmt.Sprintf("%v%v", firstWord, otherWords), nil
}

// SeparatorFieldToPascalCase transforms a separated/underscore field-name to pascal-case
func SeparatorFieldToPascalCase(text string, sep string) (string, error) {
	// validate acceptable separators [" ", "_", "__", ".", "|"]
	sepArr := PermittedSeparators
	if !ArrayContains(sepArr, sep) {
		return "", errors.New(fmt.Sprintf("missing or unacceptable separator: %v | Acceptable-separators: %v", sep, strings.Join(sepArr, ", ")))
	}
	// split text by separator/sep
	textArray := strings.Split(text, sep)
	// convert all words: transform first letter to upper case and other letters to lowercase
	allWords := ""
	for _, v := range textArray {
		// transform first letter to upper case and other letters to lowercase
		allWords += strings.ToUpper(string(v[0])) + strings.ToLower(v[1:])
	}

	return fmt.Sprintf("%v", allWords), nil
}

// CaseFieldToUnderscore transforms camelCase or PascalCase name to underscore name, in lowercase
func CaseFieldToUnderscore(caseString string) string {
	// Create slice of words from the cased-Value, separate at Uppercase-character
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)
	// transform first character to Uppercase
	caseValue := strings.ToUpper(string(caseString[0])) + caseString[1:]
	// compose separate/matched words as slice
	textArray := re.FindAllString(caseValue, -1)
	var wordsArray []string
	for _, v := range textArray {
		wordsArray = append(wordsArray, strings.ToLower(v))
	}
	if len(wordsArray) < 1 {
		return ""
	}
	if len(wordsArray) == 1 {
		return wordsArray[0]
	}
	return strings.Join(wordsArray, "_")
}

// IsLeapYear determines if a given year is a leap year
func IsLeapYear(year int) bool {
	// by setting the day to the 29th and checking if the day remains
	return year%400 == 0 || (year%4 == 0 && year%100 != 0) || time.Date(year, time.February, 29, 23, 0, 0, 0, time.UTC).Day() == 29
}

// GetLanguage returns the specified user-language or defaultLanguage as default
func GetLanguage(userLang string) string {
	// Define/set default language variable
	defaultLang := DefaultLanguage
	// Set defaultLang to current userLang, set from the UI
	if userLang != "" {
		defaultLang = userLang
	}
	return defaultLang
}

// GetLocale function returns the locale for the specified language
func GetLocale(localeFiles LocaleFilesType, options LocaleOptions) Locale {
	// localeType := options.LocaleType
	language := ""
	if lang := options.Language; lang != "" {
		language = lang
	} else {
		language = DefaultLanguage
	}
	// set the locale file contents
	myLocale := localeFiles[language]
	return myLocale
}

// ShortString returns part of string, based on the specified maxLength
func ShortString(str string, maxLength uint) string {
	if len(str) > int(maxLength) {
		// return slice of the string, up to/including the maxLength, and append "..."
		return str[:int(maxLength)+1] + "..."
	}
	// return whole string, if len(str) <= maxLength
	return str
}

// StringToBool returns the boolean Value of the specified string-Value
func StringToBool(val string) bool {
	// convert val to lowercase and trim any trailing whitespaces
	strVal := strings.ToLower(val)
	strVal = strings.Trim(strVal, " ")
	// perform the conversion
	if strVal == "true" || strVal == "t" || strVal == "yes" || strVal == "y" {
		return true
	} else if intVal, err := strconv.Atoi(strVal); err == nil && intVal > 0 {
		return true
	} else {
		return false
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
			return nil, errors.New(fmt.Sprintf("Error parsing raw-row-Value: %v", err.Error()))
		} else {
			values = append(values, value)
		}
	}
	return values, nil
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

func AccessAllowed(accessList []string, item string) bool {
	return ArrayContains(accessList, item)
}

func EncryptPassword(newPassword string, cryptCode string) (string, error) {
	pwd := []byte(newPassword + cryptCode)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(userPassword string, inputPassword string, cryptCode string) (bool, error) {
	inputPwd := []byte(inputPassword + cryptCode)
	userPwd := []byte(userPassword)
	err := bcrypt.CompareHashAndPassword(userPwd, inputPwd)
	if err != nil {
		log.Println(err)
		return false, errors.New(fmt.Sprintf("password comparison error: %v", err.Error()))
	}

	return true, nil
}

func ComparePassword2(userPassword string, inputPassword string, cryptCode string) (bool, error) {
	inputPwd := []byte(inputPassword + cryptCode)
	newHash, err := bcrypt.GenerateFromPassword(inputPwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return false, errors.New(fmt.Sprintf("password comparison error: %v", err.Error()))
	}

	return string(newHash) == userPassword, nil
}

func GetHashValue(info interface{}, cryptCode string) (string, error) {
	infoStr := []byte(fmt.Sprintf("%v", info) + cryptCode + "\n")
	h := sha256.New()
	_, err := h.Write(infoStr)
	if err != nil {
		log.Println(err)
		return "", errors.New(fmt.Sprintf("hashing / encryption error: %v", err.Error()))
	}

	res := fmt.Sprintf("%x", h.Sum(nil))
	return res, nil
}

// CurrencyNumFormatter format a number Value based on the specified currency and precision
func CurrencyNumFormatter(num *big.Float, currency string, precision int) (string, error) {
	// return '$' + num.toFixed(2).replace(/(\d)(?=(\d{3})+(?!\d))/g, '$1,')
	//ac := accounting.Accounting{Symbol: "$", Precision: 2}
	if currency == "" {
		currency = "$"
	}
	if precision < 0 {
		precision = 0
	}
	ac := accounting.Accounting{Symbol: currency, Precision: precision}
	res := fmt.Sprintf("%v", ac.FormatMoneyBigFloat(num))

	return res, nil
}

// NumFormatter format a number Value based on the specified precision
func NumFormatter(num *big.Float, precision int) (string, error) {
	if precision < 0 {
		precision = 0
	}
	ac := accounting.Accounting{Symbol: "", Precision: precision}
	res := fmt.Sprintf("%v", ac.FormatMoneyBigFloat(num))

	return res, nil
}

// JsonToStruct converts json input, in []byte, to the equivalent struct data type specification
// rec must be a pointer to a type matching the jsonRec
func JsonToStruct(jsonRec []byte, rec interface{}) error {
	if err := json.Unmarshal(jsonRec, &rec); err == nil {
		return nil
	} else {
		return errors.New(fmt.Sprintf("Error converting json-to-record-format: %v", err.Error()))
	}
}

//type ActionParamType map[string]interface{}
//type ActionParamsType []ActionParamType

// StructToMapUnderscoreData transform a struct data type to the map[string]interface{} Value type, camelCase fields.
// map keys are converted to underscore types to match database-table-field/columns specs
func StructToMapUnderscoreData(rec interface{}) (map[string]interface{}, error) {
	// validate recs as struct{} type
	recType := fmt.Sprintf("%v", reflect.TypeOf(rec).Kind())
	switch recType {
	case "struct":
		dataValue := map[string]interface{}{}
		v := reflect.ValueOf(rec)
		typeOfS := v.Type()

		for i := 0; i < v.NumField(); i++ {
			dataValue[CaseFieldToUnderscore(typeOfS.Field(i).Name)] = v.Field(i).Interface()
			// dataValue[govalidator.CamelCaseToUnderscore(typeOfS.Field(i).Name)] = v.Field(i).Interface()
			//fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).ItemName, v.Field(i).Interface())
		}
		return dataValue, nil
	default:
		return nil, errors.New(fmt.Sprintf("rec parameter must be of type struct{}"))
	}
}

// TagField returns the field-tag (e.g. table-column-name) for mcorm/db/other tag, for the specified struct data fieldName.
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
	// check fieldName, as specified
	field, found := t.FieldByName(fieldName)
	if !found {
		// check public field/PascalCase, converts the first-letter to upper-case
		field, found = t.FieldByName(strings.ToUpper(string(fieldName[0]) + fieldName[1:]))
		if !found {
			return "", errors.New(fmt.Sprintf("error retrieving tag-field for fieldName: %v", fieldName))
		}
	}
	//tagValue := field.Tag
	return field.Tag.Get(tag), nil
}

// StructToMap function transforms struct to map. Map keys match the json-fields of the struct data.
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

// StructToMapTag function converts struct to map (with key of tag/underscore name), for crud-db-table-record
func StructToMapTag(rec interface{}, tag string) (map[string]interface{}, error) {
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
		caseUnderscoreMapData[CaseFieldToUnderscore(key)] = val
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
		uMapData[CaseFieldToUnderscore(key)] = val
	}
	return uMapData, nil
}

// MapUnderscoreToMapCamelCase converts map underscore-fields to camelCase-fields
func MapUnderscoreToMapCamelCase(rec interface{}, sep string) (map[string]interface{}, error) {
	// validate acceptable separators [" ", "_", "__", ".", "|"]
	sepArr := PermittedSeparators
	if !ArrayContains(sepArr, sep) {
		return nil, errors.New(fmt.Sprintf("missing or unacceptable separator: %v | Acceptable-separators: %v", sep, strings.Join(sepArr, ", ")))
	}
	// validate recs as map type
	recMap, ok := rec.(map[string]interface{})
	if !ok || recMap == nil {
		return nil, errors.New(fmt.Sprintf("rec parameter must be of type map[string]interface{}"))
	}

	uMapData := map[string]interface{}{}
	// compose uMapData
	for key, val := range recMap {
		if strings.Contains(key, sep) {
			if camelCaseField, err := SeparatorFieldToCamelCase(key, sep); err != nil {
				return nil, err
			} else {
				uMapData[camelCaseField] = val
			}
		} else {
			uMapData[key] = val
		}

	}
	return uMapData, nil
}

// ArrayMapToMapUnderscore converts slice of map to map with underscore keys/fields.
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

// StructToFieldValues converts struct to record fields(underscore) and associated values - columns and values.
func StructToFieldValues(rec interface{}) (tableFields []string, fieldValues []interface{}, err error) {
	// validate recs as struct{} type
	recType := fmt.Sprintf("%v", reflect.TypeOf(rec).Kind())
	switch recType {
	case "struct":
		break
	default:
		return nil, nil, errors.New(fmt.Sprintf("rec parameter must be of type struct{}"))
	}
	//var tableFields []string
	//var fieldValues []interface{}
	mapDataValue, err := StructToMap(rec)
	if err != nil {
		return nil, nil, errors.New("error computing struct to map")
	}
	// compose table fields/column(underscore) and values
	for key, val := range mapDataValue {
		tableFields = append(tableFields, CaseFieldToUnderscore(key))
		fieldValues = append(fieldValues, val)
	}
	return tableFields, fieldValues, nil
}

// ArrayMapToStruct converts []map/actParams to []struct/model-type
func ArrayMapToStruct(actParams []map[string]interface{}, recs interface{}) (interface{}, error) {
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

// GetParamsMessage compose the message-object into mcresponse.ResponseMessage.
func GetParamsMessage(msgObject MessageObject) mcresponse.ResponseMessage {
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

// ConvertJsonStringToMapValue converts the db-json-string-Value to the map-type.
// Returns map Value on success, or on error, returns map nil-Value and error.
func ConvertJsonStringToMapValue(jsonStr string) (map[string]interface{}, error) {
	mapVal := map[string]interface{}{}
	jErr := json.Unmarshal([]byte(jsonStr), &mapVal)
	if jErr != nil {
		return nil, jErr
	}
	return mapVal, nil
}

// ConvertJsonStringToTypeValue converts the db-json-string-Value to the base-type.
// On success, store the transformed/converted Value to the base-type address (typePointer).
func ConvertJsonStringToTypeValue(jsonStr string, typePointer interface{}) error {
	jErr := json.Unmarshal([]byte(jsonStr), typePointer)
	if jErr != nil {
		return jErr
	}
	return nil
}

// ConvertJsonBase64StringToTypeValue converts the db-json-string-Value to the base-type.
// On success, store and returns the transformed/converted Value to the base-type address (typePointer).
func ConvertJsonBase64StringToTypeValue(base64Str interface{}, typePointer interface{}) (interface{}, error) {
	// assert the base64String Value as of string-type
	strVal, ok := base64Str.(string)
	if !ok {
		return nil, errors.New(fmt.Sprintf("unable to convert base64-string [%v] to string", base64Str))
	}
	// decode the base64StringValue
	decoded, err := base64.StdEncoding.DecodeString(strVal)
	if err != nil {
		return nil, err
	}
	// transform/un-marshal the decoded Value to the base-type
	jErr := json.Unmarshal(decoded, typePointer)
	if jErr != nil {
		return nil, jErr
	}
	return typePointer, nil
}

// ConvertJsonBase64StringToMap converts the db-json-string-Value to the map-type
// On success, returns the resulting map-Value.
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

// ConvertByteSliceToBase64Str converts slice of byte to base64-string type.
func ConvertByteSliceToBase64Str(fileContent []byte) string {
	return base64.StdEncoding.EncodeToString(fileContent)
}

// ConvertStringToBase64Str converts string Value to slice of byte.
func ConvertStringToBase64Str(fileContent string) string {
	return base64.StdEncoding.EncodeToString([]byte(fileContent))
}

// ExcludeEmptyIdFromMapRecord excludes id field with zero Value("").
// Return the map record with id fields removed.
func ExcludeEmptyIdFromMapRecord(rec map[string]interface{}) map[string]interface{} {
	mapVal := map[string]interface{}{}
	for key, val := range rec {
		if key == "id" && (val == nil || val == "") {
			continue
		}
		mapVal[key] = val
	}
	return mapVal
}

// ExcludeFieldFromMapRecord excludes the specified field name from the map record.
// Return the map record with the specified field removed.
func ExcludeFieldFromMapRecord(rec map[string]interface{}, field string) map[string]interface{} {
	mapVal := map[string]interface{}{}
	for key, val := range rec {
		if key == field {
			continue
		}
		mapVal[key] = val
	}
	return mapVal
}

// ExcludeEmptyIdFields excludes fields with name id or ending with Id/ID/iD, with zero Value("").
// Return the map record with id fields removed.
func ExcludeEmptyIdFields(recs []map[string]interface{}) []map[string]interface{} {
	var mapValues []map[string]interface{}
	for _, rec := range recs {
		mapVal := map[string]interface{}{}
		for key, val := range rec {
			if (key == "id" || strings.HasSuffix(strings.ToLower(key), "id")) && (val == nil || val == "") {
				continue
			}
			mapVal[key] = val
		}
		mapValues = append(mapValues, mapVal)
	}
	return mapValues
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

// CryptoRandomNumbers generates cryptographic random numbers using cryptoRand.Prime and return string of digits.
func CryptoRandomNumbers(bits int) (string, error) {
	if bits < 2 {
		bits = 2
	}
	p, err := cryptoRand.Prime(cryptoRand.Reader, bits)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", p), nil
}
