package triangle

type Kind int

const (
    NaT = iota// not a triangle
    Equ // equilateral
    Iso // isosceles
    Sca // scalene
)
var types = []Kind{NaT, Equ, Iso, Sca}

func KindFromSides(a, b, c float64) Kind {
    if a <= 0 || b <= 0 || c <= 0 || a + b <= c || b + c <= a || c + a <= b {
        return NaT
    }
    return types[len(map[interface{}]float64{a: a, b: b, c: c})]
}
