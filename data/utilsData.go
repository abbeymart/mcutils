package data

var emptyObjectValue = map[string]interface{}{}
var nonEmptyObjectValue = map[string]interface{}{"name": "Abi", "location": "Toronto"}

const shortStringParam = "This is a great title for testing"
const short21StringResult = "This is a great title..."
const short20StringResult = "This is a great titl..."

var paramObjectMsgData = map[string]string{"name": "required", "location": "optional"}

const paramObjectMsgResult = `name: required | location: optional`

// const strTrueToBool = "true"
// const strYesToBool = "yes"
// const strYToBool = "Y"
// const strFalseToBool = "false"
// const strNoToBool = "no"
// const strNToBool = "N"
// const str1ToBool = 1
// const str0ToBool = 0
// const strEmptyToBool = ""

const firstname = "Abi"
const middlename = "John"
const lastname = "Owo"
const fullnameTwo = "Abi Owo"
const fullnameThree = "Abi John Owo"

const camelCaseValue = "countryCode"
const underscoreValue = "country_code"
const pascalCaseValue = "CountryCode"
const dotSepParam = "country.code"
const pipeSepParam = "country|code"
const spaceSepParam = "country code"
