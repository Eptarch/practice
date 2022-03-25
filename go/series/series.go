package series

func All(n int, s string) (result []string) {
    for i := 0; i <= len(s) - n; i++ {
        result = append(result, s[i:i + n])
    }
    return
}

func UnsafeFirst(n int, s string) string {
    result, ok := First(n, s)
    if ok {
        return result
    }
    panic(nil)
}

func First(n int, s string) (first string, ok bool) {
    if n > len(s) {
        return s, false
    }
    return s[0:n], true
}

