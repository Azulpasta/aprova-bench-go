package transpose

import "strings"

func Transpose(input []string) []string {
	if len(input) == 0 {
		return []string{}
	}

	// Convert strings to runes and find max length
	runes := make([][]rune, len(input))
	maxLen := 0
	for i, line := range input {
		runes[i] = []rune(line)
		if len(runes[i]) > maxLen {
			maxLen = len(runes[i])
		}
	}

	// Transpose
	result := make([]string, maxLen)
	for col := 0; col < maxLen; col++ {
		var sb strings.Builder
		for _, lineRunes := range runes {
			if col < len(lineRunes) {
				sb.WriteRune(lineRunes[col])
			} else {
				sb.WriteRune(' ')
			}
		}
		result[col] = strings.TrimRight(sb.String(), " ")
	}

	return result
}
