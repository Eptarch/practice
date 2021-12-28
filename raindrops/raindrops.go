package raindrops

import "strconv"

// Pair is a replacement structure.
type Pair struct {
	d int
	w string
}

var replacements = []Pair{
	{d: 3, w: "Pling"},
	{d: 5, w: "Plang"},
	{d: 7, w: "Plong"},
}

// Convert number to a string considering its divisibility by 3, 5, and 7.
func Convert(number int) string {
	r := ""
	for _, p := range replacements {
		if number%p.d == 0 {
			r += p.w
		}
	}
	if r == "" {
		return strconv.Itoa(number)
	}
	return r
}
