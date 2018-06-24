package numero

import (
	"strconv"
	"strings"
)

var zero_starts = [...]rune{
	'0', '٠', '۰', '߀', '०', '০', '੦', '૦', '୦',
	'௦', '౦', '೦', '൦', '෦', '๐', '໐', '༠', '၀',
	'႐', '០', '᠐', '᥆', '᧐', '᪀', '᪐', '᭐', '᮰',
	'᱀', '᱐', '꘠', '꣐', '꤀', '꧐', '꧰', '꩐', '꯰',
	'０', '𐒠', '𑁦', '𑃰', '𑄶', '𑇐', '𑋰', '𑑐', '𑓐',
	'𑙐', '𑛀', '𑜰', '𑣠', '𑱐', '𑵐', '𖩠', '𖭐', '𝟎',
	'𝟘', '𝟢', '𝟬', '𝟶', '𞥐'}

const zero_code = 48

func isDigit(char rune) (bool, int) {
	for _, start := range zero_starts {
		if s := rune(start); char >= s && char <= s+9 {
			return true, int(char - s)
		}
	}
	return false, -1
}

func DigitOnly(str string) bool {
	for _, c := range str {
		if ok, _ := isDigit(c); !ok {
			return false
		}
	}
	return true
}

func Normalize(numberStr string) string {
	normalized := ""
	for _, char := range numberStr {
		if ok, index := isDigit(char); ok {
			normalized += string(zero_code + index)
		} else {
			normalized += string(char)
		}
	}
	return normalized
}

func NormalizeAsNumber(numberStr string) (interface{}, error) {
	if strings.Contains(numberStr, ".") {
		return strconv.ParseFloat(Normalize(numberStr), 64)
	} else {
		return strconv.Atoi(Normalize(numberStr))
	}
}

func RemoveNonDigits(str string, exceptions ...string) string {
	normalized := ""
	for _, char := range str {
		if ok, index := isDigit(char); ok {
			normalized += string(zero_code + index)
		} else if len(exceptions) > 0 && strings.Contains(exceptions[0], string(char)) {
			normalized += string(char)
		}
	}
	return normalized
}
