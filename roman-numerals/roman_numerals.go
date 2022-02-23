package romannumerals

import (
    "fmt"
    "strings"
)

type Rule struct {
    arabic int
    roman string
}

var rules = []Rule{
    {arabic: 1000, roman: "M"},
    {arabic: 900,  roman: "CM"},
    {arabic: 500,  roman: "D"},
    {arabic: 400,  roman: "CD"},
    {arabic: 100,  roman: "C"},
    {arabic: 90,   roman: "XC"},
    {arabic: 50,   roman: "L"},
    {arabic: 40,   roman: "XL"},
    {arabic: 10,   roman: "X"},
    {arabic: 9,    roman: "IX"},
    {arabic: 5,    roman: "V"},
    {arabic: 4,    roman: "IV"},
    {arabic: 1,    roman: "I"},
}

func ToRomanNumeral(input int) (string, error) {
    if input <= 0 || input > 3000 {
        return "", fmt.Errorf("")
    }
	var sb strings.Builder
    for _, rule := range rules {
        for input >= rule.arabic {
            input -= rule.arabic
            sb.WriteString(rule.roman)
        }
    }
    return sb.String(), nil
}
