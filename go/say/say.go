package say

import "strings"

type Scale struct {
    n int64
    name string
}

var scales = []Scale {
	{1e18, "quintillion"},
	{1e15, "quadrillion"},
	{1e12, "trillion"},
	{1e9, "billion"},
	{1e6, "million"},
	{1e3, "thousand"},
	{1e2, "hundred"},
}

var verbatim = []string{
	0:  "zero",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func Say(n int64) (string, bool) {
	switch {
        case n < 0, n >= 1e12:
            return "", false
        case n < int64(len(verbatim)) && verbatim[n] != "":
            return verbatim[n], true
        case n < 100:
            return verbatim[n - n % 10] + "-" + verbatim[n % 10], true
    }
    parts := make([]string, 0)
    for _, scale := range scales {
        if n < scale.n {
            continue
        }
        said, _ := Say(n / scale.n)
        parts = append(parts, said, scale.name)
        n %= scale.n
    }
    if n != 0 {
        said, _ := Say(n)
        parts = append(parts, said)
    }
    return strings.Join(parts, " "), true
}

