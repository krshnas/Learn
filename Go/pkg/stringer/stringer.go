package stringer

import (
	"unicode"
)

// Reverse reverses the input string and returns the reversed result.
func Reverse(input string) string {
	var result string
	for i := len(input) - 1; i >= 0; i-- {
		result += string(input[i])
	}
	return result
}

// Inspect inspects the input string and counts either characters or digits based on the `digits` flag.
// It returns the count and the kind ("char" or "digit").
func Inspect(input string, digits bool) (count int, kind string) {
	if digits {
		count = countDigits(input)
		kind = "digit"
	} else {
		count = len(input)
		kind = "char"
	}
	return
}

// countDigits counts the number of digits in the input string.
func countDigits(input string) (count int) {
	for _, c := range input {
		if unicode.IsDigit(c) {
			count++
		}
	}
	return count
}
