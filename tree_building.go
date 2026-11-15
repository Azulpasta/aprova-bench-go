package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	seen := make(map[int]bool)
	for _, record := range records {
		if record.ID < 0 || record.ID >= len(records) {
			return nil, errors.New("invalid record ID")
		}
		if seen[record.ID] {
			return nil, errors.New("duplicate record ID")
		}
		seen[record.ID] = true
	}

	nodes := make([]*Node, len(records))
	for i := range nodes {
		nodes[i] = &Node{ID: i}
	}

	var root *Node
	for _, record := range records {
		if record.ID == record.Parent {
			if record.ID != 0 {
				return nil, errors.New("root node must have ID 0")
			}
			if root != nil {
				return nil, errors.New("duplicate root node")
			}
			root = nodes[0]
		} else {
			if record.Parent >= record.ID {
				return nil, errors.New("parent ID must be less than node ID")
			}
			nodes[record.Parent].Children = append(nodes[record.Parent].Children, nodes[record.ID])
		}
	}

	if root == nil {
		return nil, errors.New("no root node found")
	}

	sortChildren(root)
	return root, nil
}

func sortChildren(node *Node) {
	sort.Slice(node.Children, func(i, j int) bool {
		return node.Children[i].ID < node.Children[j].ID
	})
	for _, child := range node.Children {
		sortChildren(child)
	}
}
