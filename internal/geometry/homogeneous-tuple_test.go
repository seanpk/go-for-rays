package geometry

import (
	"math"
	"testing"
)

func TestTuple(t *testing.T) {
	tests := []struct {
		name     string
		tuple    HomogeneousTuple
		expected [2]bool
	}{
		{name: "w=1 is a point", tuple: NewTuple(1, 2, 3, 1), expected: [2]bool{true, false}},
		{name: "w=0 is a vector", tuple: NewTuple(1, 2, 3, 0), expected: [2]bool{false, true}},
		{name: "w!=0 and w!=1 is neither", tuple: NewTuple(1, 2, 3, 2), expected: [2]bool{false, false}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tuple.IsPoint(); got != tt.expected[0] {
				t.Errorf("IsPoint() = %v, want %v", got, tt.expected[0])
			}
			if got := tt.tuple.IsVector(); got != tt.expected[1] {
				t.Errorf("IsVector() = %v, want %v", got, tt.expected[1])
			}
		})
	}
}

func TestTupleConversion(t *testing.T) {
	tests := []struct {
		name     string
		tuple    HomogeneousTuple
		expected string
	}{
		{name: "convert point tuple", tuple: NewTuple(1, 2, 3, 1), expected: "Point(1.000000, 2.000000, 3.000000)"},
		{name: "convert vector tuple", tuple: NewTuple(1, 2, 3, 0), expected: "Vector(1.000000, 2.000000, 3.000000)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.tuple.IsPoint() {
				if got := tt.tuple.String(); got != tt.expected {
					t.Errorf("ToPoint() = %v, want %v", got, tt.expected)
				}
			} else if tt.tuple.IsVector() {
				if got := tt.tuple.String(); got != tt.expected {
					t.Errorf("ToVector() = %v, want %v", got, tt.expected)
				}
			}
		})
	}
}

