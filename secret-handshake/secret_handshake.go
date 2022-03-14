package secret

var actions = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(code uint) []string {
    if code <= 0 {
        return []string{}
    }
    var result = make([]string, 0, len(actions))
    reverse := code / 16 % 2 == 1
    code %= 16
    bytes := byte(code)
    for i := 0; bytes > 0; i++ {
        if bytes&1 == 1 {
            result = append(result, actions[i])
        }
        bytes >>= 1
    }
    if reverse {
        last := len(result) - 1
        for i := 0; i < len(result) / 2; i++ {
            result[i], result[last - i] = result[last - i], result[i]
        }
    }
    return result
}

