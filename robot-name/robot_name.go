package robotname

import (
	"errors"
	"fmt"
)

const namespacesCapacity int = 26 * 26 * 10 * 10 * 10
var namespace map[int]bool = make(map[int]bool, namespacesCapacity)
var resetCounter int = 0
var ls string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Define the Robot type.
type Robot struct {
    name int
    named bool
}

func (r *Robot) Name() (string, error) {
    if r.named == false {
        if resetCounter == namespacesCapacity {
            return "", errors.New("namespace exhausted")
        }
        r.Reset()
    }
    lp, np := r.name / 1000, r.name % 1000  // Letter part, number part
    st, nd := lp / 26, lp % 26    // First letter, second letter   
    return fmt.Sprintf("%s%s%03d", string(ls[st]), string(ls[nd]), np), nil 
}

func (r *Robot) Reset() {
    if len(namespace) == 0 {
        for i := 0; i < namespacesCapacity; i++ {
            namespace[i] = false
        }
    }
    for k := range namespace {
        r.name = k
        break
    }
    delete(namespace, r.name)
	r.named = true
    resetCounter++
}

