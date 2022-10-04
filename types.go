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

type Number interface {
	int64 | float64
}

type TestFuncType[T ValueType] func(val T) bool

// types

type McObjectString struct {
	value string
}
type McObjectFloat struct {
	value float64
}
type McObjectBool struct {
	value bool
}
type McObjectInt struct {
	value int
}
type McObjectInterface[T ValueType] struct {
	value T
}
type McObjectStringSlice struct {
	value []string
}
type McObjectFloatSlice struct {
	value []float64
}
type McObjectBoolSlice struct {
	value []bool
}
type McObjectIntSlice struct {
	value []int
}
type McObjectInterfaceSlice[T ValueType] struct {
	value []T
}
