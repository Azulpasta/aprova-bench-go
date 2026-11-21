package matrix

import "testing"

func TestDobro(t *testing.T) {
	casos := map[int]int{0: 0, 1: 2, -3: -6, 21: 42}
	for entrada, esperado := range casos {
		if obtido := Dobro(entrada); obtido != esperado {
			t.Errorf("Dobro(%d) = %d, esperava %d", entrada, obtido, esperado)
		}
	}
}
