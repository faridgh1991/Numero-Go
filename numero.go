// A micro library for converting non-english UTF8 digits. (like ۱=1, ۲=2) .
// Almost all numbers defined in Unicode is supported in Numero.
package numero

import (
	"strconv"
	"strings"
)

// zero character in diffrent languages
var zeroStarts = [...]rune{
	'0', '٠', '۰', '߀', '०', '০', '੦', '૦', '୦',
	'௦', '౦', '೦', '൦', '෦', '๐', '໐', '༠', '၀',
	'႐', '០', '᠐', '᥆', '᧐', '᪀', '᪐', '᭐', '᮰',
	'᱀', '᱐', '꘠', '꣐', '꤀', '꧐', '꧰', '꩐', '꯰',
	'０', '𐒠', '𑁦', '𑃰', '𑄶', '𑇐', '𑋰', '𑑐', '𑓐',
	'𑙐', '𑛀', '𑜰', '𑣠', '𑱐', '𑵐', '𖩠', '𖭐', '𝟎',
	'𝟘', '𝟢', '𝟬', '𝟶', '𞥐'}

// english zerro character code
const zeroCode = 48

// check character is digit or not and if true return integer value of character
func IsDigit(char rune) (bool, int) {
	for _, zero := range zeroStarts {
		if char >= zero && char <= zero+9 {
			return true, int(char - zero)
		}
	}
	return false, -1
}

// Checking if a string is all numbers
func DigitOnly(str string) bool {
	for _, c := range str {
		if ok, _ := IsDigit(c); !ok {
			return false
		}
	}
	return true
}

// Normalize all numbers in input string
func Normalize(numberStr string) string {
	normalized := ""
	for _, char := range numberStr {
		if ok, index := IsDigit(char); ok {
			normalized += string(zeroCode + index)
		} else {
			normalized += string(char)
		}
	}
	return normalized
}

// Convert numbers to Integer or Float based on input string
func NormalizeAsNumber(numberStr string) (interface{}, error) {
	if strings.Contains(numberStr, ".") {
		return strconv.ParseFloat(Normalize(numberStr), 64)
	} else {
		return strconv.Atoi(Normalize(numberStr))
	}
}

//Strip all non numeric chars from a string
func RemoveNonDigits(str string, exceptions ...string) string {
	normalized := ""
	for _, char := range str {
		if ok, index := IsDigit(char); ok {
			normalized += string(zeroCode + index)
		} else if len(exceptions) > 0 && strings.Contains(exceptions[0], string(char)) {
			normalized += string(char)
		}
	}
	return normalized
}
