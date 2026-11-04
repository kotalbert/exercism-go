package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, nc int) (float64, error) {
	// ignore errors for now
	fa, err := fc.FodderAmount(nc)
	if err != nil {
		return 0, err
	}
	ff, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return (fa * ff) / float64(nc), nil

}

func ValidateInputAndDivideFood(fc FodderCalculator, nc int) (float64, error) {
	if nc <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, nc)
}

func ValidateNumberOfCows(nc int) error {
	switch {
	case nc < 0:
		return &InvalidCowsError{
			nc:      nc,
			message: "there are no negative cows",
		}
	case nc == 0:
		return &InvalidCowsError{
			nc:      nc,
			message: "no cows don't need food",
		}
	default:
		return nil
	}
}

type InvalidCowsError struct {
	nc      int
	message string
}

func (e InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.nc, e.message)
}
