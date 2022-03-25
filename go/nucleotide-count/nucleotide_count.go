package dna

import "fmt"

type Histogram map[rune]int
type DNA string

func (d DNA) Counts() (Histogram, error) {
    var histogram = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
    for _, nucleotide := range d {
        if _, ok := histogram[nucleotide]; !ok {
            return nil, fmt.Errorf("invalid nucleotide: %v", nucleotide)
        }
        histogram[nucleotide]++
    }
	return histogram, nil
}
