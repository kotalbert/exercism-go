package thefarm

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
	panic("Implement")
}

func ValidateNumberOfCows(nc int) error {
	panic("Implement")
}
