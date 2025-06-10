package geometry

import (
	"fmt"
	"math"
)

const EPSILON = 1e-6

func epsilonOrDefault(epsilon ...float64) float64 {
	if len(epsilon) > 0 && epsilon[0] > 0 {
		return epsilon[0]
	}
	return EPSILON
}

func IsNearTo(a, b float64, epsilon ...float64) bool {
	eps := epsilonOrDefault(epsilon...)
	return math.Abs(a-b) < eps
}

// Can represent both points and vectors in 3D space using homogeneous coordinates (x, y, z, w).
// A point has w=1, a vector has w=0.
type HomogeneousTuple interface {
	X() float64
	Y() float64
	Z() float64
	W() float64

	Add(HomogeneousTuple) Tuple
	Equals(HomogeneousTuple, ...float64) bool
	String() string
}

func IsPoint(t HomogeneousTuple) bool {
	return IsNearTo(t.W(), 1.0)
}

func IsVector(t HomogeneousTuple) bool {
	return IsNearTo(t.W(), 0.0)
}

func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{
		x: x,
		y: y,
		z: z,
		w: w,
	}
}

func ToTuple(t HomogeneousTuple) Tuple {
	return NewTuple(t.X(), t.Y(), t.Z(), t.W())
}

func Add(a, b HomogeneousTuple) Tuple {
	var w float64
	if IsPoint(a) && IsPoint(b) {
		w = 0.0 // the result of adding two points will be interpreted as a vector
	} else {
		w = a.W() + b.W()
	}

	return NewTuple(a.X()+b.X(), a.Y()+b.Y(), a.Z()+b.Z(), w)
}

func IsEqual(a, b HomogeneousTuple, epsilon ...float64) bool {
	eps := epsilonOrDefault(epsilon...)
	return IsNearTo(a.X(), b.X(), eps) &&
		IsNearTo(a.Y(), b.Y(), eps) &&
		IsNearTo(a.Z(), b.Z(), eps) &&
		IsNearTo(a.W(), b.W(), eps)
}

// Tuple:
// A generic implementation of a homogeneous tuple; it can represent both points and vectors.
type Tuple struct {
	x, y, z, w float64
}

func (t Tuple) X() float64 {
	return t.x
}

func (t Tuple) Y() float64 {
	return t.y
}

func (t Tuple) Z() float64 {
	return t.z
}

func (t Tuple) W() float64 {
	return t.w
}

func (t Tuple) IsPoint() bool {
	return IsPoint(t)
}

func (t Tuple) IsVector() bool {
	return IsVector(t)
}

// Convert a Tuple to a Point
func (t Tuple) ToPoint() Point {
	return NewPoint(t.X(), t.Y(), t.Z())
}

// Convert a Tuple to a Vector
func (t Tuple) ToVector() Vector {
	return NewVector(t.X(), t.Y(), t.Z())
}

// Converts a Tuple to a Point; panics if the Tuple is not a point.
func (t Tuple) AsPoint() Point {
	if !t.IsPoint() {
		panic("Tuple is not a point")
	}
	return ToPoint(t)
}

// Converts a Tuple to a Vector; panics if the Tuple is not a vector.
func (t Tuple) AsVector() Vector {
	if !t.IsVector() {
		panic("Tuple is not a vector")
	}
	return ToVector(t)
}

func (t Tuple) Add(other HomogeneousTuple) Tuple {
	return Add(t, other)
}

func (t Tuple) Equals(other HomogeneousTuple, epsilon ...float64) bool {
	return IsEqual(t, other, epsilon...)
}

func (t Tuple) String() string {
	return fmt.Sprintf("Tuple(%f, %f, %f, %f)", t.X(), t.Y(), t.Z(), t.W())
}
