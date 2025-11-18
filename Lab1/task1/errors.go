package task1

import "errors"

// Error constants shared across rectangle operations.
var (
	ErrInvalidRectangle = errors.New("task1: rectangle requires positive width and height")
	ErrNonPositiveSize  = errors.New("task1: width and height must remain positive")
)

func orderPair(a, b int) (int, int) {
	if a <= b {
		return a, b
	}
	return b, a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
