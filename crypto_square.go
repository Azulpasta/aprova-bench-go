package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	var sb strings.Builder
	for _, r := range pt {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sb.WriteRune(unicode.ToLower(r))
		}
	}
	normalized := sb.String()

	length := len(normalized)
	if length == 0 {
		return ""
	}

	c := int(math.Ceil(math.Sqrt(float64(length))))
	r := c
	if c*(c-1) >= length {
		r = c - 1
	}

	rows := make([]string, r)
	for i := 0; i < r; i++ {
		start := i * c
		end := start + c
		if end > length {
			end = length
		}
		row := normalized[start:end]
		if len(row) < c {
			row += strings.Repeat(" ", c-len(row))
		}
		rows[i] = row
	}

	chunks := make([]string, c)
	for col := 0; col < c; col++ {
		var chunk strings.Builder
		for row := 0; row < r; row++ {
			chunk.WriteByte(rows[row][col])
		}
		chunks[col] = chunk.String()
	}

	return strings.Join(chunks, " ")
}
