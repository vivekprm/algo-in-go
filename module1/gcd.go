package module01

// GCD stands for greatest common divisor
// Given two numbers, GCD calculates the largest number you could divide both numbers by without getting a remainder.

// Examples:
// A = 10, B = 5, GCD = 5 (10%5 == 0, 5%5 == 0)
// A = 25, B = 5, GCD = 5
// A = 30, B = 15, GCD = 15
// A = 30, B = 9, GCD = 3
// A = 100, B = 9, GCD = 1
func GCD(a, b int) int {
	return GCD(b/a, a)
}
