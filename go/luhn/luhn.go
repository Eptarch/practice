package luhn

import "unicode"

func Reverse(s string) string {
    n := len(s)
    rs := make([]rune, n)
    for _, r := range s {
        n--
        rs[n] = r
    }
    return string(rs[n:])
}

func Valid(id string) bool {
    double := true
    prep := []int{}
    for _, r := range Reverse(id) {
        if unicode.IsDigit(r) {
            double = !double
            d := int(r-'0')
            if double {
                d *= 2
                if d > 9 {
                    d -= 9
                }
            }
            prep = append(prep, d)
        } else if unicode.IsSpace(r) {
            continue
        } else {
            return false
        }
    }
    if len(prep) < 2 {
        return false
    }
    s := 0
    for _, v := range prep  {
        s += v
    }
    return s % 10 == 0
}
