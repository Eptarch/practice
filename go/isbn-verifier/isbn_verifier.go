package isbn

import "unicode"

func IsValidISBN(isbn string) bool {
	var numbers = []int{}
    var sum int
    for i, r := range isbn {
        switch {
            case unicode.IsDigit(r): numbers = append(numbers, int(r) - '0')
            case r == 'X' && i == len(isbn) - 1: numbers = append(numbers, 10)
        }
    }
    if len(numbers) != 10 {
        return false
    }
    for i := 0; i < 10; i++ {
        sum += numbers[i] * (10 - i)
    }
    return sum % 11 == 0
}
