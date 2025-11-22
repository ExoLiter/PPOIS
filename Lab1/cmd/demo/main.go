package main

import (
	"fmt"

	"lab1/task1"
	"lab1/task2"
)

const (
	rectAX1               = 0
	rectAY1               = 0
	rectAX2               = 4
	rectAY2               = 3
	rectBX1               = 2
	rectBY1               = 1
	rectBX2               = 6
	rectBY2               = 5
	rectMoveDX            = 1
	rectMoveDY            = -2
	rectResizeDeltaWidth  = 2
	rectResizeDeltaHeight = 1
	polyEvalPoint         = 2.0
)

var (
	basePolynomial   = []float64{2, -3, 1}
	secondPolynomial = []float64{-1, 4, 0}
)

func main() {
	rectA, err := task1.NewRectangle(rectAX1, rectAY1, rectAX2, rectAY2)
	if err != nil {
		panic(err)
	}
	rectB, err := task1.NewRectangle(rectBX1, rectBY1, rectBX2, rectBY2)
	if err != nil {
		panic(err)
	}

	fmt.Println("Rectangle demo")
	fmt.Printf("A vertices: %v\n", rectA.Vertices())
	fmt.Printf("B vertices: %v\n", rectB.Vertices())
	fmt.Printf("A initial size: width=%d height=%d\n", rectA.Width(), rectA.Height())
	fmt.Printf("B initial size: width=%d height=%d\n", rectB.Width(), rectB.Height())

	movedA := rectA.Move(rectMoveDX, rectMoveDY)
	fmt.Printf("A Move by (%d,%d): %v\n", rectMoveDX, rectMoveDY, movedA.Vertices())

	movedB := rectB.Move(rectMoveDX, rectMoveDY)
	fmt.Printf("B Move by (%d,%d): %v\n", rectMoveDX, rectMoveDY, movedB.Vertices())

	if resized, err := rectA.Resize(rectResizeDeltaWidth, rectResizeDeltaHeight); err == nil {
		fmt.Printf("A Resize (+%d,+%d): width=%d height=%d\n",
			rectResizeDeltaWidth, rectResizeDeltaHeight, resized.Width(), resized.Height())
	}

	if resized, err := rectB.Resize(rectResizeDeltaWidth, rectResizeDeltaHeight); err == nil {
		fmt.Printf("B Resize (+%d,+%d): width=%d height=%d\n",
			rectResizeDeltaWidth, rectResizeDeltaHeight, resized.Width(), resized.Height())
	}

	expandedA := rectA.Clone()
	expandedA.IncrementAssign()
	fmt.Printf("A Increment ++: width=%d height=%d\n", expandedA.Width(), expandedA.Height())
	if err := expandedA.DecrementAssign(); err == nil {
		fmt.Printf("A Decrement --: width=%d height=%d\n", expandedA.Width(), expandedA.Height())
	}

	expandedB := rectB.Clone()
	expandedB.IncrementAssign()
	fmt.Printf("B Increment ++: width=%d height=%d\n", expandedB.Width(), expandedB.Height())
	if err := expandedB.DecrementAssign(); err == nil {
		fmt.Printf("B Decrement --: width=%d height=%d\n", expandedB.Width(), expandedB.Height())
	}

	union := rectA.Union(rectB)
	fmt.Printf("Union result: %v\n", union.Vertices())

	unionAssigned := rectA.Clone()
	unionAssigned.UnionAssign(rectB)
	fmt.Printf("UnionAssign result: %v\n", unionAssigned.Vertices())

	if intersection, ok := rectA.Intersection(rectB); ok {
		fmt.Printf("Intersection result: %v\n", intersection.Vertices())
	}

	intersectionAssigned := rectB.Clone()
	if intersectionAssigned.IntersectionAssign(rectA) {
		fmt.Printf("IntersectionAssign result: %v\n", intersectionAssigned.Vertices())
	} else {
		fmt.Println("IntersectionAssign: rectangles do not intersect")
	}

	fmt.Println()

	p1 := task2.New(basePolynomial)
	p2 := task2.New(secondPolynomial)

	fmt.Println("Polynomial demo")
	fmt.Printf("p1 coefficients: %v\n", p1.Coefficients())
	fmt.Printf("p2 coefficients: %v\n", p2.Coefficients())
	fmt.Printf("p1 coefficient for x^1: %.1f\n", p1.Coefficient(1))
	fmt.Printf("p1(%.1f) using operator (): %.1f\n", polyEvalPoint, p1.Evaluate(polyEvalPoint))

	sum := p1.Add(p2)
	fmt.Printf("p1 + p2: %v\n", sum.Coefficients())
	sumAssigned := p1.Clone()
	sumAssigned.AddAssign(p2)
	fmt.Printf("p1 += p2: %v\n", sumAssigned.Coefficients())

	diff := p1.Subtract(p2)
	fmt.Printf("p1 - p2: %v\n", diff.Coefficients())
	diffAssigned := p1.Clone()
	diffAssigned.SubtractAssign(p2)
	fmt.Printf("p1 -= p2: %v\n", diffAssigned.Coefficients())

	product := p1.Multiply(p2)
	fmt.Printf("p1 * p2: %v\n", product.Coefficients())
	productAssigned := p1.Clone()
	productAssigned.MultiplyAssign(p2)
	fmt.Printf("p1 *= p2: %v\n", productAssigned.Coefficients())

	quotient, remainder, err := product.Divide(p1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("(p1*p2) / p1: quotient=%v remainder=%v\n", quotient.Coefficients(), remainder.Coefficients())

	divAssigned := product.Clone()
	rem, err := divAssigned.DivideAssign(p2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("(p1*p2) /= p2: quotient=%v remainder=%v\n", divAssigned.Coefficients(), rem.Coefficients())
}
