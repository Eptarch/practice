package binarysearchtree

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// NewBst creates and returns a new SearchTreeData.
func NewBst(i int) SearchTreeData {
    return SearchTreeData{data: i}
}

// Insert inserts an int into the SearchTreeData.
// Inserts happen based on the rules of a BinarySearchTree
func (std *SearchTreeData) Insert(i int) {
    switch {
        case i <= std.data && std.left == nil:
            std.left = &SearchTreeData{data: i}
        case i > std.data && std.right == nil: 
            std.right = &SearchTreeData{data: i}
        case i <=std.data:
            std.left.Insert(i)
        case i > std.data:
            std.right.Insert(i)
    }
}

// MapString returns the ordered contents of SearchTreeData as a []string.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []string ["1", "3", "5", "7"].
func (std *SearchTreeData) MapString(f func(int) string) (result []string) {
    ints := std.MapInt(func(i int) int { return i })
    for _, n := range ints {
        result = append(result, f(n))
    }
	return 
}

// MapInt returns the ordered contents of SearchTreeData as an []int.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (std *SearchTreeData) MapInt(f func(int) int) (result []int) {
    if std.left != nil {
        result = std.left.MapInt(f)
    }
    result = append(result, f(std.data))
    if std.right != nil {
        result = append(result, std.right.MapInt(f)...)
    }
	return
}
