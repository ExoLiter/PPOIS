package task2_test

import (
	"math"
	"testing"

	"lab1/task2"
)

func TestCoefficientAccessAndNormalization(t *testing.T) {
	p := task2.New([]float64{3, 0, 0})
	if p.Degree() != 0 {
		t.Fatalf("expected degree 0 after trimming, got %d", p.Degree())
	}
	if coeff := p.Coefficient(0); coeff != 3 {
		t.Fatalf("expected coefficient 3, got %f", coeff)
	}
	if coeff := p.Coefficient(5); coeff != 0 {
		t.Fatalf("out-of-range coefficient should be 0, got %f", coeff)
	}
}

func TestEvaluate(t *testing.T) {
	p := task2.New([]float64{1, -3, 2}) // 2x^2 - 3x + 1
	if got := p.Evaluate(2); math.Abs(got-3) > 1e-9 {
		t.Fatalf("expected 3, got %f", got)
	}
}

func TestAdditionAndAssignment(t *testing.T) {
	a := task2.New([]float64{1, 2, 3})
	b := task2.New([]float64{3, 2, 1, 4})
	sum := a.Add(b)
	compareCoeffs(t, sum.Coefficients(), []float64{4, 4, 4, 4})

	a.AddAssign(b)
	compareCoeffs(t, a.Coefficients(), []float64{4, 4, 4, 4})
}

func TestSubtractionAndAssignment(t *testing.T) {
	a := task2.New([]float64{5, 4, 3})
	b := task2.New([]float64{1, 2})
	diff := a.Subtract(b)
	compareCoeffs(t, diff.Coefficients(), []float64{4, 2, 3})

	a.SubtractAssign(b)
	compareCoeffs(t, a.Coefficients(), []float64{4, 2, 3})
}

func TestMultiplyAndAssignment(t *testing.T) {
	a := task2.New([]float64{1, 1})  // x + 1
	b := task2.New([]float64{-1, 1}) // x - 1
	product := a.Multiply(b)
	compareCoeffs(t, product.Coefficients(), []float64{-1, 0, 1})

	a.MultiplyAssign(b)
	compareCoeffs(t, a.Coefficients(), []float64{-1, 0, 1})
}

func TestDivideProducesQuotientAndRemainder(t *testing.T) {
	dividend := task2.New([]float64{-1, 0, 0, 1}) // x^3 - 1
	divisor := task2.New([]float64{-1, 1})        // x - 1
	quot, rem, err := dividend.Divide(divisor)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	compareCoeffs(t, quot.Coefficients(), []float64{1, 1, 1})
	compareCoeffs(t, rem.Coefficients(), []float64{0})
}

func TestDivideAssignReturnsRemainder(t *testing.T) {
	dividend := task2.New([]float64{2, -3, 1}) // x^2 -3x +2
	divisor := task2.New([]float64{-1, 1})     // x -1
	remainder, err := dividend.DivideAssign(divisor)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	compareCoeffs(t, dividend.Coefficients(), []float64{-2, 1})
	compareCoeffs(t, remainder.Coefficients(), []float64{0})
}

func TestDivideWithSmallerDegreeDividend(t *testing.T) {
	dividend := task2.New([]float64{1, 2})
	divisor := task2.New([]float64{0, 0, 1}) // x^2
	quot, rem, err := dividend.Divide(divisor)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	compareCoeffs(t, quot.Coefficients(), []float64{0})
	compareCoeffs(t, rem.Coefficients(), []float64{1, 2})
}

func TestDivideByZeroPolynomial(t *testing.T) {
	dividend := task2.New([]float64{1, 2})
	zero := task2.New([]float64{0})
	if _, _, err := dividend.Divide(zero); err != task2.ErrZeroDivisor {
		t.Fatalf("expected ErrZeroDivisor, got %v", err)
	}
}

func TestEqual(t *testing.T) {
	a := task2.New([]float64{1, 2, 3})
	b := task2.New([]float64{1, 2, 3})
	if !a.Equal(b) {
		t.Fatal("expected polynomials to be equal")
	}
	c := task2.New([]float64{1, 2})
	if a.Equal(c) {
		t.Fatal("expected polynomials to differ")
	}
}

func compareCoeffs(t *testing.T, got, want []float64) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("length mismatch: got %v want %v", got, want)
	}
	for i := range got {
		if math.Abs(got[i]-want[i]) > 1e-9 {
			t.Fatalf("coeff mismatch at %d: got %f, want %f", i, got[i], want[i])
		}
	}
}
