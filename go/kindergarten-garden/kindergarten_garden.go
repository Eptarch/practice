package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

type Garden map[string][]string

var plants = map[rune]string{
    'V': "violets",
    'R': "radishes",
    'G': "grass",
    'C': "clover",
}
// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	if children == nil {
        return nil, errors.New("no children in kindergarten :(")
    }
    ordered := make([]string, len(children))
    copy(ordered, children)
    sort.Sort(sort.StringSlice(ordered))
    garden := Garden{}
    rows := strings.Split(diagram, "\n")
    if len(rows) != 3 {
        return nil, errors.New("invalid diagram format")
    }
    row0, row1 := []rune(rows[1]), []rune(rows[2])
    if len(row0) != 2 * len(ordered) || len(row1) != 2 * len(ordered) {
        return nil, errors.New("invalid row length") 
    }
    for i, child := range ordered {
        plant0, ok0 := plants[row0[2 * i]]
        plant1, ok1 := plants[row0[2 * i + 1]]
        plant2, ok2 := plants[row1[2 * i]]
        plant3, ok3 := plants[row1[2 * i + 1]]
        if !(ok0 && ok1 && ok2 && ok3) {
            return nil, errors.New("invalid items in pot")
        }
        plants := []string{plant0, plant1, plant2, plant3}
        garden[child] = plants
    }
    if len(garden) != len(children) {
        return nil, errors.New("duplicate name")
    }
    return &garden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
    result, ok := (*g)[child]
    return result, ok
}
