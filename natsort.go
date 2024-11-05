package natsort

import (
	"unicode"
	"unicode/utf8"
)

// Compare compares two strings using natural alphanumeric order
func Compare(a, b string) bool {
	ai, bi := 0, 0
	an, bn := len(a), len(b)

	for ai < an && bi < bn {
		ar, aw := utf8.DecodeRuneInString(a[ai:])
		br, bw := utf8.DecodeRuneInString(b[bi:])

		if unicode.IsDigit(ar) && unicode.IsDigit(br) {
			aNum, aLen := extractNumber(a, ai)
			bNum, bLen := extractNumber(b, bi)

			if aNum != bNum {
				return aNum < bNum
			}

			ai += aLen
			bi += bLen
		} else {
			// Convert runes to lower case for case-insensitive comparison
			arLower := unicode.ToLower(ar)
			brLower := unicode.ToLower(br)

			if arLower != brLower {
				return arLower < brLower
			}
			ai += aw
			bi += bw
		}
	}

	return an < bn
}

// extractNumber parses a number starting at index `start` in string `s`
// Returns the parsed number and the length of the numeric substring
func extractNumber(s string, start int) (int, int) {
	end := start
	for end < len(s) && '0' <= s[end] && s[end] <= '9' {
		end++
	}
	num := 0
	for i := start; i < end; i++ {
		num = num*10 + int(s[i]-'0')
	}
	return num, end - start
}
