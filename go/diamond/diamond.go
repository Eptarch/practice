package diamond

import (
    "bytes"
    "errors"
)

func Gen(char byte) (string, error) {
    if char < 'A' || char > 'Z' {
        return "", errors.New("out of bounds")
    }
    distance := int(char - 'A')
    rows := make([][]byte, 2 * distance + 1)
    for i, j := distance, 0; i >= 0; i, j = i - 1, j + 1 {
        line := bytes.Repeat([]byte{' '}, 2 * distance + 2)
        line[2 * distance + 1] = '\n'
        c := 'A' + byte(j)
        line[i], line[2 * distance - i] = c, c
        rows[j], rows[2 * distance - j] = line, line
    }
	return string(bytes.Join(rows, nil)), nil
}

