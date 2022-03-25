package perfect

import (
	"errors"
	"math"
)

type Classification uint8

const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
)

var ErrOnlyPositive = errors.New("ErrOnlyPositive")

func Classify(number int64) (Classification, error) {
	if number < 1 {
		return ClassificationDeficient, ErrOnlyPositive
	}
	sum := AliquotSum(number)
	switch {
	    default: return ClassificationDeficient, nil
	    case sum == number: return ClassificationPerfect, nil
	    case sum > number: return ClassificationAbundant, nil
	}
}

func AliquotSum(number int64) int64 {
	sum := int64(1)
	p := int64(math.Sqrt(float64(number)))
	c := int64(2)
	for number > 1 && c <= p {
		ex := float64(1)
		for number % c == 0 {
			number /= c
			ex++
		}
		if ex > 1 {
			sum *= int64(math.Pow(float64(c), ex) - 1) / (c - 1)
		}
		c++
	}
	if number > 1 {
		sum *= (number * number - 1) / (number - 1)
	}
	return sum / 2
}
