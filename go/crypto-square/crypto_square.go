package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	var normalized = normalize(pt)
	var rows, columns = dimensions(normalized)
	result := make([]string, columns)
	for i := 0; i < columns; i++ {
		row := make([]uint8, rows)
		for j := 0; j < rows; j++ {
			var char uint8 = ' '
			if len(normalized) > columns * j + i {
				char = normalized[columns * j + i]
			}
			row[j] = char
		}
		result[i] = string(row)
	}
	return strings.Join(result, " ")
}

func dimensions(text string) (rows int, columns int) {
	sqrt := math.Sqrt(float64(len(text)))
	return int(math.Round(sqrt)), int(math.Ceil(sqrt))
}

func normalize(text string) string {
    var normalized strings.Builder
    for _, r := range text {
        if unicode.IsDigit(r) || unicode.IsLetter(r) {
            normalized.WriteRune(unicode.ToLower(r))
        }
    }
	return normalized.String()
}

