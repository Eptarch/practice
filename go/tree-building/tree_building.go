// Package tree implements a method to build a simple
// tree data structure.
package tree

import (
	"errors"
	"sort"
)

// Record represent a single record
type Record struct {
	ID int
	Parent int
}

// Node holds tree node data
type Node struct {
	ID int
	Children []*Node
}


// Build simple tree from a given array of records.
func Build(records []Record) (*Node, error) {
	node := make(map[int]*Node, len(records))
	// Sort by ascending ID.
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, errors.New("not in sequence or has bad parent")
		}
		node[r.ID] = &Node{ID: r.ID}
		if r.ID != 0 {
			node[r.Parent].Children = append(node[r.Parent].Children, node[r.ID])
		}
	}
	return node[0], nil
}
