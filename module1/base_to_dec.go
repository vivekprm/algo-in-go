package module1

import (
	"strings"
)

const charset = "0123456789ABCDEF"
// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(val string, base int) int {
	res := 0
	multiplier := 1
	for i := len(val) - 1; i >= 0; i-- {
		res += GetIdx(string(val[i])) * multiplier
		multiplier *= base
	}
	return res
}

func GetIdx(char string) int {
	idx := strings.Index(charset, char)
	return idx
}