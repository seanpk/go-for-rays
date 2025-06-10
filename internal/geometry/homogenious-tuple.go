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

// Can represent both points and vectors in 3D space
// using homogeneous coordinates (x, y, z, w).
// A point has w=1, a vector has w=0.
type HomogeneousTuple interface {
	X() float64
	Y() float64
	Z() float64
	W() float64
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

func NewPoint(x, y, z float64) Point {
	return Point{
		x: x,
		y: y,
		z: z,
	}
}

func NewVector(x, y, z float64) Vector {
	return Vector{
		x: x,
		y: y,
		z: z,
	}
}

func ToTuple(t HomogeneousTuple) Tuple {
	return NewTuple(t.X(), t.Y(), t.Z(), t.W())
}

func ToPoint(t HomogeneousTuple) Point {
	return NewPoint(t.X(), t.Y(), t.Z())
}

func ToVector(t HomogeneousTuple) Vector {
	return NewVector(t.X(), t.Y(), t.Z())
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

func (t Tuple) String() string {
	return fmt.Sprintf("Tuple(%f, %f, %f, %f)", t.X(), t.Y(), t.Z(), t.W())
}

func (t Tuple) Equals(other HomogeneousTuple, epsilon ...float64) bool {
	return IsEqual(t, other, epsilon...)
}

// Point:
// Represents a point in 3D space (x, y, z). The w value (always 1.0) is included for interoperability with Vectors and using homogeneous coordinates (x, y, z, w).
type Point struct {
	x, y, z float64
}

func (p Point) X() float64 {
	return p.x
}

func (p Point) Y() float64 {
	return p.y
}

func (p Point) Z() float64 {
	return p.z
}

func (p Point) W() float64 {
	return 1.0
}

func (p Point) ToVector() Vector {
	return NewVector(p.X(), p.Y(), p.Z())
}

func (p Point) ToTuple() Tuple {
	return NewTuple(p.X(), p.Y(), p.Z(), p.W())
}

func (p Point) String() string {
	return fmt.Sprintf("Point(%f, %f, %f)", p.X(), p.Y(), p.Z())
}

func (p Point) Equals(other HomogeneousTuple, epsilon ...float64) bool {
	return IsEqual(p, other, epsilon...)
}

// Vector:
// Represents a vector in 3D space (x, y, z). The w value (always 0.0) is included for interoperability with Points and using homogeneous coordinates (x, y, z, w).
type Vector struct {
	x, y, z float64
}

func (v Vector) X() float64 {
	return v.x
}

func (v Vector) Y() float64 {
	return v.y
}

func (v Vector) Z() float64 {
	return v.z
}

func (v Vector) W() float64 {
	return 0.0
}

func (v Vector) ToPoint() Point {
	return NewPoint(v.X(), v.Y(), v.Z())
}

func (v Vector) ToTuple() Tuple {
	return NewTuple(v.X(), v.Y(), v.Z(), v.W())
}

func (v Vector) String() string {
	return fmt.Sprintf("Vector(%f, %f, %f)", v.X(), v.Y(), v.Z())
}

func (v Vector) Equals(other HomogeneousTuple, epsilon ...float64) bool {
	return IsEqual(v, other, epsilon...)
}
