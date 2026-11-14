package transpose

import "strings"

func Transpose(input []string) []string {
	maxLen := 0
	for _, row := range input {
		if len(row) > maxLen {
			maxLen = len(row)
		}
	}

	result := make([]string, maxLen)
	for c := 0; c < maxLen; c++ {
		var sb strings.Builder
		for r, row := range input {
			if c < len(row) {
				sb.WriteByte(row[c])
				continue
			}
			laterHasColumn := false
			for _, later := range input[r+1:] {
				if len(later) > c {
					laterHasColumn = true
					break
				}
			}
			if !laterHasColumn {
				break
			}
			sb.WriteByte(' ')
		}
		result[c] = sb.String()
	}
	return result
}
