package bowling

import "errors"

// Define the Game type here.
type Game struct{
    rolls []int
    secondFrame bool
}

func NewGame() *Game {
    return &Game{rolls: make([]int, 0, 21)}
}

func (g *Game) Roll(pins int) error {
    pinsOut := 0
    if g.secondFrame == true {
        pinsOut = g.rolls[len(g.rolls) - 1]
    }
    if _, err := g.Score(); (err == nil) || pins < 0 || pins + pinsOut > 10 {
        return errors.New("can't roll like that")
    }
    g.secondFrame, g.rolls = !g.secondFrame, append(g.rolls, pins)
    if pins == 10 {
        g.secondFrame = false
    }
	return nil
}

func (g *Game) Score() (result int, err error) {
    for roll, frame := 0, 0; frame < 10; roll, frame = roll + 1, frame + 1 {
        if len(g.rolls) <= roll + 1 {
            return 0, errors.New("insufficient rolls")
        }
        first, second := g.rolls[roll], g.rolls[roll + 1]
        result += first + second
        if first == 10 || first + second == 10 {
            if len(g.rolls) <= roll + 2 {
                return 0, errors.New("insufficient rolls")
            }
            result += g.rolls[roll + 2]
        }
        if first != 10 {
            roll++
        }
    }
	return
}
