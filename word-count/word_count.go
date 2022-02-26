package wordcount

import (
	"regexp"
    "strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
    r, _ := regexp.Compile(`\b[\'\w]+\b`)
    words := r.FindAllString(phrase, -1)
    result := Frequency{}
    for _, word := range words {
        result[strings.ToLower(word)]++
    }
    return result
}
