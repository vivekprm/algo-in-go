package module1

import "strings"

// Reverse will return the provided word in reverse
// order. E.g.
//
// Reverse("cat") => "tac"
// Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	res := ""

	for i := len(word)-1; i >= 0; i-- {
		// When we access a string index, we get byte back.
		res = res + string(word[i])
	}
	return res
}

func ReverseSB(word string) string {
	var sb strings.Builder

	for i := len(word)-1; i >= 0; i-- {
		sb.WriteByte(word[i])
	}	

	return sb.String()
}

// Rather than dealing with single chars as byte deal them as rune.
func ReverseRune(word string) string {
	res := ""

	for _, r := range word {
		res = string(r) + res
	}
	return res
}
