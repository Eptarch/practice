package rotationalcipher

import "strings"

func RotationalCipher(plain string, shiftKey int) string {
    var sb strings.Builder
    for _, r := range plain {
        switch {
            default:
                sb.WriteRune(r)
            case r >= 0x41 && r <= 0x5a:
                sb.WriteRune((r - 0x41 + rune(shiftKey)) % 26 + 0x41)
            case r >= 0x61 && r <= 0x7a:
                sb.WriteRune((r - 0x61 + rune(shiftKey)) % 26 + 0x61)
        }          
    }
    return sb.String()
}
