package railfence

func Encode(message string, rails int) string {
    return Swap(message, rails, true)
}

func Decode(message string, rails int) string {
	return Swap(message, rails, false)
}

func Swap(message string, rails int, encode bool) string {
    var insert, index, delta int
    var r, s *int
    switch encode {
        default:
            r = &insert
            s = &index
        case true:
            r = &index
            s = &insert

    }
    cycle := (rails - 1) * 2
    result := []byte(message)
    for rail := 0; rail < rails; rail++ {
        delta = rail * 2
        for insert = rail; insert < len(message); insert += delta {
            result[*r] = message[*s]
            index++
            if delta == cycle {
                continue
            }
            delta = cycle - delta
        }
    }
    return string(result)
}
