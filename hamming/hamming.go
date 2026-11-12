package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("strands must be of equal lengths")
	}
	d := 0
	for i, _ := range a {
		if a[i] != b[i] {
			d++
		}
	}
	return d, nil
}
