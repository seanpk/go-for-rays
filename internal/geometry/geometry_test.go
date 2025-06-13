package geometry

import (
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
