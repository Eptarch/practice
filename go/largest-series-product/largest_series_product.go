package lsproduct

import (
    "fmt"
    "unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
    switch {
        case len(digits) < span:
            return 0, fmt.Errorf("span less than the amount of digits")
        case span < 0:
            return 0, fmt.Errorf("span should be positive")
    }
    var maxProduct int64
    for i := 0; i <= len(digits) - span; i++ {
        var product int64 = 1
        for _, r := range digits[i:i + span] {
            if !unicode.IsDigit(r) {
                return 0, fmt.Errorf("digits contain non-digit")
            } else if int64(r) - '0' == 0 {
                product = 0
                break
            }
            product *= int64(r) - '0'
        }
        if product > maxProduct {
            maxProduct = product
        }
    }
    return maxProduct, nil
}
