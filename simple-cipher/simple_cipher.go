package cipher

import (
	"strings"
	"unicode"
)

type shift struct {
    distance int
}

type vigenere struct {
    key string
}

func NewCaesar() Cipher {
    return shift{distance: 3}
}

func NewShift(distance int) Cipher {
    if distance < -25 || distance == 0 || distance > 25 {
        return nil
    }
    return shift{distance: distance}
}

func (c shift) Encode(input string) string {
    return shiftTranslate(normalize(input), c.distance)
}

func (c shift) Decode(input string) string {
    return shiftTranslate(input, -c.distance)
}

func NewVigenere(key string) Cipher {
    var diff int = 0
    for _, r := range key {
        if r < 'a' || r > 'z' {
            return nil
        }
        diff += int(r - 'a')
    }
    if diff <= 0 {
        return nil
    }
    return vigenere{key: key}
}

func (v vigenere) Encode(input string) string {
	return vigenereTranslate(normalize(input), v.key, true)
}

func (v vigenere) Decode(input string) string {
	return vigenereTranslate(normalize(input), v.key, false)
}

func normalize(text string) string {
    var normalized strings.Builder
    for _, r := range text {
        if unicode.IsLetter(r) {
            normalized.WriteRune(unicode.ToLower(r))
        }
    }
	return normalized.String()
}

func rotate(r rune, distance int) rune {
	var n rune = r + rune(distance)
	if n > 'z' {
		n -= 'z' - 'a' + 1
	} else if n < 'a' {
		n += 'z' - 'a' + 1
	}
	return n
}

func shiftTranslate(s string, distance int) string {
    var sb strings.Builder
    for _, r := range s {
        sb.WriteRune(rotate(r, distance))
    }
    return sb.String()
}

func vigenereTranslate(s string, key string, encode bool) string {
    var sb strings.Builder
	for i, r := range(s) {
		if key[i % len(key)] == 'a' {
			sb.WriteRune(r)
            continue
		} else {
            c := NewShift(int(key[i % len(key)] - 'a'))
            if encode {
                sb.WriteString(c.Encode(string(r)))
            } else {
                sb.WriteString(c.Decode(string(r)))
            }
        }
	}
    return sb.String()
}
