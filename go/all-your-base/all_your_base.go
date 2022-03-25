package allyourbase

import (
	"fmt"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
    switch {
        case inputBase < 2: return []int{}, fmt.Errorf("input base must be >= 2")   
        case outputBase < 2: return []int{}, fmt.Errorf("output base must be >= 2")
        case len(inputDigits) == 0: return []int{0}, nil
    }
    var repr int
    var result []int
    for i, d := range inputDigits {
        if d < 0 || d >= inputBase {
            return []int{}, fmt.Errorf("all digits must satisfy 0 <= d < input base")
        }
        if d == 0 && len(result) == 0 {
            continue
        }
        repr += d * int(math.Pow(float64(inputBase), float64(len(inputDigits) - i - 1)))
    }

    for repr > 0 {
        result = append([]int{repr % outputBase}, result...)
        repr = repr / outputBase
    }
    if len(result) == 0 {
        return []int{0}, nil
    }
    return result, nil
}

