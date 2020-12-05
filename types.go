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
