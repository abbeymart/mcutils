package data



// collection
const arrayOfNumber: ArrayOfNumber = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10,]
const arrayOfString: ArrayOfString = ["abc", "ab2", "abc3", "ab4", "abc5", "ab6", "abc7", "ab8", "abc9",
"ab10",]
// const arrayOfSymbol: ArrayOfSymbol = [Symbol("abc"), Symbol("ab2"), Symbol("ab3"), Symbol("ab4"), Symbol("ab5"),
//     Symbol("ab6"), Symbol("ab7"), Symbol("ab8"), Symbol("ab9"), Symbol("ab10"),]

const filterEvenNumFunc = (val: number): boolean => val % 2 === 0
const filterEvenNumFuncResult = [2, 4, 6, 8, 10,]
const filterOddNumFunc = (val: number): boolean => val % 2 !== 0
const filterOddNumFuncResult = [1, 3, 5, 7, 9,]
const filterStringIncludeABC = (val: string): boolean => val.includes("abc")
const filterStringIncludeABCResult = ["abc", "abc3", "abc5", "abc7", "abc9",]

const mapDoubleNumFunc = (val: number) => val * 2
const mapDoubleNumFuncResult = [2, 4, 6, 8, 10, 12, 14, 16, 18, 20,]

const take7NumResult = [1, 2, 3, 4, 5, 6, 7]
const take7CountResult = 7
const take7StringResult = ["abc", "ab2", "abc3", "ab4", "abc5", "ab6", "abc7",]

// getLocale
const localeLabelOptions = LocaleOptions{
type    : "mcLabels",
language: "en-CA",
}
const localeConstantOptions: LocaleOptions = {
type    : "mcConstants",
language: "en-CA",
}
var localeLabelObject = Locale{
code      : "Code",
name      : "Name",
desc      : "Description",
postalCode: "Postal Code",
}

const localeConstantObject: Locale = {
SHORT_DESC  : 20,
DEFAULT_LANG: "en-US",
}

const localeLabelFiles = Locale{
"en-US": localeLabelObject,
"en-CA": localeLabelObject,
}

const localeConstantFiles = LocalFilesType {
"en-US": localeConstantObject,
"en-CA": localeConstantObject,
}

// stats

const meanResult = 5.5
const medianResult = 5.5
const minResult = 1
const maxResult = 10
var minMaxResult = []float64{1, 10}

const stdDeviationResult = 3.0276503540974917   // 16 decimal places
const stdDeviationResultEst = 3.02765

// populationStandardDeviation
var arrayOfNumber2 = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5, 9.5}
const stdDeviationResultEst2 = 2.74

// variance


// interval


// frequency


// frequencyStat


// IQRange
