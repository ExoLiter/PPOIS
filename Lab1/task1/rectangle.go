package task1

import "fmt"

// Rectangle models an axis-aligned rectangle with integer vertices.
// It is represented by the bottom-left (left, bottom) and top-right (right, top) corners.
type Rectangle struct {
	left   int
	bottom int
	right  int
	top    int
}

// NewRectangle constructs a rectangle from any two opposite vertices.
func NewRectangle(x1, y1, x2, y2 int) (Rectangle, error) {
	left, right := orderPair(x1, x2)
	bottom, top := orderPair(y1, y2)
	if left == right || bottom == top {
		return Rectangle{}, ErrInvalidRectangle
	}
	return Rectangle{left: left, bottom: bottom, right: right, top: top}, nil
}

// String implements fmt.Stringer to provide a compact textual representation.
func (r Rectangle) String() string {
	return fmt.Sprintf("[(%d,%d),(%d,%d)]", r.left, r.bottom, r.right, r.top)
}

// Equal reports whether two rectangles share identical coordinates.
func (r Rectangle) Equal(other Rectangle) bool {
	return r.left == other.left &&
		r.right == other.right &&
		r.top == other.top &&
		r.bottom == other.bottom
}

// Width returns the number of units along the X axis.
func (r Rectangle) Width() int {
	return r.right - r.left
}

// Height returns the number of units along the Y axis.
func (r Rectangle) Height() int {
	return r.top - r.bottom
}

// Vertices returns the vertices in the order: top-left, top-right, bottom-right, bottom-left.
func (r Rectangle) Vertices() [4]Point {
	return [4]Point{
		{X: r.left, Y: r.top},
		{X: r.right, Y: r.top},
		{X: r.right, Y: r.bottom},
		{X: r.left, Y: r.bottom},
	}
}

// Move shifts the rectangle by the provided offset.
func (r Rectangle) Move(dx, dy int) Rectangle {
	return Rectangle{
		left:   r.left + dx,
		right:  r.right + dx,
		top:    r.top + dy,
		bottom: r.bottom + dy,
	}
}

// Resize changes the width and height by deltaWidth and deltaHeight keeping the bottom-left vertex fixed.
func (r Rectangle) Resize(deltaWidth, deltaHeight int) (Rectangle, error) {
	newWidth := r.Width() + deltaWidth
	newHeight := r.Height() + deltaHeight
	if newWidth <= 0 || newHeight <= 0 {
		return Rectangle{}, ErrNonPositiveSize
	}
	return Rectangle{
		left:   r.left,
		bottom: r.bottom,
		right:  r.left + newWidth,
		top:    r.bottom + newHeight,
	}, nil
}

// Increment increases the size by one unit per axis keeping the bottom-left vertex fixed.
func (r Rectangle) Increment() Rectangle {
	return Rectangle{
		left:   r.left,
		bottom: r.bottom,
		right:  r.right + 1,
		top:    r.top + 1,
	}
}

// Decrement decreases the size by one unit per axis.
func (r Rectangle) Decrement() (Rectangle, error) {
	if r.Width() <= 1 || r.Height() <= 1 {
		return Rectangle{}, ErrNonPositiveSize
	}
	return Rectangle{
		left:   r.left,
		bottom: r.bottom,
		right:  r.right - 1,
		top:    r.top - 1,
	}, nil
}

// IncrementAssign expands the rectangle in place and returns the modified value.
func (r *Rectangle) IncrementAssign() {
	*r = r.Increment()
}

// DecrementAssign shrinks the rectangle by one unit along each axis.
func (r *Rectangle) DecrementAssign() error {
	dec, err := r.Decrement()
	if err != nil {
		return err
	}
	*r = dec
	return nil
}

// Union returns the minimal rectangle that contains both rectangles.
func (r Rectangle) Union(other Rectangle) Rectangle {
	left := min(r.left, other.left)
	right := max(r.right, other.right)
	bottom := min(r.bottom, other.bottom)
	top := max(r.top, other.top)
	return Rectangle{left: left, right: right, bottom: bottom, top: top}
}

// UnionAssign replaces the current rectangle with the union result.
func (r *Rectangle) UnionAssign(other Rectangle) {
	*r = r.Union(other)
}

// Intersection computes the overlapping area of two rectangles.
// The boolean result indicates whether a non-empty intersection exists.
func (r Rectangle) Intersection(other Rectangle) (Rectangle, bool) {
	left := max(r.left, other.left)
	right := min(r.right, other.right)
	bottom := max(r.bottom, other.bottom)
	top := min(r.top, other.top)
	if left >= right || bottom >= top {
		return Rectangle{}, false
	}
	return Rectangle{left: left, right: right, bottom: bottom, top: top}, true
}

// IntersectionAssign keeps the overlapping area in the receiver and reports whether it exists.
func (r *Rectangle) IntersectionAssign(other Rectangle) bool {
	if inter, ok := r.Intersection(other); ok {
		*r = inter
		return true
	}
	return false
}

// Clone returns a copy of the rectangle.
func (r Rectangle) Clone() Rectangle {
	return Rectangle{left: r.left, right: r.right, bottom: r.bottom, top: r.top}
}
