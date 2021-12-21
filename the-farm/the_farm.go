package thefarm

import (
    "errors"
    "fmt"
)

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
    arg int
}

func (e *SillyNephewError) Error() string {
    return fmt.Sprintf("silly nephew, there cannot be %d cows", e.arg)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {

    if cows == 0 {
        return 0, errors.New("Division by zero")
    } else if cows < 0 {
        return 0, &SillyNephewError{arg: cows}
    }

    fodder, err := weightFodder.FodderAmount()
    
    if fodder < 0 {
        return 0, errors.New("Negative fodder")
    } else if err == ErrScaleMalfunction {
        fodder *= 2
    } else if err != nil {
        return 0, err
    }

    return fodder / float64(cows), nil
}
