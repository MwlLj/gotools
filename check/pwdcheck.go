package check

import (
	"unicode"
)

func ComplicatedPwdCheck(pwd []rune, min int, max int) bool {
	result := 0
	for _, r := range pwd {
		isDigit := unicode.IsDigit(r)
		if isDigit && (result&1 == 0) {
			result |= 1
		}
		isLower := unicode.IsLower(r)
		if isLower && (result&2 == 0) {
			result |= 2
		}
		isUpper := unicode.IsUpper(r)
		if isUpper && (result&4 == 0) {
			result |= 4
		}
		if (!isDigit && !isLower && !isUpper) && (result&8 == 0) {
			result |= 8
		}
	}
	length := len(pwd)
	if (result == 7 || result == 11 || result == 13 || result == 14 || result == 15) && (length >= min && length <= max) {
		return true
	} else {
		return false
	}
}
