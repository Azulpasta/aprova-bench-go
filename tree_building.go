package tree

import "fmt"

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	n := len(records)
	if n == 0 {
		return nil, nil
	}

	nodes := make([]*Node, n)
	parentOf := make([]int, n)
	seen := make([]bool, n)

	for _, r := range records {
		if r.ID < 0 || r.ID >= n {
			return nil, fmt.Errorf("invalid record id %d: out of range [0, %d)", r.ID, n)
		}
		if seen[r.ID] {
			return nil, fmt.Errorf("duplicate record id %d", r.ID)
		}
		seen[r.ID] = true
		parentOf[r.ID] = r.Parent
		nodes[r.ID] = &Node{ID: r.ID}
	}

	for id := 0; id < n; id++ {
		if !seen[id] {
			return nil, fmt.Errorf("missing record id %d", id)
		}
	}

	if parentOf[0] != 0 {
		return nil, fmt.Errorf("root record must have itself as parent, got parent %d", parentOf[0])
	}

	for id := 1; id < n; id++ {
		p := parentOf[id]
		if p < 0 || p >= id {
			return nil, fmt.Errorf("record %d has invalid parent %d", id, p)
		}
		nodes[p].Children = append(nodes[p].Children, nodes[id])
	}

	return nodes[0], nil
}
