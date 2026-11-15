package sublist

// Relation type is defined in relations.go file.

func isContiguousSubsequence(sub, list []int) bool {
	if len(sub) == 0 {
		return true
	}
	if len(sub) > len(list) {
		return false
	}

	for i := 0; i <= len(list)-len(sub); i++ {
		match := true
		for j := 0; j < len(sub); j++ {
			if list[i+j] != sub[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func Sublist(l1, l2 []int) Relation {
	if len(l1) == len(l2) {
		equal := true
		for i := 0; i < len(l1); i++ {
			if l1[i] != l2[i] {
				equal = false
				break
			}
		}
		if equal {
			return RelationEqual
		}
		return RelationUnequal
	}

	if isContiguousSubsequence(l1, l2) {
		return RelationSublist
	}

	if isContiguousSubsequence(l2, l1) {
		return RelationSuperlist
	}

	return RelationUnequal
}
