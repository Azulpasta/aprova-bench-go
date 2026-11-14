package auth

import "unicode"

func SenhaForte(senha string) bool {
	if len(senha) < 12 {
		return false
	}
	for _, c := range senha {
		if unicode.IsDigit(c) {
			return true
		}
	}
	return false
}
