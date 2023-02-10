package data

var EmptyObjectValue = map[string]interface{}{}
var NonEmptyObjectValue = map[string]interface{}{"name": "Abi", "location": "Toronto"}

const ShortStringParam = "This is a great title for testing"
const Short21StringResult = "This is a great title..."
const Short20StringResult = "This is a great titl..."

var ParamObjectMsgData = map[string]string{"name": "required", "location": "optional"}

const ParamObjectMsgResult = `name: required | location: optional`

const StrTrueToBool = "true"
const StrYesToBool = "yes"
const StrYToBool = "Y"
const StrFalseToBool = "false"
const StrNoToBool = "no"
const StrNToBool = "N"
const Str1ToBool = 1
const Str0ToBool = 0
const StrEmptyToBool = ""

const Firstname = "Abi"
const Middlename = "John"
const Lastname = "Owo"
const FullnameTwo = "Abi Owo"
const FullnameThree = "Abi John Owo"

const CamelCaseValue = "countryCode"
const UnderscoreValue = "country_code"
const PascalCaseValue = "CountryCode"
const DotSepParam = "country.code"
const PipeSepParam = "country|code"
const SpaceSepParam = "country code"
