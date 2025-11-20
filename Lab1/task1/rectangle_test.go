package task1_test

import (
	"lab1/task1"
	"testing"
)

func TestNewRectangleValid(t *testing.T) {
	rect := mustRect(t, 3, 1, -1, 5)
	if rect.Width() != 4 || rect.Height() != 4 {
		t.Fatalf("unexpected size width=%d height=%d", rect.Width(), rect.Height())
	}
}

func TestNewRectangleInvalid(t *testing.T) {
	if _, err := task1.NewRectangle(0, 0, 0, 5); err != task1.ErrInvalidRectangle {
		t.Fatalf("expected ErrInvalidRectangle, got %v", err)
	}
}

func TestVerticesOrder(t *testing.T) {
	rect := mustRect(t, 0, 0, 2, 1)
	vertices := rect.Vertices()
	expected := [4]task1.Point{{X: 0, Y: 1}, {X: 2, Y: 1}, {X: 2, Y: 0}, {X: 0, Y: 0}}
	if vertices != expected {
		t.Fatalf("expected %v, got %v", expected, vertices)
	}
}

func TestMove(t *testing.T) {
	rect := mustRect(t, 0, 0, 1, 1)
	moved := rect.Move(3, -2)
	want := mustRect(t, 3, -2, 4, -1)
	if !moved.Equal(want) {
		t.Fatalf("expected %v, got %v", want, moved)
	}
}

func TestResize(t *testing.T) {
	rect := mustRect(t, 0, 0, 2, 2)
	resized, err := rect.Resize(2, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resized.Width() != 4 || resized.Height() != 3 {
		t.Fatalf("unexpected size width=%d height=%d", resized.Width(), resized.Height())
	}
}

func TestResizeError(t *testing.T) {
	rect := mustRect(t, 0, 0, 2, 2)
	if _, err := rect.Resize(-5, 0); err != task1.ErrNonPositiveSize {
		t.Fatalf("expected ErrNonPositiveSize, got %v", err)
	}
}

func TestIncrementAndDecrement(t *testing.T) {
	rect := mustRect(t, 0, 0, 2, 2)
	incremented := rect.Increment()
	if incremented.Width() != 3 || incremented.Height() != 3 {
		t.Fatalf("increment failed width=%d height=%d", incremented.Width(), incremented.Height())
	}
	decremented, err := incremented.Decrement()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !rect.Equal(decremented) {
		t.Fatalf("expected %v, got %v", rect, decremented)
	}
}

func TestDecrementError(t *testing.T) {
	rect := mustRect(t, 0, 0, 1, 2)
	if _, err := rect.Decrement(); err != task1.ErrNonPositiveSize {
		t.Fatalf("expected ErrNonPositiveSize, got %v", err)
	}
}

func TestDecrementAssignError(t *testing.T) {
	rect := mustRect(t, 0, 0, 1, 1)
	if err := rect.DecrementAssign(); err != task1.ErrNonPositiveSize {
		t.Fatalf("expected ErrNonPositiveSize, got %v", err)
	}
}

func TestUnionAndIntersection(t *testing.T) {
	a := mustRect(t, 0, 0, 3, 3)
	b := mustRect(t, 2, 2, 4, 4)
	union := a.Union(b)
	if union.Width() != 4 || union.Height() != 4 {
		t.Fatalf("unexpected union size width=%d height=%d", union.Width(), union.Height())
	}
	intersection, ok := a.Intersection(b)
	if !ok {
		t.Fatal("expected intersection")
	}
	expected := mustRect(t, 2, 2, 3, 3)
	if !expected.Equal(intersection) {
		t.Fatalf("expected %v, got %v", expected, intersection)
	}
}

func TestStringFormat(t *testing.T) {
	rect := mustRect(t, -1, -2, 3, 4)
	if rect.String() != "[(-1,-2),(3,4)]" {
		t.Fatalf("unexpected string output: %s", rect.String())
	}
}

func TestUnionAssign(t *testing.T) {
	a := mustRect(t, 0, 0, 1, 1)
	b := mustRect(t, -1, -1, 0, 0)
	a.UnionAssign(b)
	expected := mustRect(t, -1, -1, 1, 1)
	if !a.Equal(expected) {
		t.Fatalf("expected %v, got %v", expected, a)
	}
}

func TestIntersectionAssign(t *testing.T) {
	a := mustRect(t, 0, 0, 2, 2)
	b := mustRect(t, 1, 1, 3, 3)
	if !a.IntersectionAssign(b) {
		t.Fatal("expected intersection")
	}
	expected := mustRect(t, 1, 1, 2, 2)
	if !a.Equal(expected) {
		t.Fatalf("expected %v, got %v", expected, a)
	}

	far := mustRect(t, 5, 5, 6, 6)
	if a.IntersectionAssign(far) {
		t.Fatal("expected no intersection")
	}
}

func TestIncrementAndDecrementAssign(t *testing.T) {
	rect := mustRect(t, 0, 0, 2, 2)
	rect.IncrementAssign()
	if rect.Width() != 3 || rect.Height() != 3 {
		t.Fatalf("expected 3x3, got %dx%d", rect.Width(), rect.Height())
	}
	if err := rect.DecrementAssign(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rect.Width() != 2 || rect.Height() != 2 {
		t.Fatalf("expected 2x2, got %dx%d", rect.Width(), rect.Height())
	}
}

func TestCloneProducesIndependentCopy(t *testing.T) {
	original := mustRect(t, 0, 0, 3, 3)
	cloned := original.Clone()
	if !original.Equal(cloned) {
		t.Fatalf("clone should be equal")
	}
	cloned.IncrementAssign()
	if original.Equal(cloned) {
		t.Fatalf("expected clone to diverge after modification")
	}
}

func mustRect(t *testing.T, x1, y1, x2, y2 int) task1.Rectangle {
	t.Helper()
	rect, err := task1.NewRectangle(x1, y1, x2, y2)
	if err != nil {
		t.Fatalf("unexpected error creating rectangle: %v", err)
	}
	return rect
}
