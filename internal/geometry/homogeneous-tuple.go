package geometry

import (
	"fmt"
	"math"
)

func NewHomogeneousTuple(x, y, z, w float64) HomogeneousTuple {
	return HomogeneousTuple{
		x: x,
		y: y,
		z: z,
		w: w,
	}
}

// Shorthand for creating a tuple with homogeneous coordinates (x, y, z, w).
// This can be used for both points and vectors, depending on the value of w.
// If w=1, it represents a point; if w=0, it represents a vector.
// If w is neither 0 nor 1, it represents a general tuple.
func NewTuple(x, y, z, w float64) HomogeneousTuple {
	return NewHomogeneousTuple(x, y, z, w)
}

// Creates a new point in 3D space with homogeneous coordinates (x, y, z).
// The resulting tuple has w=1, indicating it is a point.
func NewPoint(x, y, z float64) HomogeneousTuple {
	return NewHomogeneousTuple(x, y, z, 1.0)
}

// Creates a new vector in 3D space with homogeneous coordinates (x, y, z).
// The resulting tuple has w=0, indicating it is a vector.
func NewVector(x, y, z float64) HomogeneousTuple {
	return NewHomogeneousTuple(x, y, z, 0.0)
}

// Can represent both points and vectors in 3D space using homogeneous coordinates (x, y, z, w).
// A point has w=1, a vector has w=0.
type HomogeneousTuple struct {
	x, y, z, w float64
}

func (t HomogeneousTuple) X() float64 {
	return t.x
}

func (t HomogeneousTuple) Y() float64 {
	return t.y
}

func (t HomogeneousTuple) Z() float64 {
	return t.z
}

func (t HomogeneousTuple) W() float64 {
	return t.w
}

func (t HomogeneousTuple) IsPoint() bool {
	return IsNearTo(t.w, 1.0)
}

func (t HomogeneousTuple) IsVector() bool {
	return IsNearTo(t.w, 0.0)
}

func ToPoint(t HomogeneousTuple) HomogeneousTuple {
	return NewPoint(t.x, t.y, t.z)
}

func ToVector(t HomogeneousTuple) HomogeneousTuple {
	return NewVector(t.x, t.y, t.z)
}

func (t HomogeneousTuple) Add(other HomogeneousTuple) HomogeneousTuple {
	var w float64
	if t.IsPoint() && other.IsPoint() {
		w = 0.0 // the result of adding two points will be interpreted as a vector
	} else {
		w = t.W() + other.W()
	}

	return NewTuple(
		t.X()+other.X(),
		t.Y()+other.Y(),
		t.Z()+other.Z(),
		w,
	)
}

func (t HomogeneousTuple) Divide(scalar float64) HomogeneousTuple {
	if scalar == 0 {
		panic("division by zero is not allowed")
	}

	return t.Multiply(1.0 / scalar)
}

func (t HomogeneousTuple) Magnitude() float64 {
	if t.IsPoint() {
		return 0.0 // points do not have a magnitude
	}

	squareSum := t.X()*t.X() + t.Y()*t.Y() + t.Z()*t.Z()
	if !t.IsVector() {
		squareSum += t.W() * t.W() // for general tuples, include w in the magnitude calculation
	}

	return math.Sqrt(squareSum)
}

func (t HomogeneousTuple) Multiply(scalar float64) HomogeneousTuple {
	var w float64
	if t.IsPoint() || t.IsVector() {
		w = t.W() // the multiplication of a point or vector keeps the same w value
	} else {
		w = t.W() * scalar // for general tuples, multiply the w value
	}

	return NewTuple(
		t.X()*scalar,
		t.Y()*scalar,
		t.Z()*scalar,
		w,
	)
}

func (t HomogeneousTuple) Negate() HomogeneousTuple {
	var w float64
	if t.IsPoint() || t.IsVector() {
		w = t.W() // the negation of a point or vector keeps the same w value
	} else {
		w = -t.W() // for general tuples, negate the w value
	}

	return NewTuple(
		-t.X(),
		-t.Y(),
		-t.Z(),
		w,
	)
}

func (t HomogeneousTuple) Normalize() HomogeneousTuple {
	return t.Divide(t.Magnitude())
}

func (t HomogeneousTuple) Subtract(other HomogeneousTuple) HomogeneousTuple {
	var w float64
	if t.IsPoint() && other.IsVector() || t.IsVector() && other.IsPoint() {
		w = 1.0 // the result of subtraction with a point and a vector will be interpreted as a point
	} else {
		w = t.W() - other.W()
	}

	return NewTuple(
		t.X()-other.X(),
		t.Y()-other.Y(),
		t.Z()-other.Z(),
		w,
	)
}

func (t HomogeneousTuple) Equals(other HomogeneousTuple, epsilon ...float64) bool {
	eps := epsilonOrDefault(epsilon...)

	return IsNearTo(t.X(), other.X(), eps) &&
		IsNearTo(t.Y(), other.Y(), eps) &&
		IsNearTo(t.Z(), other.Z(), eps) &&
		IsNearTo(t.W(), other.W(), eps)
}

func (t HomogeneousTuple) String() string {
	if t.IsPoint() {
		return fmt.Sprintf("Point(%f, %f, %f)", t.X(), t.Y(), t.Z())
	}
	if t.IsVector() {
		return fmt.Sprintf("Vector(%f, %f, %f)", t.X(), t.Y(), t.Z())
	}

	return fmt.Sprintf("Tuple(%f, %f, %f, %f)", t.X(), t.Y(), t.Z(), t.W())
}
