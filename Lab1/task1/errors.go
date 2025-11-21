package task1

import "errors"

var (
	ErrInvalidRectangle = errors.New("Rectangle requires positive width and height")
	ErrNonPositiveSize  = errors.New("Width and height must remain positive")
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
