package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	switch {
	case equal(l1, l2):
		return RelationEqual
	case contains(l2, l1):
		return RelationSublist
	case contains(l1, l2):
		return RelationSuperlist
	default:
		return RelationUnequal
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func contains(big, small []int) bool {
	if len(small) > len(big) {
		return false
	}
	for i := 0; i+len(small) <= len(big); i++ {
		if equal(big[i:i+len(small)], small) {
			return true
		}
	}
	return false
}
