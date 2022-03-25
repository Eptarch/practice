package brackets

var t = map[rune]rune{'(': ')', '[': ']', '{': '}'}

func Bracket(input string) bool {
    stack := []rune{}
    for _, r := range input {
        switch r {
            case '(', '{', '[': stack = append(stack, r)
            case ')', '}', ']':
                if len(stack) == 0 || r != t[stack[len(stack) - 1]] {
                    return false
                }
                stack = stack[:len(stack) - 1]
        }
    }
    return len(stack) == 0
}

