package anagram

import (
    "strings"
    "reflect"
)

func buildMap(s string) map[rune]int {
    var stringmap = map[rune]int{}
    for _, r := range strings.ToLower(s) {
        stringmap[r] += 1
    }
    return stringmap
}

func Detect(subject string, candidates []string) []string {
    var result = []string{}
    subjectmap := buildMap(subject)
    for index, candidate := range candidates {
        if reflect.DeepEqual(subjectmap, buildMap(candidate)) && strings.ToLower(subject) != strings.ToLower(candidate) {
            result = append(result, candidates[index])
        }
    }
    return result
}
