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

// Returns a tuple with NaN values for each component.
func NaNTuple() HomogeneousTuple {
	return NewHomogeneousTuple(math.NaN(), math.NaN(), math.NaN(), math.NaN())
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

// Add returns a new HomogeneousTuple that is the sum of the two tuples.
// If both tuples are points, the result is a vector (w=0).
// If both tuples are vectors, the result is a vector (w=0).
// If one tuple is a point and the other is a vector, the result is a point (w=1).
// If both tuples are general tuples, the result is a general tuple with w being the sum of the two w values.
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

// Subtract returns a new HomogeneousTuple that is the result of subtracting the other tuple from this one.
// If both tuples are points, the result is a vector (w=0).
// If both tuples are vectors, the result is a vector (w=0).
// If one tuple is a point and the other is a vector, the result is a point (w=1).
// If both tuples are general tuples, the result is a general tuple with w being the difference of the two w values.
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

// Multiply returns a new HomogeneousTuple that is the result of multiplying each component by the scalar.
// For points and vectors, it keeps the w value unchanged.
// For general tuples, it multiplies the w value by the scalar.
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

// Divide returns a new HomogeneousTuple that is the result of dividing each component by the scalar.
// If the scalar is zero, it returns a tuple with NaN values for each component.
// The behavior for points and vectors follows the same logic as Multiply.
func (t HomogeneousTuple) Divide(scalar float64) HomogeneousTuple {
	if scalar == 0 {
		return NaNTuple()
	}

	return t.Multiply(1.0 / scalar)
}

// Negate returns a new HomogeneousTuple with each component negated.
// For points and vectors, it keeps the w value unchanged.
// For general tuples, it negates the w value as well.
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

// DotProduct calculates the dot product of two HomogeneousTuples.
// If either tuple is a point, it returns NaN since points do not have a meaningful dot product.
// If both tuples are vectors, it calculates the dot product as the sum of the products of their components (x1*x2 + y1*y2 + z1*z2).
// If either tuples is a general tuple, it includes the w component in the calculation (x1*x2 + y1*y2 + z1*z2 + w1*w2).
func (t HomogeneousTuple) DotProduct(other HomogeneousTuple) float64 {
	if t.IsPoint() || other.IsPoint() {
		return math.NaN()
	}

	threeProduct := t.X()*other.X() + t.Y()*other.Y() + t.Z()*other.Z()
	if t.IsVector() && other.IsVector() {
		return threeProduct
	}

	return threeProduct + t.W()*other.W()
}

// Magnitude returns the magnitude of the tuple.
// For points, it returns 0.0 since points do not have a magnitude.
// For vectors, it calculates the Euclidean norm (sqrt(x^2 + y^2 + z^2)).
// For general tuples, it includes the w component in the magnitude calculation.
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

// Normalize returns a new HomogeneousTuple that is the normalized version of the tuple.
// For points, it returns a tuple with NaN values since points cannot be normalized.
// For vectors, it divides each component by the magnitude of the vector.
// For general tuples, it divides each component by the magnitude of the tuple, including the w component.
func (t HomogeneousTuple) Normalize() HomogeneousTuple {
	return t.Divide(t.Magnitude())
}

func (t HomogeneousTuple) Equals(other HomogeneousTuple, epsilon ...float64) bool {
	eps := epsilonOrDefault(epsilon...)

	return IsNearTo(t.X(), other.X(), eps) &&
		IsNearTo(t.Y(), other.Y(), eps) &&
		IsNearTo(t.Z(), other.Z(), eps) &&
		IsNearTo(t.W(), other.W(), eps)
}

func (t HomogeneousTuple) IsNaN() bool {
	return math.IsNaN(t.X()) || math.IsNaN(t.Y()) || math.IsNaN(t.Z()) || math.IsNaN(t.W())
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
