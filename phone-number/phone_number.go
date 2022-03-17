package phonenumber

import (
    "fmt"
	"errors"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
    digits := []rune{}
    for _, r := range phoneNumber {
        if unicode.IsDigit(r) {
            digits = append(digits, r)
        }
    }
    switch {
        case len(digits) > 10 && digits[0] != '1':
            return "", errors.New("invalid country code")
        case len(digits) < 10:
            return "", errors.New("not enough digits")
    }
    if len(digits) == 11 {
        digits = digits[1:]
    }
    switch {
        case digits[0] == '0' || digits[0] == '1':
            return "", errors.New("invalid area code")
        case digits[3] == '0' || digits[3] == '1':
            return "", errors.New("invalid exchange code")          
    }
    return string(digits), nil
}

func AreaCode(phoneNumber string) (string, error) {
    number, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
    number, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return fmt.Sprintf("(%s) %s-%s", number[:3], number[3:6], number[6:]), nil
}

