package strand

import "strings"

var replacer = strings.NewReplacer("G", "C", "C", "G", "T", "A", "A", "U")

func ToRNA(dna string) string {
	return replacer.Replace(dna)
}

