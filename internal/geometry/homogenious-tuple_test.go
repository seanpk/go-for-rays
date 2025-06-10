package geometry

import (
	"testing"
)

func TestTuple(t *testing.T) {
	tests := []struct {
		name     string
		tuple    Tuple
		expected [2]bool
	}{
		{name: "w=1 is a point", tuple: NewTuple(1, 2, 3, 1), expected: [2]bool{true, false}},
		{name: "w=0 is a vector", tuple: NewTuple(1, 2, 3, 0), expected: [2]bool{false, true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPoint(tt.tuple); got != tt.expected[0] {
				t.Errorf("IsPoint() = %v, want %v", got, tt.expected[0])
			}
			if got := IsVector(tt.tuple); got != tt.expected[1] {
				t.Errorf("IsVector() = %v, want %v", got, tt.expected[1])
			}
		})
	}
}

func TestTupleConversion(t *testing.T) {
	tests := []struct {
		name     string
		tuple    Tuple
		expected string
	}{
		{name: "convert point tuple", tuple: NewTuple(1, 2, 3, 1), expected: "Point(1.000000, 2.000000, 3.000000)"},
		{name: "convert vector tuple", tuple: NewTuple(1, 2, 3, 0), expected: "Vector(1.000000, 2.000000, 3.000000)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.tuple.IsPoint() {
				point := tt.tuple.ToPoint()
				if got := point.String(); got != tt.expected {
					t.Errorf("ToPoint() = %v, want %v", got, tt.expected)
				}
			} else if tt.tuple.IsVector() {
				vector := tt.tuple.ToVector()
				if got := vector.String(); got != tt.expected {
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
		{name: "equal tuples", a: NewTuple(1, 2, 3, 1), b: NewTuple(1, 2, 3, 1), expected: true},
		{name: "unequal tuples", a: NewTuple(1, 2, 3, 1), b: NewTuple(1, 2, 4, 1), expected: false},
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
				result = IsEqual(tt.a, tt.b)
			} else {
				result = IsEqual(tt.a, tt.b, *tt.epsilon)
			}
			if got := result; got != tt.expected {
				t.Errorf("IsEqual() = %v, want %v", got, tt.expected)
			}
		})
	}
}
