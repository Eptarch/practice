package matrix

import (
    "fmt"
    "strconv"
    "strings"
)

type Matrix [][]int

func New(s string) (*Matrix, error) {
    var matrix = &Matrix{}
    lines := strings.Split(s, "\n")
    previousRowLength := 0
    for iteration, line := range lines {
        var row = []int{}
        for _, a := range strings.Fields(line) {
            n, err := strconv.Atoi(string(a))
            if err != nil {
                return nil, fmt.Errorf("conversion went wrong: %s", err)
            }
            row = append(row, n)
        }
        if iteration != 0 && len(row) != previousRowLength {
            return nil, fmt.Errorf("uneven rows")
        }
        previousRowLength = len(row)
        *matrix = append(*matrix, row)
    }
    return matrix, nil
}

func (m *Matrix) Cols() [][]int {
    var mx = make([][]int, len((*m)[0]))
    for i := 0; i < len((*m)[0]); i++{
        for _, row := range *m {
            mx[i] = append(mx[i], row[i])
        }
    }
    return mx
}

func (m *Matrix) Rows() [][]int {
    var mx = make([][]int, len(*m))
    for row, items := range *m {
        for _, item := range items {
            mx[row] = append(mx[row], item)
        }
    }
    return mx
}

func (m *Matrix) Set(row, col, val int) bool {
    if row < 0 || col < 0 || row > len(*m) - 1 || col > len((*m)[0]) - 1 {
        return false
    }
    (*m)[row][col] = val
    return true
}
