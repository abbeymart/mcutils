// @Author: abbeymart | Abi Akindele | @Created: 2020-12-02 | @Updated: 2020-12-02
// @Company: mConnect.biz | @License: MIT
// @Description: mcutils types and constants

package mcutils

type LocaleContent map[string]interface{}

//type Locale map[string]LocaleContent

type LocaleOptions struct {
	LocaleType string
	Language   string
}

type MessageObject map[string]string

const (
	DefaultLanguage = "en-US"
)

type ValueType interface {
	string | int | int64 | float64 | float32 | bool
}

type ValueTypeSlice interface {
	[]string | int | []int64 | []float64 | float32 | []bool | []map[string]interface{}
}

type Number interface {
	int | int64 | float64 | float32
}

type TestFuncType[T ValueType] func(val T) bool

type LocaleFunc[T ValueType] func() T

type LocaleValueType[T ValueType] interface {
	ValueType | LocaleFunc[T]
}

type ObjectType map[string]interface{}

type Locale[T ValueType] map[string]LocaleValueType[T]

type LocaleFilesType[T ValueType] map[string]Locale[T] // key => language ("en-US", "en-CA", "yoruba", "fr-CA", "fr-FR" etc.

// types

var PermittedSeparators = []string{" ", "_", "__", ".", "|"}

// Collection function types

type IntPredicate func(val int) bool
type FloatPredicate func(val float64) bool
type StringPredicate func(val string) bool
type NumberPredicate[T Number] func(val T) bool
type Predicate[T ValueType] func(val T) bool
type BinaryPredicate[T ValueType, U ValueType] func(val1 T, val2 U) bool
type UnaryOperator[T ValueType] func(val1 T) T
type BinaryOperator[T ValueType] func(val1 T, val2 T) T

type Function[T ValueType, R ValueType] func(val T) R
type BiFunction[T ValueType, U ValueType, R ValueType] func(val1 T, val2 U) R
type Consumer[T ValueType] func(val T)
type BiConsumer[T ValueType, U ValueType] func(val1 T, val2 U)
type Supplier[R ValueType] func() R
type Comparator[T ValueType] func(val1 T, val2 T) int

type StringType struct {
	Value string `json:"Value"`
}
type FloatType struct {
	Value float64 `json:"value"`
}
type BoolType struct {
	Value bool `json:"value"`
}
type IntType struct {
	Value int `json:"value"`
}
type GenericType[T ValueType] struct {
	Value T `json:"value"`
}
type StringSliceType struct {
	Value []string `json:"value"`
}
type FloatSliceType struct {
	Value []float64 `json:"value"`
}

//type FloatSliceType []float64

type BoolSliceType struct {
	Value []bool `json:"Value"`
}

//type BoolSliceType []bool

type IntSliceType struct {
	Value []int `json:"Value"`
}
type GenericSliceType[T ValueType] struct {
	Value []T `json:"Value"`
}

type QuartilesType struct {
	Min   float64 `json:"min"` // Lowest value
	Q1    float64 `json:"q1"`
	Q2    float64 `json:"q2"` // Median
	Q3    float64 `json:"q3"`
	Q4    float64 `json:"q4"` // Highest value, Max
	IQR   float64 `json:"IQR"`
	Max   float64 `json:"max"`   // Highest value, Q4
	Range float64 `json:"range"` // Q3 - Q1
}
