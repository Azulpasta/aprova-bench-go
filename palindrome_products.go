package palindrome

import (
	"fmt"
	"strconv"
)

type Product struct {
	Value          int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, fmt.Errorf("fmin > fmax")
	}

	palindromes := make(map[int][][2]int)

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			product := i * j
			if isPalindrome(product) {
				palindromes[product] = append(palindromes[product], [2]int{i, j})
			}
		}
	}

	if len(palindromes) == 0 {
		return Product{}, Product{}, fmt.Errorf("no palindromes")
	}

	var minVal, maxVal int
	var minFactors, maxFactors [][2]int
	first := true

	for val, factors := range palindromes {
		if first {
			minVal = val
			maxVal = val
			minFactors = factors
			maxFactors = factors
			first = false
		} else {
			if val < minVal {
				minVal = val
				minFactors = factors
			}
			if val > maxVal {
				maxVal = val
				maxFactors = factors
			}
		}
	}

	return Product{minVal, minFactors}, Product{maxVal, maxFactors}, nil
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
