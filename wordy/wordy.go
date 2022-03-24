package wordy

import (
	"regexp"
	"strconv"
)

func Answer(s string) (int, bool) {
	if match, _ := regexp.MatchString(`What is -?\d+(?: (?:plus|minus|divided by|multiplied by) -?\d+)*\?`, s); !match {
		return 0, false
	}
	tokens := regexp.MustCompile(`(plus|minus|divided|multiplied|-?\d+)`).FindAllString(s, -1)
	result, _ := strconv.Atoi(tokens[0])
	for i := 2; i < len(tokens); i += 2 {
		n, _ := strconv.Atoi(tokens[i])
		switch tokens[i-1] {
		case "plus": result += n
		case "minus": result -= n
		case "divided":	result /= n
		case "multiplied": result *= n
		}
	}
	return result, true
}

