package task2

import (
	"errors"
	"math"
)

type Polynomial struct {
	coeffs []float64
}

var ErrZeroDivisor = errors.New("task2: division by zero polynomial")

func New(coefficients []float64) Polynomial {
	copied := append([]float64(nil), coefficients...)
	return Polynomial{coeffs: normalize(copied)}
}

func (p Polynomial) Coefficients() []float64 {
	return append([]float64(nil), p.coeffs...)
}

func (p Polynomial) Clone() Polynomial {
	return Polynomial{coeffs: append([]float64(nil), p.coeffs...)}
}

func (p Polynomial) Degree() int {
	return len(p.coeffs) - 1
}

func (p Polynomial) Coefficient(power int) float64 {
	if power < 0 || power >= len(p.coeffs) {
		return 0
	}
	return p.coeffs[power]
}

func (p Polynomial) Evaluate(x float64) float64 {
	result := 0.0
	for i := len(p.coeffs) - 1; i >= 0; i-- {
		result = result*x + p.coeffs[i]
	}
	return result
}

func (p Polynomial) Add(other Polynomial) Polynomial {
	maxLen := max(len(p.coeffs), len(other.coeffs))
	result := make([]float64, maxLen)
	for i := 0; i < maxLen; i++ {
		result[i] = p.Coefficient(i) + other.Coefficient(i)
	}
	return Polynomial{coeffs: normalize(result)}
}

func (p *Polynomial) AddAssign(other Polynomial) {
	*p = p.Add(other)
}

func (p Polynomial) Subtract(other Polynomial) Polynomial {
	maxLen := max(len(p.coeffs), len(other.coeffs))
	result := make([]float64, maxLen)
	for i := 0; i < maxLen; i++ {
		result[i] = p.Coefficient(i) - other.Coefficient(i)
	}
	return Polynomial{coeffs: normalize(result)}
}

func (p *Polynomial) SubtractAssign(other Polynomial) {
	*p = p.Subtract(other)
}

func (p Polynomial) Multiply(other Polynomial) Polynomial {
	result := make([]float64, p.Degree()+other.Degree()+2)
	for i := range p.coeffs {
		for j := range other.coeffs {
			result[i+j] += p.coeffs[i] * other.coeffs[j]
		}
	}
	return Polynomial{coeffs: normalize(result)}
}

func (p *Polynomial) MultiplyAssign(other Polynomial) {
	*p = p.Multiply(other)
}

func (p Polynomial) Divide(divisor Polynomial) (Polynomial, Polynomial, error) {
	if divisor.isZero() {
		return Polynomial{}, Polynomial{}, ErrZeroDivisor
	}
	dividend := Polynomial{coeffs: append([]float64(nil), p.coeffs...)}
	if dividend.Degree() < divisor.Degree() {
		return Polynomial{coeffs: []float64{0}}, dividend, nil
	}

	quotient := make([]float64, dividend.Degree()-divisor.Degree()+1)
	for dividend.Degree() >= divisor.Degree() {
		power := dividend.Degree() - divisor.Degree()
		scale := dividend.leading() / divisor.leading()
		quotient[power] = scale

		for j := range divisor.coeffs {
			idx := power + j
			dividend.coeffs[idx] -= scale * divisor.coeffs[j]
		}
		dividend.coeffs = normalize(dividend.coeffs)
	}
	return Polynomial{coeffs: normalize(quotient)}, Polynomial{coeffs: normalize(dividend.coeffs)}, nil
}

func (p *Polynomial) DivideAssign(divisor Polynomial) (Polynomial, error) {
	quot, remainder, err := p.Divide(divisor)
	if err != nil {
		return Polynomial{}, err
	}
	*p = quot
	return remainder, nil
}

func (p Polynomial) Equal(other Polynomial) bool {
	if len(p.coeffs) != len(other.coeffs) {
		return false
	}
	for i := range p.coeffs {
		if !almostEqual(p.coeffs[i], other.coeffs[i]) {
			return false
		}
	}
	return true
}

func (p Polynomial) leading() float64 {
	return p.coeffs[len(p.coeffs)-1]
}

func (p Polynomial) isZero() bool {
	return len(p.coeffs) == 1 && almostEqual(p.coeffs[0], 0)
}

func normalize(coeffs []float64) []float64 {
	n := len(coeffs)
	for n > 1 && almostEqual(coeffs[n-1], 0) {
		n--
	}
	return append([]float64(nil), coeffs[:n]...)
}

func almostEqual(a, b float64) bool {
	const tolerance = 1e-9
	return math.Abs(a-b) <= tolerance
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
