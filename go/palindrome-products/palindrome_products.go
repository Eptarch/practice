package palindrome

import "errors"

type Product struct {
    Value int
    Factorizations [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
        return Product{}, Product{}, errors.New("fmin > fmax")
    }
    min, max := Product{Value: 9223372036854775807}, Product{Value: 0}
    for i := fmin; i <= fmax; i++ {
        for j := i; j <= fmax; j++ {
            product := i * j
            if isPalindrome(product) != true {
                continue
            }
            if product < min.Value {
                min = Product{Value: product, Factorizations: [][2]int{{i, j}}}
            } else if product == min.Value {
                min.Factorizations = append(min.Factorizations, [2]int{i, j})
            }
            if product > max.Value {
                max = Product{Value: product, Factorizations: [][2]int{{i, j}}}
            } else if product == max.Value {
                max.Factorizations = append(max.Factorizations, [2]int{i, j})
            }
        }    
    }
    if min.Value == 9223372036854775807 || max.Value == 0 {
        return Product{}, Product{}, errors.New("no palindromes")
    }
    return min, max, nil
}

func isPalindrome(n int) bool {
    reversed := 0
    for i := n; i > 0; i /= 10 {
        reversed = reversed * 10 + i % 10
    }
    return n == reversed
}

