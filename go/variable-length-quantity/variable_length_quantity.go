package variablelengthquantity

import "fmt"

func EncodeVarint(input []uint32) (result []byte) {
    for _, n := range input {
        encoded := []byte{byte(n&0x7f)}
        for n >>= 7; n != 0; n >>= 7 {
            encoded = append([]byte{byte(n&0x7f) | 0x80}, encoded...) 
        }
        result = append(result, encoded...)
    }
    return 
}

func DecodeVarint(input []byte) (result []uint32, err error) {
    var number uint32
    var complete bool
    for _, n := range input {
        number = (number << 7) + uint32(n&0x7f)
        if complete = (n&0x80 == 0); complete {
            result = append(result, number)
            number = 0
        }
    }
    if !complete {
        return []uint32{}, fmt.Errorf("incomplete sequence")
    }
    return
}
