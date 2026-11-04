package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	f, ok := cb[file]
	if !ok {
		return 0
	}
	var c int
	for _, e := range f {
		if e {
			c++
		}
	}
	return c
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	var c int
	for _, file := range cb {
		if file[rank-1] {
			c++
		}
	}
	return c
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var c int
	for _, file := range cb {
		for _ = range file {
			c++
		}
	}
	return c
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var c int
	for _, file := range cb {
		for _, w := range file {
			if w {
				c++
			}
		}
	}
	return c
}
