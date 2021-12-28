package isogram

import "strings"

func IsIsogram(word string) bool {
    var m = map[rune]int{}
    for _, r := range strings.ToLower(word) {
        m[r]++
        if !(r == '-' || r == ' ') && m[r] > 1 {
            return false
        }
    }
    return true
}
