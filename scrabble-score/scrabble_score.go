package scrabble

import (
    "strings"
)
// LetterScore decides each letter worth.
func LetterScore(letter rune) int {
    switch letter {
        case 'q', 'z': return 10
        case 'j', 'x': return 8
        case 'k': return 5
        case 'f', 'h', 'v', 'w', 'y': return 4
        case 'b', 'c', 'm', 'p': return 3
        case 'd', 'g': return 2
        default: return 1
    }
}
// Score converts the word to lowercase and calls LetterScore on each of its letters.
func Score(word string) int {
    r := 0
    for _, l := range strings.ToLower(word) {
        r += LetterScore(l)
    }
    return r
}
