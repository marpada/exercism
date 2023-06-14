package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError struct {
	cows int
}

func (error SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", error.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}

	amount, err := weightFodder.FodderAmount()

	if err != nil {
		if err == ErrScaleMalfunction {
			amount *= 2
			if amount < 0 {
				return 0.0, errors.New("negative fodder")
			}
			return amount / float64(cows), nil
		}
		return 0, err
	}

	if amount < 0 {
		return 0.0, errors.New("negative fodder")
	}

	if cows < 0 {
		return 0, SillyNephewError{cows: cows}
	}

	return amount / float64(cows), nil
}
