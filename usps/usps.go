// Package tools provides various functions for working with USPS tracking numbers.
package usps

// Format accepts a tracking number.
// It returns the tracking number with a space inserted between every 4 characters.
func Format(in string) (s string) {
	r := []rune(in)
	for i := 0; i < len(r); i++ {
		if i != 0 && i%4 == 0 {
			s += " "
		}
		s += string(r[i])
	}
	return s
}
