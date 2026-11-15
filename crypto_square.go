package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	// Normalize: keep only alphanumeric, lowercase
	var normalized strings.Builder
	for _, r := range pt {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			normalized.WriteRune(unicode.ToLower(r))
		}
	}

	norm := normalized.String()
	if len(norm) == 0 {
		return ""
	}

	n := len(norm)

	// Find r (rows) and c (columns) such that:
	// - r * c >= n
	// - c >= r
	// - c - r <= 1
	r := int(math.Ceil(math.Sqrt(float64(n))))
	for r*r < n {
		r++
	}

	c := (n + r - 1) / r // ceiling division
	for c < r {
		r--
		c = (n + r - 1) / r
	}
	for c-r > 1 {
		r--
		c = (n + r - 1) / r
	}

	// Pad the normalized string to fill the rectangle
	for len(norm) < r*c {
		norm += " "
	}

	// Read down columns left to right and build output
	var result strings.Builder
	for j := 0; j < c; j++ {
		if j > 0 {
			result.WriteRune(' ')
		}
		for i := 0; i < r; i++ {
			result.WriteByte(norm[i*c+j])
		}
	}

	return result.String()
}
