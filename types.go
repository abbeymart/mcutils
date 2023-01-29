// @Author: abbeymart | Abi Akindele | @Created: 2020-12-02 | @Updated: 2020-12-02
// @Company: mConnect.biz | @License: MIT
// @Description: mcutils types and constants

package mcutils

type LocaleContent map[string]interface{}
type Locale map[string]LocaleContent
type LocaleOptions struct {
	LocaleType string
	Language   string
}

type MessageObject map[string]string

const (
	DefaultLanguage = "en-US"
)

type ValueType interface {
	string | int64 | float64 | bool
}

type ValueTypeSlice interface {
	[]string | []int64 | []float64 | []bool | []map[string]interface{}
}

type Number interface {
	int64 | float64
}

type TestFuncType[T ValueType] func(val T) bool

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
	value string
}
type FloatType struct {
	value float64
}
type BoolType struct {
	value bool
}
type IntType struct {
	value int
}
type GenericType[T ValueType] struct {
	value T
}
type StringSliceType struct {
	value []string
}
type FloatSliceType struct {
	value []float64
}
type BoolSliceType struct {
	value []bool
}
type IntSliceType struct {
	value []int
}
type GenericSliceType[T ValueType] struct {
	value []T
}
