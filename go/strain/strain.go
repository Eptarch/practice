package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) (ns Ints) {
    for _, n := range i {
        if filter(n) {
            ns = append(ns, n)
        }
    }
    return
}

func (i Ints) Discard(filter func(int) bool) Ints {
    return i.Keep(func(x int) bool { return !filter(x)})
}

func (l Lists) Keep(filter func([]int) bool) (ls Lists) {
    for _, lt := range l {
        if filter(lt) {
            ls = append(ls, lt)
        }
    }
    return
}

func (s Strings) Keep(filter func(string) bool) (ss Strings) {
    for _, st := range s {
        if filter(st) {
            ss = append(ss, st)
        }
    }
    return
}
