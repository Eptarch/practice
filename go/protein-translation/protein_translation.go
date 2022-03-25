package protein

import "fmt"

var translation = map[string]string{
    "AUG": "Methionine",
    "UUU": "Phenylalanine",
    "UUC": "Phenylalanine",
    "UUA": "Leucine",
    "UUG": "Leucine",
    "UCU": "Serine",
    "UCC": "Serine",
    "UCA": "Serine",
    "UCG": "Serine",
    "UAU": "Tyrosine",
    "UAC": "Tyrosine",
    "UGU": "Cysteine",
    "UGC": "Cysteine",
    "UGG": "Tryptophan",
    "UAA": "STOP",
    "UAG": "STOP",
    "UGA": "STOP",
}

var ErrStop = fmt.Errorf("encountered stop codon")
var ErrInvalidBase = fmt.Errorf("invalid codon base")

func FromRNA(rna string) ([]string, error) {
    var result = []string{}
    for i := 0; i < len(rna) / 3; i++ {
        protein, err := FromCodon(rna[i*3:i*3+3])
        switch {
            default: result = append(result, protein)
            case err == ErrStop: return result, nil
            case err == ErrInvalidBase: return []string{}, err
        }
    }
    return result, nil
}

func FromCodon(codon string) (string, error) {
    protein, exists := translation[codon]
    switch {
        case !exists: return "", ErrInvalidBase
        case protein == "STOP": return "", ErrStop
    }
    return protein, nil
}
