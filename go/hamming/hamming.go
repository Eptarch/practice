package hamming

import "fmt"

type SequenceLengthMismatch struct {
    Diff int
}

func (e *SequenceLengthMismatch) Error() string {
    return fmt.Sprintf("Sequences have %d items length mismatch", e.Diff)
}

func Distance(a, b string) (int, error) {
    diff := len(a) - len(b)
    if diff != 0 {
        return -1, &SequenceLengthMismatch{Diff: diff}
    }
    result := 0
    for i := range a {
        if a[i] != b[i] {
            result++
        }
    }
    return result, nil
}
