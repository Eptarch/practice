package atbash

import "strings"

func Atbash(s string) string {
	var sb strings.Builder
    for _, r := range strings.ToLower(s) {
        switch {
            case r >= 0x30 && r <= 0x39: sb.WriteRune(r)
            case r >= 0x41 && r <= 0x5a: sb.WriteRune(0x5a - r + 0x41)
            case r >= 0x61 && r <= 0x7a: sb.WriteRune(0x7a - r + 0x61)
        }
        if sb.Len() % 6 == 5 {
            sb.WriteRune(0x20)
        }
    }
    return strings.TrimSpace(sb.String())
}
