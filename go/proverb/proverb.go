package proverb

import "fmt"

func Proverb(rhyme []string) (result []string) {
    var previous string
    for index, current := range rhyme {
        if previous != "" {
            result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", previous, current))
        }
        previous = current
        if index + 1 == len(rhyme) {
            result = append(result, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
        }
    }
	return
}
