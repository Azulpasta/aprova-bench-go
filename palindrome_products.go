package palindrome

import (
	"fmt"
	"strconv"
)

// Product represents a palindromic product and the factor pairs,
// within the given range, that produce it.
type Product struct {
	Value          int
	Factorizations [][2]int
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, fmt.Errorf("fmin > fmax: %d > %d", fmin, fmax)
	}

	var pmin, pmax Product
	found := false

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			product := i * j
			if !isPalindrome(product) {
				continue
			}
			if !found || product < pmin.Value {
				pmin = Product{product, [][2]int{{i, j}}}
			} else if product == pmin.Value {
				pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
			}
			if !found || product > pmax.Value {
				pmax = Product{product, [][2]int{{i, j}}}
			} else if product == pmax.Value {
				pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
			}
			found = true
		}
	}

	if !found {
		return Product{}, Product{}, fmt.Errorf("no palindromes in range %d to %d", fmin, fmax)
	}

	return pmin, pmax, nil
}
