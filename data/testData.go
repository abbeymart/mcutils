package data

import (
	"github.com/abbeymart/mcutils"
	"math"
	"strings"
)

// collection

var ArrayOfNumber = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var ArrayOfString = []string{"abc", "ab2", "abc3", "ab4", "abc5", "ab6", "abc7", "ab8", "abc9", "ab10"}

// const arrayOfSymbol: ArrayOfSymbol = [Symbol("abc"), Symbol("ab2"), Symbol("ab3"), Symbol("ab4"), Symbol("ab5"),
//     Symbol("ab6"), Symbol("ab7"), Symbol("ab8"), Symbol("ab9"), Symbol("ab10"),]

func FilterEvenNumFunc[T mcutils.Number](val T) bool {
	return math.Abs(math.Remainder(float64(val), 2)) == 0.00
}

var FilterEvenNumFuncResult = []int{2, 4, 6, 8, 10}

func FilterOddNumFunc[T mcutils.Number](val T) bool {
	return math.Abs(math.Remainder(float64(val), 2)) != 0.00
}

var FilterOddNumFuncResult = []int{1, 3, 5, 7, 9}

func FilterStringIncludeABC(val string) bool {
	return strings.Contains(val, "abc")
}

var FilterStringIncludeABCResult = []string{"abc", "abc3", "abc5", "abc7", "abc9"}

func MapDoubleNumFunc[T mcutils.Number](val T) T {
	return val * 2
}

var MapDoubleNumFuncResult = []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

var Take7NumResult = []int{1, 2, 3, 4, 5, 6, 7}

const take7CountResult = 7

var Take7StringResult = []string{"abc", "ab2", "abc3", "ab4", "abc5", "ab6", "abc7"}

// getLocale

var LocaleLabelOptions = mcutils.LocaleOptions{
	LocaleType: "mcLabels",
	Language:   "en-CA",
}

var LocaleConstantOptions = mcutils.LocaleOptions{
	LocaleType: "mcConstants",
	Language:   "en-CA",
}

// code      : "Code",
//name      : "Name",
//desc      : "Description",
//postalCode: "Postal Code",

var LabelObject = mcutils.Locale{
	"code":       "Code",
	"name":       "Name",
	"desc":       "Description",
	"postalCode": "Postal Code",
}

var ConstantObject = mcutils.Locale{
	"SHORT_DESC":   20,
	"DEFAULT_LANG": "en-US",
}

var LocaleLabelFiles = mcutils.LocaleFilesType{
	"en-US": LabelObject,
	"en-CA": LabelObject,
}

var LocaleConstantFiles = mcutils.LocaleFilesType{
	"en-US": ConstantObject,
	"en-CA": ConstantObject,
}

// stats

const MeanResult = 5.5
const MedianResult = 5.5
const MinResult = 1
const MaxResult = 10

var MinMaxResult = []float64{1, 10}

const StdDeviationResult = 3.0276503540974917 // 16 decimal places
const StdDeviationResultEst = 3.02765

// populationStandardDeviation

var ArrayOfNumber2 = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5, 9.5}

const StdDeviationResultEst2 = 2.74

// variance

// interval

// frequency

// frequencyStat

// IQRange
