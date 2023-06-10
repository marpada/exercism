package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
// CountInFile returns how many squares are occupied in the chessboard,
type Chessboard map[string]File

// within the given file.
func CountInFile(cb Chessboard, file string) int {
	f, exists := cb[file]
	if ! exists {
		return 0
	}
	var count int
	for _, square := range f {
		if square {
			count++
		}
	}
	return count

}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	var count int
	for _, file := range cb {
		if file[rank-1] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var count int
	for _, row := range cb {
		for range row {
			count++
		}
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var count int
	for _, row := range cb {
		for _, rank := range row {
			if rank {
				count++
			}

		}
	}
	return count
}