func TestEquality(t *testing.T) {
	tests := []struct {
		name     string
		a        HomogeneousTuple
		b        HomogeneousTuple
		epsilon  *float64
		expected bool
	}{
		{name: "equal tuples", a: NewTuple(1, 2, 3, 5), b: NewTuple(1, 2, 3, 5), expected: true},
		{name: "unequal tuples", a: NewTuple(1, 2, 3, 5), b: NewTuple(1, 2, 4, -5), expected: false},
		{name: "equal vectors (w tuple)", a: NewVector(1, 2, 3), b: NewTuple(1, 2, 3, 0), expected: true},
		{name: "unequal vectors", a: NewVector(1, 2, 3), b: NewVector(1, 2, 4), expected: false},
		{name: "equal points (w tuple)", a: NewPoint(1, 2, 3), b: NewTuple(1, 2, 3, 1), expected: true},
		{name: "unequal points", a: NewPoint(1, 2, 3), b: NewPoint(1, 2, 4), expected: false},
		{name: "equal with epsilon", a: NewTuple(1, 2, 3, 1), b: NewTuple(1.0000001, 2, 3, 1), epsilon: &[]float64{0.000001}[0], expected: true},
		{name: "unequal with epsilon", a: NewTuple(1, 2, 3, 1), b: NewTuple(1.0001, 2, 3, 1), epsilon: &[]float64{0.000001}[0], expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result bool
			if tt.epsilon == nil {
				result = tt.a.Equals(tt.b)
			} else {
				result = tt.a.Equals(tt.b, *tt.epsilon)
			}
			if got := result; got != tt.expected {
				t.Errorf("IsEqual() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAddition(t *testing.T) {
	tests := []struct {
		name     string
		a        HomogeneousTuple
		b        HomogeneousTuple
		expected string
	}{
		{name: "add point and vector (is a point)", a: NewPoint(1, 2, 3), b: NewVector(4, 5, 6), expected: "Point(5.000000, 7.000000, 9.000000)"},
		{name: "add vector and point (is a point)", a: NewVector(4, 5, 6), b: NewPoint(1, 2, 3), expected: "Point(5.000000, 7.000000, 9.000000)"},
		{name: "add two vectors (is a vector)", a: NewVector(1, 2, 3), b: NewVector(4, 5, 6), expected: "Vector(5.000000, 7.000000, 9.000000)"},
		{name: "add two points (is a vector)", a: NewPoint(1, 2, 3), b: NewPoint(4, 5, 6), expected: "Vector(5.000000, 7.000000, 9.000000)"},
		{name: "tuple point addition", a: NewTuple(1, 2, 3, 1), b: NewTuple(4, 5, 6, 1), expected: "Vector(5.000000, 7.000000, 9.000000)"},
		{name: "tuple vector addition", a: NewTuple(1, 2, 3, 0), b: NewTuple(4, 5, 6, 0), expected: "Vector(5.000000, 7.000000, 9.000000)"},
		{name: "tuple point and vector addition", a: NewTuple(1, 2, 3, 1), b: NewTuple(4, 5, 6, 0), expected: "Point(5.000000, 7.000000, 9.000000)"},
		{name: "addition of irregular tuples is allowed", a: NewTuple(1, 2, 3, 3), b: NewTuple(4, 5, 6, 5), expected: "Tuple(5.000000, 7.000000, 9.000000, 8.000000)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.a.Add(tt.b)
			if got := result.String(); got != tt.expected {
				t.Errorf("Add() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSubtraction(t *testing.T) {
	tests := []struct {
		name     string
		a        HomogeneousTuple
		b        HomogeneousTuple
		expected string
	}{
		{name: "subtract point from vector (is a point)", a: NewVector(5, 7, 9), b: NewPoint(4, 5, 6), expected: "Point(1.000000, 2.000000, 3.000000)"},
		{name: "subtract vector from point (is a point)", a: NewPoint(5, 7, 9), b: NewVector(1, 2, 3), expected: "Point(4.000000, 5.000000, 6.000000)"},
		{name: "subtract two vectors (is a vector)", a: NewVector(5, 7, 9), b: NewVector(4, 5, 6), expected: "Vector(1.000000, 2.000000, 3.000000)"},
		{name: "subtract two points (is a vector)", a: NewPoint(5, 7, 9), b: NewPoint(4, 5, 6), expected: "Vector(1.000000, 2.000000, 3.000000)"},
		{name: "tuple point subtraction", a: NewTuple(5, 7, 9, 1), b: NewTuple(4, 5, 6, 1), expected: "Vector(1.000000, 2.000000, 3.000000)"},
		{name: "tuple vector subtraction", a: NewTuple(5, 7, 9, 0), b: NewTuple(4, 5, 6, 0), expected: "Vector(1.000000, 2.000000, 3.000000)"},
		{name: "tuple point and vector subtraction", a: NewTuple(5, 7, 9, 1), b: NewTuple(4, 5, 6, 0), expected: "Point(1.000000, 2.000000, 3.000000)"},
		{name: "subtraction of irregular tuples is allowed", a: NewTuple(5, 7, 9, 3), b: NewTuple(4, 5, 6, 5), expected: "Tuple(1.000000, 2.000000, 3.000000, -2.000000)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.a.Subtract(tt.b)
			if got := result.String(); got != tt.expected {
				t.Errorf("Subtract() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNegation(t *testing.T) {
	tests := []struct {
		name     string
		tuple    HomogeneousTuple
		expected string
	}{
		{name: "negate point", tuple: NewPoint(1, 2, 3), expected: "Point(-1.000000, -2.000000, -3.000000)"},
		{name: "negate vector", tuple: NewVector(1, 2, 3), expected: "Vector(-1.000000, -2.000000, -3.000000)"},
		{name: "negate tuple point", tuple: NewTuple(1, 2, 3, 1), expected: "Point(-1.000000, -2.000000, -3.000000)"},
		{name: "negate tuple vector", tuple: NewTuple(1, 2, 3, 0), expected: "Vector(-1.000000, -2.000000, -3.000000)"},
		{name: "negate irregular tuple", tuple: NewTuple(1, 2, 3, 2), expected: "Tuple(-1.000000, -2.000000, -3.000000, -2.000000)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rInverse := tt.tuple.Negate()
			if got := rInverse.String(); got != tt.expected {
				t.Errorf("Negated() = %v, want %v", got, tt.expected)
			}
			if rInverse.Equals(tt.tuple) {
				t.Errorf("Negated() modified the original Tuple!")
			}
		})
	}
}

func TestMultiplication(t *testing.T) {
	tests := []struct {
		name     string
		tuple    HomogeneousTuple
		factor   float64
		expected string
	}{
		{name: "multiply point by scalar", tuple: NewPoint(1, 2, 3), factor: 2, expected: "Point(2.000000, 4.000000, 6.000000)"},
		{name: "multiply vector by scalar", tuple: NewVector(1, 2, 3), factor: 2, expected: "Vector(2.000000, 4.000000, 6.000000)"},
		{name: "multiply tuple point by scalar", tuple: NewTuple(1, 2, 3, 1), factor: 2, expected: "Point(2.000000, 4.000000, 6.000000)"},
		{name: "multiply tuple vector by scalar", tuple: NewTuple(1, 2, 3, 0), factor: 2, expected: "Vector(2.000000, 4.000000, 6.000000)"},
		{name: "multiply irregular tuple by scalar", tuple: NewTuple(1, 2, 3, 5), factor: 2, expected: "Tuple(2.000000, 4.000000, 6.000000, 10.000000)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.tuple.Multiply(tt.factor)
			if got := result.String(); got != tt.expected {
				t.Errorf("Multiply() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDivision(t *testing.T) {
	tests := []struct {
		name     string
		tuple    HomogeneousTuple
		factor   float64
		expected string
	}{
		{name: "divide point by scalar", tuple: NewPoint(2, 4, 6), factor: 2, expected: "Point(1.000000, 2.000000, 3.000000)"},
		{name: "divide vector by scalar", tuple: NewVector(2, 4, 6), factor: 2, expected: "Vector(1.000000, 2.000000, 3.000000)"},
		{name: "divide tuple point by scalar", tuple: NewTuple(2, 4, 6, 1), factor: 2, expected: "Point(1.000000, 2.000000, 3.000000)"},
		{name: "divide tuple vector by scalar", tuple: NewTuple(2, 4, 6, 0), factor: 2, expected: "Vector(1.000000, 2.000000, 3.000000)"},
		{name: "divide irregular tuple by scalar", tuple: NewTuple(2, 4, 6, 10), factor: 2, expected: "Tuple(1.000000, 2.000000, 3.000000, 5.000000)"},
		{name: "divide by zero", tuple: NewTuple(2, 4, 6, 10), factor: 0, expected: "Tuple(NaN, NaN, NaN, NaN)"}, // Division by zero should return NaN
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.tuple.Divide(tt.factor)
			if got := result.String(); got != tt.expected {
				t.Errorf("Divide() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestMagnitude(t *testing.T) {
	tests := []struct {
		name     string
		tuple    HomogeneousTuple
		expected float64
	}{
		{name: "magnitude of point", tuple: NewPoint(2, 3, 6), expected: 0.0},
		{name: "magnitude of vector", tuple: NewVector(2, 3, 6), expected: 7.0},
		{name: "magnitude of tuple point", tuple: NewTuple(2, 3, 6, 1), expected: 0.0},
		{name: "magnitude of tuple vector", tuple: NewTuple(2, 3, 6, 0), expected: 7.0},
		{name: "magnitude of irregular tuple", tuple: NewTuple(0, 3, 4, 12), expected: 13.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tuple.Magnitude(); got != tt.expected {
				t.Errorf("Magnitude() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNormalization(t *testing.T) {
	tests := []struct {
		name     string
		tuple    HomogeneousTuple
		expected HomogeneousTuple
	}{
		{name: "normalize point", tuple: NewPoint(2, 3, 6), expected: NewTuple(math.NaN(), math.NaN(), math.NaN(), math.NaN())},
		{name: "normalize vector", tuple: NewVector(2, 3, 6), expected: HomogeneousTuple{x: 2.0 / 7, y: 3.0 / 7, z: 6.0 / 7, w: 0}},
		{name: "normalize tuple point", tuple: NewTuple(2, 3, 6, 1), expected: NewTuple(math.NaN(), math.NaN(), math.NaN(), math.NaN())},
		{name: "normalize tuple vector", tuple: NewTuple(2, 3, 6, 0), expected: HomogeneousTuple{x: 2.0 / 7, y: 3.0 / 7, z: 6.0 / 7, w: 0}},
		{name: "normalize irregular tuple", tuple: NewTuple(0, 3, 4, 12), expected: HomogeneousTuple{x: 0, y: 3.0 / 13, z: 4.0 / 13, w: 12.0 / 13}},
		{name: "normalize zero vector", tuple: NewVector(0, 0, 0), expected: NewTuple(math.NaN(), math.NaN(), math.NaN(), math.NaN())},
		{name: "normalize zero tuple", tuple: NewTuple(0, 0, 0, 0), expected: NewTuple(math.NaN(), math.NaN(), math.NaN(), math.NaN())},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.tuple.Normalize()
			if got := result.String(); got != tt.expected.String() {
				t.Errorf("Normalize() = %v, want %v", got, tt.expected.String())
			}
		})
	}
}

func TestDotProduct(t *testing.T) {
	tests := []struct {
		name     string
		a        HomogeneousTuple
		b        HomogeneousTuple
		expected float64
	}{
		{name: "dot product of two vectors", a: NewVector(1, 2, 3), b: NewVector(4, 5, 6), expected: 32.0},
		{name: "dot product of point and vector", a: NewPoint(1, 2, 3), b: NewVector(4, 5, 6), expected: math.NaN()},
		{name: "dot product of two points", a: NewPoint(1, 2, 3), b: NewPoint(4, 5, 6), expected: math.NaN()},
		{name: "dot product of tuple vectors", a: NewTuple(1, 2, 3, 0), b: NewTuple(4, 5, 6, 0), expected: 32.0},
		{name: "dot product of irregular tuples", a: NewTuple(1, 2, 3, 2), b: NewTuple(4, 5, 6, 3), expected: 38.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.a.DotProduct(tt.b)
			if math.IsNaN(result) && math.IsNaN(tt.expected) {
				return // Both results are NaN, which is acceptable
			}
			if got := result; got != tt.expected {
				t.Errorf("DotProduct() = %v, want %v", got, tt.expected)
			}
		})
	}
}
