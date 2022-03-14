package sieve

func Sieve(limit int) []int {
    if limit < 2 {
        return []int{}
    }
	var numbers = make([]bool, limit + 1)
    var primes []int
    for i := 2; i <= limit; i++ {
        if numbers[i] == false {
            for j := 0; i * i + i * j <= limit; j++ {
                numbers[i * i + i * j] = true
            }
        }
    }
    for i, n := range numbers {
        if i > 1 && n == false {
            primes = append(primes, i)
        }
    }
    return primes
}
