package thefarm

import "fmt"

// Define the SillyNephewError type.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood tries to split the fodder among the cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
    switch {
        case fodder > 0 && err == ErrScaleMalfunction:
	        return 2 * fodder / float64(cows), nil
        case fodder < 0 && err == ErrScaleMalfunction:
            return 0, fmt.Errorf("negative fodder")
        case err != nil:
	        return 0, err
        case fodder < 0:
            return 0, fmt.Errorf("negative fodder")
        case cows == 0:
            return 0, fmt.Errorf("division by zero")
        case cows < 0:
            return 0, &SillyNephewError{cows: cows}
    }
    return fodder / float64(cows), nil
}
