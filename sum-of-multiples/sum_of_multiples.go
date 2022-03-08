package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var result int
    for i := 1; i < limit; i++ { 
        for _, divisor := range divisors {
            if divisor != 0 && i % divisor == 0 {
                result += i
                break
            }
        }
    }
    return result
}
