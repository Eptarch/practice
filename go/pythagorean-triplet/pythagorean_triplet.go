package pythagorean

import "math"

type Triplet [3]int

func coPrime(a, b int) bool {
	for b != 0 {
		a, b = b, a%b
	}
	return a == 1
}

func Range(min, max int) (result []Triplet) {
	maxM := int(math.Sqrt(float64(max)))
	for m := 2; m < maxM; m++ {
		for n := 1 + m%2; n < m; n += 2 {
			if !coPrime(n, m) {
				continue
			}
			m2, n2 := m*m, n*n
			a, b, c := m2-n2, 2*m*n, m2+n2
			for k := 1; ; k++ {
				if c*k > max {
					break
				}
				if a*k < min {
					continue
				}
				result = append(result, Triplet{k * a, k * b, k * c})
			}
		}
	}
	return
}

func Sum(p int) (result []Triplet) {
	for a := 1; a < p; a++ {
		b := (p*p - 2*a*p) / 2 / (p - a)
		if b < a {
			break
		}
		c := p - a - b
		if a*a+b*b == c*c {
			result = append(result, Triplet{a, b, c})
		}
	}
	return
}
