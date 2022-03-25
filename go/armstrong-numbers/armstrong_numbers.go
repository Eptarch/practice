package armstrong

import (
    "math"
    "strconv"
)

func IsNumber(n int) bool {
    var srepr = strconv.Itoa(n)
    var sum int
    for _, r := range srepr {
        sum += int(math.Pow(float64(int(r - '0')), float64(len(srepr))))
    }
    return n == sum
}
