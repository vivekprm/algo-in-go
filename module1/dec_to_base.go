package module1

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// this is limited to base 2-16 for simplicity.
//
// E.g.
//
// DecToBase(14, 16) => "E"
// DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	const charset = "0123456789ABCDEF"
	var res string
	for dec > 0 {
		rem := dec % base
		res = string(charset[rem]) + res
		dec = dec/base
	}
	return res
}