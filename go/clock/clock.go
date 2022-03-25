package clock

import "fmt"

type Clock struct {
    m int
}

func (c Clock) Add(m int) Clock {
	c.m = rotate(c.m + m)
    return c
}

func (c Clock) Subtract(m int) Clock {
	c.m = rotate(c.m - m)
    return c
}

func (c Clock) String() string {
    return fmt.Sprintf("%02d:%02d", c.m / 60 , c.m % 60)
}

func New(h, m int) Clock {
    return Clock{ m: rotate(60 * h + m) }
}

func rotate(m int) int {
    if m < 0 {
        return 1440 + m % 1440
    }
    return m % 1440
}

