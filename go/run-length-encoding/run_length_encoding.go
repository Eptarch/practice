package encode

import (
	"strconv"
	"strings"
	"unicode"
)

type Group struct {
    content rune
    amount int
}

func RunLengthEncode(input string) string {
    var sb strings.Builder
    var groups = []Group{}
    group := Group{}
    for i, current := range input {
        switch {
            default: group.amount++
            case i == 0: group = Group{content: current, amount: 1}
            case group.content != current:
                groups = append(groups, group)
                group = Group{content: current, amount: 1}
        }
        if i == len(input) - 1 {
            groups = append(groups, group)
        }
    }
    for _, item := range groups {
        if item.amount > 1 {
            sb.WriteString(strconv.Itoa(item.amount))
        }
        sb.WriteRune(item.content)
    }
    return sb.String()
}

func RunLengthDecode(input string) string {
	var sb strings.Builder
    var runes = []rune{}
    for _, current := range input {
        if unicode.IsDigit(current) {
            runes = append(runes, current)
        } else if !unicode.IsDigit(current) {
            if len(runes) > 0 {
                amount, _ := strconv.Atoi(string(runes))
                for i := 0; i < amount; i++ {
                    sb.WriteRune(current)
                }
                runes = []rune{}
            } else {
                sb.WriteRune(current)
            }
        }
    }
    return sb.String()
}
