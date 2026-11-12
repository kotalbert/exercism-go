package differenceofsquares

// inclusiveRange generate slice of numbers [1, n]
func inclusiveRange(n int) []int {
	numbers := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

func SquareOfSum(n int) int {
	s := 0
	for _, v := range inclusiveRange(n) {
		s += v
	}
	return s * s
}

func SumOfSquares(n int) int {
	s := 0
	for _, v := range inclusiveRange(n) {
		s += v * v
	}
	return s
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
